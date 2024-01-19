/*
 Package zxl_go_sdk zxl_go_sdk-ev_obtain
    - @File:  ev_obtain
    - @Description: // TODO:
    - @Author: suxiongye
    - @Date: 2023/12/3 15:24
    - @Copyright: Tencent
*/

package zxl_go_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type CaptureVideoOption struct {
	WebUrls  string
	Duration int
}

// ContentCaptureVideo 下发录屏任务到取证工具服务
func (sdk *ZxlImpl) ContentCaptureVideo(webUrls string, timeout time.Duration) (string, error) {
	op := CaptureVideoOption{WebUrls: webUrls, Duration: DEFAULT_VIDEO_DURATION}
	return sdk.NewContentCaptureVideo(&op, timeout)
}

// NewContentCaptureVideo 下发录屏任务到取证工具服务增加录屏时长
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

// ContentCapturePic 下发截屏任务到取证工具服务
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

// GetContentStatus 查询取证状态
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

// EvidenceObtainVideo 视频取证接口
func (sdk *ZxlImpl) EvidenceObtainVideo(webUrls, title, remark string, timeout time.Duration) (string, error) {
	return sdk.NewEvidenceObtainVideo(&ObtainVideoOption{WebUrls: webUrls, Title: title, Remark: remark, RepresentAppId: "", Duration: DEFAULT_VIDEO_DURATION}, timeout)
}

// RepresentEvidenceObtainVideo 代理用户视频取证接口
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

type ObtainMobileOption struct {
	ShareUrl       string
	Title          string
	Remark         string
	AppName        string
	Duration       int
	ReqTime        int64
	RepresentAppId string
}

// EvidenceObtainMobile 手机取证接口
func (sdk *ZxlImpl) EvidenceObtainMobile(shareUrl, appName, title, remark string, duration int,
	timeout time.Duration) (string, error) {
	return sdk.NewEvidenceObtainMobile(&ObtainMobileOption{ShareUrl: shareUrl, AppName: appName, Title: title,
		Remark: remark, RepresentAppId: "", Duration: duration}, timeout)
}

// RepresentEvidenceObtainMobile 代理用户手机取证接口
func (sdk *ZxlImpl) RepresentEvidenceObtainMobile(shareUrl, appName, title, remark,
	representAppId string, duration int,
	timeout time.Duration) (string, error) {
	return sdk.NewEvidenceObtainMobile(&ObtainMobileOption{ShareUrl: shareUrl, AppName: appName, Title: title, Remark: remark,
		RepresentAppId: representAppId, Duration: duration}, timeout)
}

func (sdk *ZxlImpl) NewEvidenceObtainMobile(obtainMobileOption *ObtainMobileOption, timeout time.Duration) (string,
	error) {
	if len(obtainMobileOption.ShareUrl) == 0 || len(obtainMobileOption.Title) == 0 {
		return "", errors.New("shareUrl or title 不能为空")
	}
	duration := obtainMobileOption.Duration
	if obtainMobileOption.Duration > 60*60 {
		return "", errors.New("duration 录屏任务不能超过1小时")
	}
	if obtainMobileOption.Duration <= 0 {
		return "", errors.New("duration 录屏任务时间错误")
	}
	param := EvObtainTask{AppId: sdk.appId, WebUrls: obtainMobileOption.ShareUrl, Title: obtainMobileOption.Title,
		Type: 5, ShareUrl: obtainMobileOption.ShareUrl, AppName: obtainMobileOption.AppName, ReqTime: time.Now().Unix(),
		Duration: duration, RepresentAppId: obtainMobileOption.RepresentAppId, Remark: obtainMobileOption.Remark,
		RequestType: "POST", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/mobileobtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.appId, sdk.appKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New(err.Error() + ", requestId:" + cri.RequestId)
	}
	// 打印收到返回的时间
	fmt.Println(fmt.Sprintf("NewEvidenceObtainMobile already send at: %d", time.Now().Unix()))
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var orderNo = txRetDetail.OrderNo
	return orderNo, nil
}

// EvidenceObtainPic 图片取证接口
func (sdk *ZxlImpl) EvidenceObtainPic(webUrls, title, remark string, timeout time.Duration) (string, error) {
	return sdk.evidenceObtainPic(webUrls, title, remark, "", timeout)
}

// RepresentEvidenceObtainPic 代理用户图片取证接口
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

// EvidenceObtainCvd 云桌面取证接口
func (sdk *ZxlImpl) EvidenceObtainCvd(title, remark string, timeout time.Duration) (*TxRetDetail, error) {
	return sdk.NewEvidenceObtainCvd(&ObtainVideoOption{Title: title, Remark: remark, RepresentAppId: "",
		Duration: DEFAULT_VIDEO_DURATION}, timeout)
}

// RepresentEvidenceObtainCvd 代理用户云桌面取证接口
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

// RepresentGetEvidenceStatus 获取取证证书任务状态及结果
func (sdk *ZxlImpl) RepresentGetEvidenceStatus(orderNo, representAppId string, timeout time.Duration) (*EvIdData, error) {
	return sdk.getEvidenceStatus(orderNo, representAppId, timeout)
}

// RepresentZblGetEvidenceStatus 获取取证证书任务状态及结果
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
		Status:      txRetDetail.Status,
		ResultMsg:   txRetDetail.ResultMsg,
		EvidUrl:     txRetDetail.EvIdUrl,
		VoucherUrl:  txRetDetail.VoucherUrl,
		AbnormalTag: 0,
		RequestId:   cri.RequestId,
		Evid:        txRetDetail.Evid,
		EvHash:      txRetDetail.EvHash,
		TxHash:      txRetDetail.TxHash,
		BlockHeight: txRetDetail.BlockHeight,
		StorageTime: txRetDetail.StorageTime,
		Duration:    txRetDetail.Duration,
	}
	// 单独处理异常情况
	if txRetDetail.WebTitle != "" && strings.HasPrefix(txRetDetail.WebTitle, "【异常】") {
		evIdData.AbnormalTag = 1
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
