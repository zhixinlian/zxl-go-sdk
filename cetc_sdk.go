package zxl_go_sdk

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm2"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"io/ioutil"
	"strings"
	"time"
)

type cetcSDKImpl struct {
	AppId  string
	AppKey string
}

// 生成公司钥对
// 返回值  pk公钥（string），sk私钥（string），err错误信息（error）
func (sdk *cetcSDKImpl) GenerateKeyPair() (pk string, sk string, err error) {
	privateKey, _ := sm2.GenerateKey(rand.Reader)
	pk = sm2.EncodePubKey(&privateKey.PublicKey)
	sk = sm2.EncodePrivKey(privateKey)
	return
}

func (sdk *cetcSDKImpl) Sign(prvKey string, data []byte) (string, error) {
	sk, err := sm2.DecodePrivKey(prvKey)
	if err != nil {
		return "", errors.New("Sign (DecodePriKey) error ")
	}
	signBytes, err := sk.Sign(rand.Reader, data, nil)
	if err != nil {
		return "", errors.New("Sign (Sign) error ")
	}
	return hex.EncodeToString(signBytes), nil
}

func (sdk *cetcSDKImpl) Verify(pubKey string, sign string, data []byte) (bool, error) {
	signBytes, err := hex.DecodeString(sign)
	if err != nil {
		return false, errors.New("Verify (DecodeString) error ")
	}
	pk, err := sm2.DecodePubKey(pubKey)
	if err != nil {
		return false, errors.New("Verify (DecodePubKey) error ")
	}
	return pk.Verify(data, signBytes)
}

func (sdk *cetcSDKImpl) EvidenceSave(evHash, extendInfo, sk, pk string, timeout time.Duration) (*EvSaveResult, error) {
	uid, err := generateUid()
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc generateUid) error:" + err.Error())
	}
	ed := CetcEvidenceReq{EvId: uid, EvHash: evHash, ExtendInfo: extendInfo}
	rawStr := []byte(strings.Join([]string{sdk.AppId, ed.EvHash, ed.ExtendInfo, ed.EvId}, ","))
	signStr, err := sdk.Sign(sk, rawStr)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc Sign) error:" + err.Error())
	}
	ed.Sign = signStr

	bodyData, _ := json.Marshal(&ed)
	respBytes, err := sendRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.EvidenceSave, bodyData, timeout)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc sendRequest) error:" + err.Error())
	}
	var saveResp EvSaveResult
	err = json.Unmarshal(respBytes, &saveResp)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc Unmarshal) error:" + err.Error())
	}
	saveResp.EvHash = evHash
	saveResp.EvId = uid
	return &saveResp, nil
}

func (sdk *cetcSDKImpl) CalculateHash(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.New("CalculateHash (ReadFile) error:" + err.Error())
	}
	dataHash := sm3.SumSM3(data)
	return hex.EncodeToString(dataHash), nil
}
func (sdk *cetcSDKImpl) CalculateStrHash(str string) (string, error) {
	signByte := []byte(str)
	dataHash := sm3.SumSM3(signByte)
	return hex.EncodeToString(dataHash), nil
}

//下发截屏任务到取证工具服务
//_, err =sendRequest(zxl.appId, zxl.appKey, "POST", defConf.ServerAddr + defConf.UserCert, dataBytes, timeout)
func (sdk *cetcSDKImpl) ContentCaptureVideo(webUrls string, timeout time.Duration) (string, error) {
	if len(webUrls) == 0 {
		return "", errors.New("webUrls 不能为空")
	}
	param := EvObtainTask{WebUrls: webUrls, Type: 1, AppId: sdk.AppId, RequestType: "POST", RedirectUrl: "zhixin-api/v2/screenshot/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	applyRetBytes, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST",
		defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New("下发任务异常>>error:" + err.Error())
	}
	var applyResp = applyRetBytes.OrderNo
	return applyResp, errors.New("ssss")
}

//下发录屏任务到取证工具服务
func (sdk *cetcSDKImpl) ContentCapturePic(webUrls string, timeout time.Duration) (string, error) {
	if len(webUrls) == 0 {
		return "", errors.New("webUrls 不能为空")
	}
	param := EvObtainTask{WebUrls: webUrls, Type: 0, AppId: sdk.AppId, RequestType: "POST", RedirectUrl: "zhixin-api/v2/screenshot/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New("下发任务异常>>error:" + err.Error())
	}
	var retResp = sendRetBytes.OrderNo
	return retResp, nil
}
func (sdk *cetcSDKImpl) GetContentStatus(orderNo string, timeout time.Duration) (*TaskEvData, error) {
	if len(orderNo) == 0 {
		return nil, errors.New("orderNo 不能为空")
	}
	param := EvObtainTask{AppId: sdk.AppId, OrderNo: orderNo, RequestType: "GET", RedirectUrl: "zhixin-api/v2/screenshot/evobtain/evidinfo"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	var taskEvData = TaskEvData{Hash: sendRetBytes.Hash, StatusMsg: sendRetBytes.StatusMsg, Status: sendRetBytes.Status, Url: sendRetBytes.Url}
	return &taskEvData, nil
}

//视频取证接口
func (sdk *cetcSDKImpl) EvidenceObtainVideo(webUrls, title, remark string, timeout time.Duration) (string, error) {
	if len(webUrls) == 0 || len(title) == 0 {
		return "", errors.New("webUrls or title 不能为空")
	}
	param := EvObtainTask{AppId: sdk.AppId, WebUrls: webUrls, Title: title, Type: 1, Remark: remark, RequestType: "POST", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New(err.Error())
	}
	var orderNo = sendRetBytes.OrderNo
	return orderNo, nil
}

//图片取证接口
func (sdk *cetcSDKImpl) EvidenceObtainPic(webUrls, title, remark string, timeout time.Duration) (string, error) {
	if len(webUrls) == 0 || len(title) == 0 {
		return "", errors.New("webUrls or title 不能为空")
	}
	param := EvObtainTask{AppId: sdk.AppId, WebUrls: webUrls, Title: title, Type: 0, Remark: remark, RequestType: "POST", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New(err.Error())
	}
	var orderNo = sendRetBytes.OrderNo
	return orderNo, nil
}

//获取取证证书任务状态及结果
func (sdk *cetcSDKImpl) GetEvidenceStatus(orderNo string, timeout time.Duration) (*EvIdData, error) {
	if len(orderNo) == 0 {
		return nil, errors.New("orderNo 不能为空")
	}
	param := EvObtainTask{AppId: sdk.AppId, OrderNo: orderNo, RequestType: "GET", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/evidinfo"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	var evIdData = EvIdData{Status: sendRetBytes.Status, EvidUrl: sendRetBytes.EvIdUrl, VoucherUrl: sendRetBytes.VoucherUrl}
	return &evIdData, nil
}
