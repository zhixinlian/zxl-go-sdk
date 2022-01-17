package zxl_go_sdk

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
	"time"

	"github.com/zhixinlian/zxl-go-sdk/sm/sm2"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
)

const DEFAULT_VIDEO_DURATION = 300 // 默认录屏时长5分钟

// 生成公司钥对
// 返回值  pk公钥（string），sk私钥（string），err错误信息（error）
func (sdk *ZxlImpl) GenerateKeyPair() (pk string, sk string, err error) {
	privateKey, _ := sm2.GenerateKey(rand.Reader)
	pk = sm2.EncodePubKey(&privateKey.PublicKey)
	sk = sm2.EncodePrivKey(privateKey)
	return
}

func (sdk *ZxlImpl) Sign(prvKey string, data []byte) (string, error) {
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

func (sdk *ZxlImpl) Verify(pubKey string, sign string, data []byte) (bool, error) {
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

func (sdk *ZxlImpl) EvidenceSave(evHash, extendInfo, sk, pk string, timeout time.Duration) (*EvSaveResult, error) {
	uid, err := generateUid()
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc generateUid) error:" + err.Error())
	}
	ed := CetcEvidenceReq{EvId: uid, EvHash: evHash, ExtendInfo: extendInfo}
	rawStr := []byte(strings.Join([]string{sdk.appId, ed.EvHash, ed.ExtendInfo, ed.EvId}, ","))
	signStr, err := sdk.Sign(sk, rawStr)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc Sign) error:" + err.Error())
	}
	ed.Sign = signStr

	bodyData, _ := json.Marshal(&ed)
	respBytes, cri, err := sendRequest(sdk.appId, sdk.appKey, "POST", defConf.ServerAddr+defConf.EvidenceSave, bodyData, timeout)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc sendRequest) error:" + err.Error() + ", requestId:" + cri.RequestId)
	}
	var saveResp EvSaveResult
	err = json.Unmarshal(respBytes, &saveResp)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc Unmarshal) error:" + err.Error())
	}
	saveResp.EvHash = evHash
	saveResp.EvId = uid
	saveResp.RequestId = cri.RequestId
	return &saveResp, nil
}

func (sdk *ZxlImpl) CalculateHash(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.New("CalculateHash (ReadFile) error:" + err.Error())
	}
	dataHash := sm3.SumSM3(data)
	return hex.EncodeToString(dataHash), nil
}
func (sdk *ZxlImpl) CalculateStrHash(str string) (string, error) {
	signByte := []byte(str)
	dataHash := sm3.SumSM3(signByte)
	return hex.EncodeToString(dataHash), nil
}

type CaptureVideoOption struct {
	WebUrls  string
	Duration int
}

//下发录屏任务到取证工具服务
func (sdk *ZxlImpl) ContentCaptureVideo(webUrls string, timeout time.Duration) (string, error) {
	op := CaptureVideoOption{WebUrls: webUrls, Duration: DEFAULT_VIDEO_DURATION}
	return sdk.NewContentCaptureVideo(&op, timeout)
}

//下发录屏任务到取证工具服务增加录屏时长
func (sdk *ZxlImpl) NewContentCaptureVideo(captureVideoOption *CaptureVideoOption, timeout time.Duration) (string, error) {
	if len(captureVideoOption.WebUrls) == 0 {
		return "", errors.New("webUrls 不能为空")
	}
	duration := captureVideoOption.Duration
	if captureVideoOption.Duration > 60*60 {
		return "", errors.New("duration 录屏任务不能超过1小时")
	}
	if captureVideoOption.Duration < 0 {
		return "", errors.New("duration 录屏任务时间错误")
	}
	if captureVideoOption.Duration == 0 {
		duration = DEFAULT_VIDEO_DURATION
	}
	param := EvObtainTask{WebUrls: captureVideoOption.WebUrls, Type: 2, AppId: sdk.appId, Duration: duration, RequestType: "POST", RedirectUrl: "zhixin-api/v2/screenshot/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	applyRetBytes, cri, err := sendTxMidRequest(sdk.appId, sdk.appKey, "POST",
		defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New("下发任务异常>>error:" + err.Error() + ", requestId:" + cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(applyRetBytes, &txRetDetail)
	var applyResp = txRetDetail.OrderNo
	return applyResp, nil
}

//下发截屏任务到取证工具服务
func (sdk *ZxlImpl) ContentCapturePic(webUrls string, timeout time.Duration) (string, error) {
	if len(webUrls) == 0 {
		return "", errors.New("webUrls 不能为空")
	}
	param := EvObtainTask{WebUrls: webUrls, Type: 1, AppId: sdk.appId, RequestType: "POST", RedirectUrl: "zhixin-api/v2/screenshot/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.appId, sdk.appKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New("下发任务异常>>error:" + err.Error() + ", requestId:" + cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var retResp = txRetDetail.OrderNo
	return retResp, nil
}
func (sdk *ZxlImpl) GetContentStatus(orderNo string, timeout time.Duration) (*TaskEvData, error) {
	if len(orderNo) == 0 {
		return nil, errors.New("orderNo 不能为空")
	}
	param := EvObtainTask{AppId: sdk.appId, OrderNo: orderNo, RequestType: "GET", RedirectUrl: "zhixin-api/v2/screenshot/evobtain/evidinfo"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.appId, sdk.appKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return nil, errors.New(err.Error() + ", requestId:" + cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var taskEvData = TaskEvData{
		Hash:      txRetDetail.Hash,
		StatusMsg: txRetDetail.StatusMsg,
		Status:    txRetDetail.Status,
		Url:       txRetDetail.Url,
		RequestId: cri.RequestId,
	}
	return &taskEvData, nil
}

type ObtainVideoOption struct {
	WebUrls        string
	Title          string
	Remark         string
	RepresentAppId string
	Duration       int
}

//视频取证接口
func (sdk *ZxlImpl) EvidenceObtainVideo(webUrls, title, remark string, timeout time.Duration) (string, error) {
	return sdk.NewEvidenceObtainVideo(&ObtainVideoOption{WebUrls: webUrls, Title: title, Remark: remark, RepresentAppId: "", Duration: DEFAULT_VIDEO_DURATION}, timeout)
}

//代理用户视频取证接口
func (sdk *ZxlImpl) RepresentEvidenceObtainVideo(webUrls, title, remark, representAppId string, timeout time.Duration) (string, error) {
	return sdk.NewEvidenceObtainVideo(&ObtainVideoOption{WebUrls: webUrls, Title: title, Remark: remark, RepresentAppId: representAppId, Duration: DEFAULT_VIDEO_DURATION}, timeout)
}

func (sdk *ZxlImpl) NewEvidenceObtainVideo(obtainVideoOption *ObtainVideoOption, timeout time.Duration) (string, error) {
	if len(obtainVideoOption.WebUrls) == 0 || len(obtainVideoOption.Title) == 0 {
		return "", errors.New("webUrls or title 不能为空")
	}
	duration := obtainVideoOption.Duration
	if obtainVideoOption.Duration > 60*60 {
		return "", errors.New("duration 录屏任务不能超过1小时")
	}
	if obtainVideoOption.Duration < 0 {
		return "", errors.New("duration 录屏任务时间错误")
	}
	if obtainVideoOption.Duration == 0 {
		duration = DEFAULT_VIDEO_DURATION
	}
	param := EvObtainTask{AppId: sdk.appId, WebUrls: obtainVideoOption.WebUrls, Title: obtainVideoOption.Title, Type: 2, Duration: duration, RepresentAppId: obtainVideoOption.RepresentAppId, Remark: obtainVideoOption.Remark, RequestType: "POST", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.appId, sdk.appKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New(err.Error() + ", requestId:" + cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var orderNo = txRetDetail.OrderNo
	return orderNo, nil
}

//图片取证接口
func (sdk *ZxlImpl) EvidenceObtainPic(webUrls, title, remark string, timeout time.Duration) (string, error) {
	return sdk.evidenceObtainPic(webUrls, title, remark, "", timeout)
}

//代理用户图片取证接口
func (sdk *ZxlImpl) RepresentEvidenceObtainPic(webUrls, title, remark, representAppId string, timeout time.Duration) (string, error) {
	return sdk.evidenceObtainPic(webUrls, title, remark, representAppId, timeout)
}

func (sdk *ZxlImpl) evidenceObtainPic(webUrls, title, remark, representAppId string, timeout time.Duration) (string, error) {
	if len(webUrls) == 0 || len(title) == 0 {
		return "", errors.New("webUrls or title 不能为空")
	}
	param := EvObtainTask{AppId: sdk.appId, WebUrls: webUrls, Title: title, Type: 1, RepresentAppId: representAppId, Remark: remark, RequestType: "POST", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.appId, sdk.appKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New(err.Error() + ", requestId:" + cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var orderNo = txRetDetail.OrderNo
	return orderNo, nil
}

//云桌面取证接口
func (sdk *ZxlImpl) EvidenceObtainCvd(title, remark string, timeout time.Duration) (*TxRetDetail, error) {
	return sdk.NewEvidenceObtainCvd(&ObtainVideoOption{Title: title, Remark: remark, RepresentAppId: "",
		Duration: DEFAULT_VIDEO_DURATION}, timeout)
}

//代理用户云桌面取证接口
func (sdk *ZxlImpl) RepresentEvidenceObtainCvd(title, remark, representAppId string,
	timeout time.Duration) (*TxRetDetail, error) {
	return sdk.NewEvidenceObtainCvd(&ObtainVideoOption{Title: title, Remark: remark,
		RepresentAppId: representAppId, Duration: DEFAULT_VIDEO_DURATION}, timeout)
}

func (sdk *ZxlImpl) NewEvidenceObtainCvd(obtainVideoOption *ObtainVideoOption, timeout time.Duration) (*TxRetDetail,
	error) {
	if len(obtainVideoOption.Title) == 0 {
		return nil, errors.New("title 不能为空")
	}
	param := EvObtainTask{AppId: sdk.appId, Title: obtainVideoOption.Title, Type: 4,
		RepresentAppId: obtainVideoOption.RepresentAppId, Remark: obtainVideoOption.Remark, RequestType: "POST",
		RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/cvdobtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.appId, sdk.appKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return nil, errors.New(err.Error() + ", requestId:" + cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	return &txRetDetail, nil
}

func (sdk *ZxlImpl) GetEvidenceStatus(orderNo string, timeout time.Duration) (*EvIdData, error) {
	return sdk.getEvidenceStatus(orderNo, "", timeout)
}

//获取取证证书任务状态及结果
func (sdk *ZxlImpl) RepresentGetEvidenceStatus(orderNo, representAppId string, timeout time.Duration) (*EvIdData, error) {
	return sdk.getEvidenceStatus(orderNo, representAppId, timeout)
}

//获取取证证书任务状态及结果
func (sdk *ZxlImpl) RepresentZblGetEvidenceStatus(orderNo, representAppId string,
	timeout time.Duration) (*EvIdDataZbl, error) {
	return sdk.getZblEvidenceStatus(orderNo, representAppId, timeout)
}

func (sdk *ZxlImpl) getEvidenceStatus(orderNo, representAppId string, timeout time.Duration) (*EvIdData, error) {
	if len(orderNo) == 0 {
		return nil, errors.New("orderNo 不能为空")
	}

	appId := sdk.appId
	if representAppId != "" {
		appId = representAppId
	}

	param := EvObtainTask{AppId: appId, OrderNo: orderNo, RequestType: "GET", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/evidinfo"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.appId, sdk.appKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return nil, errors.New(err.Error() + ", requestId:" + cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var evIdData = EvIdData{
		Status:     txRetDetail.Status,
		EvidUrl:    txRetDetail.EvIdUrl,
		VoucherUrl: txRetDetail.VoucherUrl,
		WebTitle:   txRetDetail.WebTitle,
		RequestId:  cri.RequestId,
		Duration:   txRetDetail.Duration,
	}
	return &evIdData, nil
}

func (sdk *ZxlImpl) getZblEvidenceStatus(orderNo, representAppId string, timeout time.Duration) (*EvIdDataZbl,
	error) {
	if len(orderNo) == 0 {
		return nil, errors.New("orderNo 不能为空")
	}

	appId := sdk.appId
	if representAppId != "" {
		appId = representAppId
	}

	param := EvObtainTask{AppId: appId, OrderNo: orderNo, RequestType: "GET", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/evidinfo"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.appId, sdk.appKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return nil, errors.New(err.Error() + ", requestId:" + cri.RequestId)
	}
	var evIdData EvIdDataZbl
	err = json.Unmarshal(sendRetBytes, &evIdData)
	if err != nil {
		return nil, errors.New(err.Error() + ", requestId:" + cri.RequestId)
	}
	return &evIdData, nil
}
