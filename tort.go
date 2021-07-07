package zxl_go_sdk

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zhixinlian/zxl-go-sdk/constants"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	SUBMIT_TORT_SEARCH_URL = "/api/v2/ipm/tort/create"
	QUERY_TORT_RESULT_URL = "/api/v2/ipm/watch_tort/get_task_result/"
)


/**
 * 提交侵权比对请求
 */
type Tort struct {
	Url         string               `json:"url"`
	Title       string               `json:"title"`
	Keyword     string               `json:"keyword"`
	Type        constants.TortType   `json:"type"`
	Source      constants.TortSource `json:"source"`
	StartDate   string               `json:"startDate"`
	EndDate     string               `json:"endDate"`
}

type TortResp struct {
	TaskId      string `json:"taskId"`
	RequestId   string
}

/**
 * 查询侵权结果
 */
type TortQuery struct {
	TaskId      string `json:"taskId"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"int"`
}

type TortQueryResp struct {
	ClueList  []ClueData `json:"clueList"`
	Count     int        `json:"count"`
	Status    int        `json:"status"`
	RequestId string
}

/**
 * 侵权线索
 */
type ClueData struct {
	ClueId    int    `json:"clueId"`
	PlayUrl   string `json:"playUrl"`
	Title     string `json:"title"`
	PubTime   string `json:"pubTime"`
	Platform  string `json:"platform"`
	Author    string `json:"author"`
	FindTime  string `json:"findTime"`
}


type commonResult struct {
	Data interface{} `json:"data"`
	RetCode int `json:"retCode"`
	RetMsg string `json:"retMsg"`
}


func (zxl *zxlImpl) SubmitTortTask(tort Tort, timeout time.Duration) (TortResp, error) {
	var resp TortResp

	if tort.Url == "" || tort.Title == "" || tort.EndDate == "" || tort.Type == 0 || tort.Source == 0 {
		return resp, errors.New("参数错误")
	}

	var startDate time.Time
	if tort.StartDate == "" {
		startDate = time.Now()
	} else {
		startDate, _ = time.Parse("2006-01-02", tort.StartDate)
	}
	endDate, _ := time.Parse("2006-01-02", tort.EndDate)

	if startDate.After(endDate) {
		return resp, errors.New("时间参数错误")
	}

	paramBytes, _ := json.Marshal(tort)
	retBytes, cri, err := sendTortRequest(zxl.appId, zxl.appKey, "POST", defConf.ServerAddr + SUBMIT_TORT_SEARCH_URL, paramBytes,timeout)

	if err != nil {
		return resp, errors.New("提交侵权比对请求错误：" + err.Error()+ ", requestId:"+ cri.RequestId)
	}

	err = json.Unmarshal(retBytes, &resp)

	if err != nil {
		return resp, errors.New("解析结果出错")
	}
	resp.RequestId = cri.RequestId
	return resp, nil
}

func (zxl *zxlImpl) QueryTortTaskResult(tortQuery TortQuery, timeout time.Duration) (TortQueryResp, error) {
	var resp TortQueryResp

	url := QUERY_TORT_RESULT_URL + tortQuery.TaskId + fmt.Sprintf("?offset=%v&limit=%v", tortQuery.Offset, tortQuery.Limit)

	retBytes, cri, err := sendTortRequest(zxl.appId, zxl.appKey, "GET", defConf.ServerAddr+url, []byte(""), timeout)

	if err != nil {
		return resp, errors.New("提交侵权查询请求错误："+err.Error()+ ", requestId:"+ cri.RequestId)
	}

	var commonData commonResult
	err = json.Unmarshal(retBytes, &commonData)
	if err != nil {
		return resp, errors.New("解析结果错误: "+err.Error()+ ", requestId:"+ cri.RequestId)
	}

	ret, _ := json.Marshal(&commonData.Data)

	err = json.Unmarshal(ret, &resp)
	if err != nil {
		return resp, errors.New("解析结果错误: "+err.Error()+ ", requestId:"+ cri.RequestId)
	}
	resp.RequestId = cri.RequestId

	return resp, nil
}

func sendTortRequest(appId, appKey, method, url string, body []byte, timeout time.Duration) ([]byte, *commReqInfo, error) {
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
		return nil, &cri, errors.New("cli.Do error:" + err.Error())
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

	var commonData commonResult
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return nil, &cri, errors.New("returned data format error:" + string(data))
	}
	if commonData.RetCode != 0 {
		return nil, &cri, errors.New("http response error info : " + retErrMsg(strconv.Itoa(commonData.RetCode)))
	}
	// 确权结果返回的数据不一致，这里需要重新返回原数据
	return data, &cri, nil
}

