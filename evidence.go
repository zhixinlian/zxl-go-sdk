package zxl_go_sdk

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"strconv"
	"strings"
	"time"
)

const (
	InvalidAppId   = "无效的appId"
	InvalidAppKey  = "无效的appKey"
	InvalidAppType = "无效的appType"
	BadCipherKey   = "加密密码不能为空"
	BadCipherData  = "加密内容不能为空"

	TypeTengXun = 0
	TypeWangAn  = 1
)

type zxlCipher interface {
	//生成公私钥对
	GenerateKeyPair() (pk string, sk string, err error)
	//签名
	Sign(sk string, rawData []byte) (string, error)
	//验证签名
	Verify(pk, signedStr string, rawData []byte) (bool, error)
	//保存证据
	EvidenceSave(evHash, extendInfo, sk, pk string, timeout time.Duration) (*EvSaveResult, error)
	//计算文件hash
	CalculateHash(path string) (string, error)
	//计算字符串hash
	CalculateStrHash(str string) (string, error)
	//录屏任务
	ContentCaptureVideo(webUrls string, timeout time.Duration) (string, error)
	//截屏任务
	ContentCapturePic(webUrls string, timeout time.Duration) (string, error)
	//获取录屏、截屏任务状态及结果
	GetContentStatus(orderNo string, timeout time.Duration) (*TaskEvData, error)
	//视频取证接口
	EvidenceObtainVideo(webUrls, title, remark string, timeout time.Duration) (string, error)
	//图片取证接口
	EvidenceObtainPic(webUrls, title, remark string, timeout time.Duration) (string, error)
	//获取取证证书任务状态及结果
	GetEvidenceStatus(orderNo string, timeout time.Duration) (*EvIdData, error)
	//代理商模式视频取证接口
	RepresentEvidenceObtainVideo(webUrls, title, remark, representAppId string, timeout time.Duration) (string, error)
	//代理商模式图片取证接口
	RepresentEvidenceObtainPic(webUrls, title, remark, representAppId string, timeout time.Duration) (string, error)
	//代理商模式获取取证证书任务状态及结果
	RepresentGetEvidenceStatus(orderNo, representAppId string, timeout time.Duration) (*EvIdData, error)
}

func NewZxlImpl(appId, appKey string) (*zxlImpl, error) {
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
	return &zxlImpl{appId: appId, appKey: appKey, appType: typeInt, zxlCipher: &cetcSDKImpl{AppId: appId, AppKey: appKey}}, nil
	//if typeInt == TypeTengXun {
	//	return &zxlImpl{appId: appId, appKey: appKey, appType:typeInt, zxlCipher: &trustSDKImpl{AppId:appId, AppKey:appKey}}, nil
	//} else if typeInt == TypeWangAn {
	//	return &zxlImpl{appId: appId, appKey: appKey, appType:typeInt, zxlCipher: &cetcSDKImpl{AppId:appId, AppKey:appKey}}, nil
	//} else {
	//	return nil, errors.New(InvalidAppType)
	//}
}

type zxlImpl struct {
	zxlCipher
	appKey  string
	appId   string
	appType int
}

//绑定用户证书
func (zxl *zxlImpl) BindUserCert(pk, sk string, timeout time.Duration) error {
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
	_, err = sendRequest(zxl.appId, zxl.appKey, "POST", defConf.ServerAddr+defConf.UserCert, dataBytes, timeout)
	if err != nil {
		return errors.New("BindUserCertError (sendRequest): " + err.Error())
	}

	//var bindResp UserCertResp
	//err = json.Unmarshal(respBytes, &bindResp)
	//if err != nil {
	//	return errors.New("BindUserCertError (Unmarshal): " + err.Error())
	//}
	return nil
}

//更新用户证书
func (zxl *zxlImpl) UpdateUserCert(pk, sk string, timeout time.Duration) error {
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
	_, err = sendRequest(zxl.appId, zxl.appKey, "PUT", defConf.ServerAddr+defConf.UserCert, dataBytes, timeout)
	if err != nil {
		return errors.New("UpdateUserCert (sendRequest): " + err.Error())
	}

	return nil
}

//加密信息
func (zxl *zxlImpl) EncryptData(pwd string, rawData []byte) (string, error) {
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
func (zxl *zxlImpl) DecryptData(pwd string, encryptedData string) ([]byte, error) {
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
func (zxl *zxlImpl) QueryWithEvId(evId string, timeout time.Duration) ([]QueryResp, error) {
	if len(evId) == 0 {
		return nil, errors.New("evId 不能为空")
	}
	respBytes, err := sendRequest(zxl.appId, zxl.appKey, "GET", defConf.ServerAddr+defConf.QueryWithEvId+evId, nil, timeout)
	if err != nil {
		return nil, errors.New("QueryWithEvId (sendRequest) error: " + err.Error())
	}
	var result []QueryResp
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, errors.New("QueryWithEvId (Unmarshal) error: " + err.Error())
	}
	return result, nil
}

//通过交易id查找证据
func (zxl *zxlImpl) QueryWithTxHash(txHash string, timeout time.Duration) ([]QueryResp, error) {
	if len(txHash) == 0 {
		return nil, errors.New("txHash 不能为空")
	}
	respBytes, err := sendRequest(zxl.appId, zxl.appKey, "GET", defConf.ServerAddr+defConf.QueryWithTxHash+txHash, nil, timeout)
	if err != nil {
		return nil, errors.New("QueryWithTxHash (sendRequest) error: " + err.Error())
	}
	var result []QueryResp
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, errors.New("QueryWithTxHash (Unmarshal) error: " + err.Error())
	}
	return result, nil
}

//通过证据hash查找证据
func (zxl *zxlImpl) QueryWithEvHash(evHash string, timeout time.Duration) ([]QueryResp, error) {

	if len(evHash) == 0 {
		return nil, errors.New("evHash 不能为空")
	}
	respBytes, err := sendRequest(zxl.appId, zxl.appKey, "GET", defConf.ServerAddr+defConf.QueryWithEvHash+evHash, nil, timeout)
	if err != nil {
		return nil, errors.New("QueryWithEvHash (sendRequest) error: " + err.Error())
	}
	var result []QueryResp
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, errors.New("QueryWithEvHash (Unmarshal) error: " + err.Error())
	}
	return result, nil
}

//任意输入查找证据
func (zxl *zxlImpl) QueryWithHash(hash string, timeout time.Duration) ([]QueryResp, error) {

	if len(hash) == 0 {
		return nil, errors.New("hash 不能为空")
	}
	respBytes, err := sendRequest(zxl.appId, zxl.appKey, "GET", defConf.ServerAddr+defConf.QueryWithHash+hash, nil, timeout)
	if err != nil {
		return nil, errors.New("QueryWithHash (sendRequest) error: " + err.Error())
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
