package zxl_go_sdk

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"github.com/zhixinlian/zxl-go-sdk/trustsql-sdk/encoding"
	"github.com/zhixinlian/zxl-go-sdk/trustsql-sdk/util/byteutils"
	"io/ioutil"
	"time"
)

type trustSDKImpl struct {
	AppId  string
	AppKey string
}

// 生成公司钥对
// 返回值  pk公钥（string），sk私钥（string），err错误信息（error）
func (sdk *trustSDKImpl) GenerateKeyPair() (pk string, sk string, err error) {
	var prvKey *btcec.PrivateKey
	prvKey, err = btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		err = errors.New("GenerateKeyPair error:" + err.Error())
		return
	}
	pubKey := prvKey.PubKey()
	sk = base64.StdEncoding.EncodeToString(prvKey.Serialize())
	pk = base64.StdEncoding.EncodeToString(pubKey.SerializeCompressed())
	return
}

func (sdk *trustSDKImpl) GeneratePubKeyFromPrvKey(prvkey string) (string, error) {
	prvkeyBytes, err := base64.StdEncoding.DecodeString(prvkey)
	if err != nil {
		return "", err
	}
	_, pubkey := btcec.PrivKeyFromBytes(btcec.S256(), prvkeyBytes)
	return base64.StdEncoding.EncodeToString(pubkey.SerializeCompressed()), nil
}

func (sdk *trustSDKImpl) Sign(prvKey string, data []byte) (string, error) {
	prvkeyBytes, err := base64.StdEncoding.DecodeString(prvKey)
	if err != nil {
		return "", err
	}

	prv, _ := btcec.PrivKeyFromBytes(btcec.S256(), prvkeyBytes)
	var datahash []byte

	hasher := sha256.New()
	hasher.Write(data)
	datahash = hasher.Sum(nil)

	sign, err := prv.Sign(datahash)
	if err != nil {
		return "", err
	}
	return byteutils.Hex(sign.Serialize()), err
}

func (sdk *trustSDKImpl) OnlySign(prvKey string, data []byte) (string, error) {
	prvkeyBytes, err := base64.StdEncoding.DecodeString(prvKey)
	if err != nil {
		return "", err
	}

	prv, _ := btcec.PrivKeyFromBytes(btcec.S256(), prvkeyBytes)
	var datahash []byte

	datahash = data

	sign, err := prv.Sign(datahash)
	if err != nil {
		return "", err
	}
	return byteutils.Hex(sign.Serialize()), err
}

func (sdk *trustSDKImpl) Verify(pubKey string, sign string, data []byte) (bool, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return false, err
	}
	pub, err := btcec.ParsePubKey(pubKeyBytes, btcec.S256())
	var datahash []byte

	hasher := sha256.New()
	hasher.Write([]byte(data))
	datahash = hasher.Sum(nil)

	signBytes, err := byteutils.FromHex(sign)
	if err != nil {
		return false, err
	}
	signature, err := btcec.ParseDERSignature(signBytes, btcec.S256())
	if err != nil {
		return false, err
	}
	return signature.Verify(datahash, pub), nil
}

func (sdk *trustSDKImpl) EvidenceSave(evHash, extendInfo, sk, pk string, timeout time.Duration) (*EvSaveResult, error) {
	uid, err := generateUid()
	if err != nil {
		return nil, errors.New("EvidenceSave (generateUid) error:" + err.Error())
	}
	params := SystemParams{MchType: 1, MchId: sdk.AppId, SignType: "ECDSA", TimeStamp: time.Now().Unix(), Version: "1.0"}
	content := ApplyContent{Content: evHash, EvId: uid, ExtendInfo: extendInfo}
	applyData := EvidenceApplyReq{SysParams: params, Content: content, PublicKey: pk}
	applyBytes, _ := json.Marshal(&applyData)
	applySign, err := sdk.Sign(sk, applyBytes)
	if err != nil {
		return nil, errors.New("EvidenceSave (Sign) error:" + err.Error())
	}
	applyBodyBytes, _ := json.Marshal(&TencentEvidenceReq{BodyData: string(applyBytes), BodySign: applySign})
	applyRetBytes, cri, err := sendRequest(sdk.AppId, sdk.AppKey, "POST",
		defConf.ServerAddr+defConf.EvidenceApply, applyBodyBytes, timeout)
	if err != nil {
		return nil, errors.New("EvidenceSave (sendRequest) error:" + err.Error()+ ", requestId:"+ cri.RequestId)
	}
	var applyResp EvidenceApplyResp
	err = json.Unmarshal(applyRetBytes, &applyResp)
	if err != nil {
		return nil, errors.New("EvidenceSave (Unmarshal) error:" + err.Error())
	}

	decodeData, _ := hex.DecodeString(applyResp.ToBeSign)
	submitSign, err := sdk.OnlySign(sk, decodeData)
	if err != nil {
		return nil, errors.New("EvidenceSave (Sign2) error:" + err.Error())
	}
	params.TimeStamp = time.Now().Unix()
	submitData := EvidenceSubmitReq{SysParams: params, Sign: submitSign,
		Session0: applyResp.Session0, Session1: applyResp.Session1}
	submitBytes, _ := json.Marshal(&submitData)
	submitSign2, err := sdk.Sign(sk, submitBytes)
	if err != nil {
		return nil, errors.New("EvidenceSave (Sign3) error:" + err.Error())
	}
	submitBodyBytes, _ := json.Marshal(&TencentEvidenceReq{BodyData: string(submitBytes), BodySign: submitSign2})
	submitRetBytes, cri, err := sendRequest(sdk.AppId, sdk.AppKey, "POST",
		defConf.ServerAddr+defConf.EvidenceSubmit, submitBodyBytes, timeout)
	if err != nil {
		return nil, errors.New("EvidenceSave (sendRequest2) error:" + err.Error()+ ", requestId:"+ cri.RequestId)
	}
	var tencentResp EvSaveResult
	err = json.Unmarshal(submitRetBytes, &tencentResp)
	if err != nil {
		return nil, errors.New("EvidenceSave (Unmarshal2) error:" + err.Error())
	}
	tencentResp.EvId = uid
	tencentResp.EvHash = evHash
	tencentResp.RequestId = cri.RequestId
	return &tencentResp, nil
}

func (sdk *trustSDKImpl) CalculateHash(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.New("CalculateHash (ReadFile) error:" + err.Error())
	}
	dataHash := sm3.SumSM3(data)
	return hex.EncodeToString(dataHash), nil
}

//下发截屏任务到取证工具服务
func (sdk *trustSDKImpl) ContentCaptureVideo(webUrls string, timeout time.Duration) (string, error) {
	if len(webUrls) == 0 {
		return "", errors.New("webUrls 不能为空")
	}
	param := EvObtainTask{WebUrls: webUrls, Type: 2, AppId: sdk.AppId, RequestType: "POST", RedirectUrl: "zhixin-api/v2/screenshot/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	applyRetBytes, cri, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST",
		defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New("下发任务异常>>error:" + err.Error()+ ", requestId:"+ cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(applyRetBytes, &txRetDetail)
	var applyResp = txRetDetail.OrderNo
	fmt.Print("输出当前结果" + applyResp)
	return applyResp, nil
}

//下发录屏任务到取证工具服务
func (sdk *trustSDKImpl) ContentCapturePic(webUrls string, timeout time.Duration) (string, error) {
	if len(webUrls) == 0 {
		return "", errors.New("webUrls 不能为空")
	}
	param := EvObtainTask{WebUrls: webUrls, Type: 1, AppId: sdk.AppId, RequestType: "POST", RedirectUrl: "zhixin-api/v2/screenshot/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New("下发任务异常>>error:" + err.Error()+ ", requestId:"+ cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var retResp = txRetDetail.OrderNo
	fmt.Println("输出当前结果:" + retResp)
	return retResp, nil
}
func (sdk *trustSDKImpl) getContentStatus(orderNo string, timeout time.Duration) (*TaskEvData, error) {
	if len(orderNo) == 0 {
		return nil, errors.New("orderNo 不能为空")
	}
	param := EvObtainTask{AppId: sdk.AppId, OrderNo: orderNo, RequestType: "GET", RedirectUrl: "zhixin-api/v2/screenshot/evobtain/evidinfo"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return nil, errors.New(err.Error()+ ", requestId:"+ cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var taskEvData = TaskEvData{Hash: txRetDetail.Hash, StatusMsg: txRetDetail.StatusMsg, Status: txRetDetail.Status, Url: txRetDetail.Url}
	return &taskEvData, nil
}

//视频取证接口
func (sdk *trustSDKImpl) evidenceObtainVideo(webUrls, title, remark string, timeout time.Duration) (string, error) {
	if len(webUrls) == 0 || len(title) == 0 {
		return "", errors.New("webUrls or title 不能为空")
	}
	param := EvObtainTask{AppId: sdk.AppId, WebUrls: webUrls, Title: title, Type: 2, Remark: remark, RequestType: "POST", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New(err.Error()+ ", requestId:"+ cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var orderNo = txRetDetail.OrderNo
	return orderNo, nil
}

//图片取证接口
func (sdk *trustSDKImpl) evidenceObtainPic(webUrls, title, remark string, timeout time.Duration) (string, error) {
	if len(webUrls) == 0 || len(title) == 0 {
		return "", errors.New("webUrls or title 不能为空")
	}
	param := EvObtainTask{AppId: sdk.AppId, WebUrls: webUrls, Title: title, Type: 1, Remark: remark, RequestType: "POST", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/obtain"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return "", errors.New(err.Error()+ ", requestId:"+ cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var orderNo = txRetDetail.OrderNo
	return orderNo, nil
}

//获取取证证书任务状态及结果
func (sdk *trustSDKImpl) getEvidenceStatus(orderNo string, timeout time.Duration) (*EvIdData, error) {
	if len(orderNo) == 0 {
		return nil, errors.New("orderNo 不能为空")
	}
	param := EvObtainTask{AppId: sdk.AppId, OrderNo: orderNo, RequestType: "GET", RedirectUrl: "sdk/zhixin-api/v2/busi/evobtain/evidinfo"}
	paramBytes, _ := json.Marshal(&param)
	sendRetBytes, cri, err := sendTxMidRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return nil, errors.New(err.Error()+ ", requestId:"+ cri.RequestId)
	}
	var txRetDetail TxRetDetail
	json.Unmarshal(sendRetBytes, &txRetDetail)
	var evIdData = EvIdData{
		Status: txRetDetail.Status,
		EvidUrl: txRetDetail.EvIdUrl,
		VoucherUrl: txRetDetail.VoucherUrl,
		RequestId: cri.RequestId,
	}
	return &evIdData, nil
}
