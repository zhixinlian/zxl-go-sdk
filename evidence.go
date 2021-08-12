package zxl_go_sdk

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
)

const (
	InvalidAppId   = "无效的appId"
	InvalidAppKey  = "无效的appKey"
	BadCipherKey   = "加密密码不能为空"
	BadCipherData  = "加密内容不能为空"

)


type ZxlConfig struct {
	AppId      string
	AppKey     string
	ServerAddr string
	IsProxy    bool
	ProxyHost  string
	ProxyPort  string
}

// Deprecated
func NewZxlImpl(appId, appKey string) (*ZxlImpl, error) {
	if len(appId) < 15 {
		return nil, errors.New(InvalidAppId)
	}
	if len(appKey) == 0 {
		return nil, errors.New(InvalidAppKey)
	}
	typeInt, err := strconv.Atoi(appId[10:11])
	if err != nil {
		return nil, errors.New(InvalidAppId)
	}
	return &ZxlImpl{
		appId: appId,
		appKey: appKey,
		appType: typeInt}, nil
}


func CreateZxlClientWithConfig(config ZxlConfig)(*ZxlImpl, error) {
	if len(config.AppId) < 15 {
		return nil, errors.New(InvalidAppId)
	}
	if len(config.AppKey) == 0 {
		return nil, errors.New(InvalidAppKey)
	}
	typeInt, err := strconv.Atoi(config.AppId[10:11])
	if err != nil {
		return nil, errors.New(InvalidAppId)
	}
	if config.ServerAddr != "" {
		defConf.ServerAddr = config.ServerAddr
	}
	// 配置代理
	if config.IsProxy == true {
		defConf.IsProxy = config.IsProxy
		defConf.ProxyHost = config.ProxyHost
		defConf.ProxyPort = config.ProxyPort
	}
	return &ZxlImpl{
		appId: config.AppId,
		appKey: config.AppKey,
		appType: typeInt,
		}, nil
}


// Deprecated
func NewZxlImplWithConfig(config ZxlConfig) (*ZxlImpl, error) {
	if len(config.AppId) < 15 {
		return nil, errors.New(InvalidAppId)
	}
	if len(config.AppKey) == 0 {
		return nil, errors.New(InvalidAppKey)
	}
	typeInt, err := strconv.Atoi(config.AppId[10:11])
	if err != nil {
		return nil, errors.New(InvalidAppId)
	}
	defConf.ServerAddr = config.ServerAddr
	return &ZxlImpl{
		appId: config.AppId,
		appKey:    config.AppKey,
		appType:   typeInt,
		}, nil
}

type ZxlImpl struct {
	appKey  string
	appId   string
	appType int
}

//绑定用户证书
func (zxl *ZxlImpl) BindUserCert(pk, sk string, timeout time.Duration) error {
	rawData := strings.Join([]string{zxl.appId, pk}, ",")
	signedStr, err := zxl.Sign(sk, []byte(rawData))
	if err != nil {
		return errors.New("BindUserCertError (Sign): " + err.Error())
	}
	certReq := UserCertReq{Pk: pk, Sign: signedStr}
	dataBytes, err := json.Marshal(&certReq)
	if err != nil {
		return errors.New("BindUserCertError (Marshal): " + err.Error())
	}
	_, cri, err := sendRequest(zxl.appId, zxl.appKey, "POST", defConf.ServerAddr+defConf.UserCert, dataBytes, timeout)
	if err != nil {
		return errors.New("BindUserCertError (sendRequest): " + err.Error() + ", requestId:" + cri.RequestId)
	}
	return nil
}

//更新用户证书
func (zxl *ZxlImpl) UpdateUserCert(pk, sk string, timeout time.Duration) error {
	rawData := strings.Join([]string{zxl.appId, pk}, ",")
	signedStr, err := zxl.Sign(sk, []byte(rawData))
	if err != nil {
		return errors.New("UpdateUserCert (Sign): " + err.Error())
	}
	certReq := UserCertReq{Pk: pk, Sign: signedStr}
	dataBytes, err := json.Marshal(&certReq)
	if err != nil {
		return errors.New("UpdateUserCert (Marshal): " + err.Error())
	}
	_, cri, err := sendRequest(zxl.appId, zxl.appKey, "PUT", defConf.ServerAddr+defConf.UserCert, dataBytes, timeout)
	if err != nil {
		return errors.New("UpdateUserCert (sendRequest): " + err.Error() + ", requestId:" + cri.RequestId)
	}

	return nil
}

//加密信息
func (zxl *ZxlImpl) EncryptData(pwd string, rawData []byte) (string, error) {
	if len(pwd) == 0 {
		return "", errors.New(BadCipherKey)
	}
	if len(rawData) == 0 {
		return "", errors.New(BadCipherData)
	}

	pwdData := sm3.SumSM3([]byte(pwd))

	block, err := aes.NewCipher(pwdData)
	if err != nil {
		return "", err
	}
	tmpData := padding(rawData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, pwdData[:block.BlockSize()])
	blockMode.CryptBlocks(tmpData, tmpData)
	return hex.EncodeToString(tmpData), nil
}

//解密信息
func (zxl *ZxlImpl) DecryptData(pwd string, encryptedData string) ([]byte, error) {
	if len(pwd) == 0 {
		return nil, errors.New(BadCipherKey)
	}
	if len(encryptedData) == 0 {
		return nil, errors.New(BadCipherData)
	}
	pwdData := sm3.SumSM3([]byte(pwd))

	block, _ := aes.NewCipher(pwdData)
	blockMode := cipher.NewCBCDecrypter(block, pwdData[:block.BlockSize()])

	encryptedBytes, err := hex.DecodeString(encryptedData)
	if err != nil {
		return nil, errors.New("data format error")
	}
	blockMode.CryptBlocks(encryptedBytes, encryptedBytes)
	return unpadding(encryptedBytes)

}

//通过证据id查找证据
func (zxl *ZxlImpl) QueryWithEvId(evId string, timeout time.Duration) ([]QueryResp, error) {
	if len(evId) == 0 {
		return nil, errors.New("evId 不能为空")
	}
	respBytes, cri, err := sendRequest(zxl.appId, zxl.appKey, "GET", defConf.ServerAddr+defConf.QueryWithEvId+evId, nil, timeout)
	if err != nil {
		return nil, errors.New("QueryWithEvId (sendRequest) error: " + err.Error() + ", requestId:" + cri.RequestId)
	}
	var result []QueryResp
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, errors.New("QueryWithEvId (Unmarshal) error: " + err.Error())
	}
	return result, nil
}

//通过交易id查找证据
func (zxl *ZxlImpl) QueryWithTxHash(txHash string, timeout time.Duration) ([]QueryResp, error) {
	if len(txHash) == 0 {
		return nil, errors.New("txHash 不能为空")
	}
	respBytes, cri, err := sendRequest(zxl.appId, zxl.appKey, "GET", defConf.ServerAddr+defConf.QueryWithTxHash+txHash, nil, timeout)
	if err != nil {
		return nil, errors.New("QueryWithTxHash (sendRequest) error: " + err.Error() + ", requestId:" + cri.RequestId)
	}
	var result []QueryResp
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, errors.New("QueryWithTxHash (Unmarshal) error: " + err.Error())
	}
	return result, nil
}

//通过证据hash查找证据
func (zxl *ZxlImpl) QueryWithEvHash(evHash string, timeout time.Duration) ([]QueryResp, error) {

	if len(evHash) == 0 {
		return nil, errors.New("evHash 不能为空")
	}
	respBytes, cri, err := sendRequest(zxl.appId, zxl.appKey, "GET", defConf.ServerAddr+defConf.QueryWithEvHash+evHash, nil, timeout)
	if err != nil {
		return nil, errors.New("QueryWithEvHash (sendRequest) error: " + err.Error() + ", requestId:" + cri.RequestId)
	}
	var result []QueryResp
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, errors.New("QueryWithEvHash (Unmarshal) error: " + err.Error())
	}
	return result, nil
}

//任意输入查找证据
func (zxl *ZxlImpl) QueryWithHash(hash string, timeout time.Duration) ([]QueryResp, error) {

	if len(hash) == 0 {
		return nil, errors.New("hash 不能为空")
	}
	respBytes, cri, err := sendRequest(zxl.appId, zxl.appKey, "GET", defConf.ServerAddr+defConf.QueryWithHash+hash, nil, timeout)
	if err != nil {
		return nil, errors.New("QueryWithHash (sendRequest) error: " + err.Error() + ", requestId:" + cri.RequestId)
	}
	var result []QueryResp
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, errors.New("QueryWithHash (Unmarshal) error: " + err.Error())
	}
	return result, nil
}

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) ([]byte, error) {
	n := len(src)
	unpadnum := int(src[n-1])
	if unpadnum > n {
		return nil, errors.New("password error")
	}
	return src[:n-unpadnum], nil
}
