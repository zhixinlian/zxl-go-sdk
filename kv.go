package zxl_go_sdk

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	KvSaveUrl  = "/api/v1/spider/sdk/kv/save"
	KvQueryUrl = "/api/v1/spider/sdk/kv/query"
)

type KvSaveReq struct {
	/**
	 * 字符串 key
	 */
	KvKey string `json:"kvKey"`
	/**
	 * 加密字符串
	 */
	KvValue string `json:"kvValue"`
	/**
	 * 代理用户Id
	 */
	RepresentAppId string `json:"representAppId"`
	/**
	 * 私钥, 如果是代理用户，则传代理用户的私钥
	 */
	Sk  string `json:"sk"`

	AppId          string `json:"appId"`

	KvKeyHash      string `json:"kvKeyHash"`

	KvKeyValueHash string `json:"kvKeyValueHash"`

	Signature      string `json:"signature"`
}

type KvSaveResp struct {
	KvKey     string `json:"kvKey"`
	RequestId string `json:"requestId"`
}

type KvQueryReq struct {
	KvKey     string `json:"kvKey"`
	KvKeyHash string `json:"kvKeyHash"`
}

type KvQueryResp struct {
	KvKey      string `json:"kvKey"`
	KvValue    string `json:"kvValue"`
	CreateTime string `json:"createTime"`
	Status     int    `json:"status"`
	RequestId  string `json:"requestId"`
}

func (zxl *zxlImpl) KvSave(req KvSaveReq, timeout time.Duration) (*KvSaveResp, error) {
	resp := &KvSaveResp{}
	url := defConf.ServerAddr + KvSaveUrl

	if req.KvKey == ""{
		return resp, errors.New("KvKey 不能为空")
	}
	if req.KvValue == "" {
		return resp, errors.New("KvValue 不能为空")
	}
	if req.Sk == "" {
		return resp, errors.New("私钥不能为空")
	}

	req.KvKeyHash = getHash(req.KvKey)
	req.KvKeyValueHash = getHash(req.KvKey+":"+req.KvValue)

	appId := zxl.appId
	if req.RepresentAppId != "" {
		appId = req.RepresentAppId
	}
	req.AppId = zxl.appId

	sign, err := zxl.Sign(req.Sk, []byte(strings.Join([]string{appId, req.KvKey, req.KvValue}, ",")))

	if err != nil {
		return resp, errors.New("get sign failed"+err.Error())
	}
	req.Signature = sign

	params, err := json.Marshal(req)
	if err != nil {
		return resp, errors.New("get req json failed")
	}

	data, cri, err := sendKvRequest(zxl.appId, zxl.appKey, "POST", url, params, timeout)
	if err != nil {
		return resp, errors.New("kv saved failed"+err.Error())
	}
	resp.RequestId = cri.RequestId

	var commonData commonResult
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return nil, errors.New("returned data format error:" + err.Error())
	}
	if commonData.RetCode != 0 {
		return nil, errors.New("http response error info : " + strconv.Itoa(commonData.RetCode) + ","+commonData.RetMsg)
	}

	jsonData, _ := json.Marshal(commonData.Data)

	err = json.Unmarshal(jsonData, resp)
	if err != nil {
		return resp, errors.New("format data error "+ err.Error())
	}

	return resp, nil
}


func (zxl *zxlImpl) KvQuery(req KvQueryReq, timeout time.Duration) (*KvQueryResp, error) {
	resp := &KvQueryResp{}
	url := defConf.ServerAddr + KvQueryUrl

	req.KvKeyHash = getHash(req.KvKey)

	params, err := json.Marshal(req)
	if err != nil {
		return resp, errors.New("get req json failed")
	}

	data, cri, err := sendKvRequest(zxl.appId, zxl.appKey, "POST", url, params, timeout)
	if err != nil {
		return resp, errors.New("kv query failed"+ err.Error())
	}
	resp.RequestId = cri.RequestId

	var commonData commonResult
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return nil, errors.New("returned data format error:" + err.Error())
	}
	if commonData.RetCode != 0 {
		return nil, errors.New("http response error info : " + strconv.Itoa(commonData.RetCode) + ","+commonData.RetMsg)
	}

	jsonData, _ := json.Marshal(commonData.Data)

	err = json.Unmarshal(jsonData, resp)
	if err != nil {
		return resp, errors.New("format data error "+ err.Error())
	}

	return resp, nil
}


func getHash(str string) string{
	return hex.EncodeToString(sm3.SumSM3([]byte(str)))
}


func sendKvRequest(appId, appKey, method, url string, body []byte, timeout time.Duration) ([]byte, *commReqInfo, error) {
	var byteReader io.Reader = nil
	if body != nil {
		byteReader = bytes.NewReader(body)
	}

	cri := commReqInfo{}

	cli := buildHttpClient(defConf.IsProxy, timeout)

	req, err := http.NewRequest(method, url, byteReader)
	if err != nil {
		return nil, &cri, errors.New("NewRequest error:" + err.Error())
	}
	req.Header.Add("appId", appId)

	// 不直接传递 appKey，和时间戳结合使用
	signatureTime := strconv.FormatInt(time.Now().UnixNano() / 1e6, 10)
	signature := hex.EncodeToString(sm3.SumSM3([]byte(appId+","+appKey+","+signatureTime)))

	req.Header.Add("signatureTime", signatureTime)
	req.Header.Add("signature", signature)
	req.Header.Add("content-type", "application/json")
	resp, err := cli.Do(req)
	if resp != nil {
		cri.RequestId = resp.Header.Get(REQUEST_ID)
	}
	if err != nil {
		return nil, &cri, errors.New("cli.Do error:" + err.Error()+ ", requestId:"+ cri.RequestId)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		if resp.StatusCode == 400 || resp.StatusCode == 500 || resp.StatusCode == 401 {
			data, _ := ioutil.ReadAll(resp.Body)
			var commonData CommonRet
			_ = json.Unmarshal(data, &commonData)
			return nil, &cri, errors.New("http response error info : " + commonData.Message)
		}
		return nil, &cri, errors.New("cli.Do error bad status : " + resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, &cri, errors.New("returned data format error:" + string(data))
	}

	return data, &cri, nil
}




