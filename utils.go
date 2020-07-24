package zxl_go_sdk

import (
	"bytes"
	"container/list"
	"crypto/tls"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"

	"time"
)

func generateUid() (string, error) {
	tmpUid := uuid.NewV1()
	idStr := strings.ReplaceAll(tmpUid.String(), "-", "")
	return idStr, nil
}

func sendRequest(appId, appKey, method, url string, body []byte, timeout time.Duration) ([]byte, error) {
	var byteReader io.Reader = nil
	if body != nil {
		byteReader = bytes.NewReader(body)
	}

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	cli := http.Client{Transport: tr, Timeout: timeout}

	req, err := http.NewRequest(method, url, byteReader)
	if err != nil {
		return nil, errors.New("NewRequest error:" + err.Error())
	}
	req.Header.Add("appId", appId)
	req.Header.Add("appKey", appKey)
	req.Header.Add("content-type", "application/json")
	resp, err := cli.Do(req)
	if err != nil {
		return nil, errors.New("cli.Do error:" + err.Error())
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 400 || resp.StatusCode == 500 {
			data, _ := ioutil.ReadAll(resp.Body)
			var commonData CommonRet
			_ = json.Unmarshal(data, &commonData)
			return nil, errors.New("http response error info : " + commonData.Message)
		}
		return nil, errors.New("cli.Do error bad status : " + resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	var commonData CommonRet
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return nil, errors.New("returned data format error:" + string(data))
	}
	if commonData.Code != 200 {
		return nil, errors.New("http response error info : " + commonData.Message)
	}

	retBytes, _ := json.Marshal(&commonData.Data)
	return retBytes, nil
}

/**tx Error Type**/
type TxErrorCodeType struct {
	Code string
	Msg  string
}

func retErrMsg(msg string) (errMsg string) {
	errorCodeLst := list.New()
	errorCodeLst.PushBack(&TxErrorCodeType{"461000", "参数格式错误"})
	errorCodeLst.PushBack(&TxErrorCodeType{"461010", "无效的用户appid"})
	errorCodeLst.PushBack(&TxErrorCodeType{"562012", "用户消费套餐已被关闭"})
	errorCodeLst.PushBack(&TxErrorCodeType{"562013", "没有可用的消费套餐"})
	errorCodeLst.PushBack(&TxErrorCodeType{"562014", "内部错误"})
	errorCodeLst.PushBack(&TxErrorCodeType{"563006", "更新消费流水状态失败"})
	for e := errorCodeLst.Front(); e != nil; e = e.Next() {
		if strings.Index(msg, (e.Value).(*TxErrorCodeType).Code) == -1 {
			continue
		}
		return (e.Value).(*TxErrorCodeType).Msg
	}
	return ""
}

//新增腾讯中间件请求操作发送
func sendTxMidRequest(appId, appKey, method, url string, body []byte, timeout time.Duration) (*TxRetDetail, error) {
	var byteReader io.Reader = nil
	if body != nil {
		byteReader = bytes.NewReader(body)
	}

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	cli := http.Client{Transport: tr, Timeout: timeout}

	req, err := http.NewRequest(method, url, byteReader)
	if err != nil {
		return nil, errors.New("NewRequest error:" + err.Error())
	}
	req.Header.Add("appId", appId)
	req.Header.Add("appKey", appKey)
	req.Header.Add("content-type", "application/json")
	resp, err := cli.Do(req)
	if err != nil {
		return nil, errors.New("cli.Do error:" + err.Error())
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 400 || resp.StatusCode == 500 {
			data, _ := ioutil.ReadAll(resp.Body)
			var commonData CommonRet
			_ = json.Unmarshal(data, &commonData)
			return nil, errors.New("http response error info : " + commonData.Message)
		}
		return nil, errors.New("cli.Do error bad status : " + resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	var commonData TxRetCommonData
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return nil, errors.New("returned data format error:" + string(data))
	}
	if commonData.RetCode != 0 {
		return nil, errors.New("http response error info : " + retErrMsg(strconv.Itoa(commonData.RetCode)))
	}
	return &commonData.Detail, nil
}
