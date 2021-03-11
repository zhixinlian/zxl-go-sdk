package zxl_go_sdk

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/zhixinlian/zxl-go-sdk/constants"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)


const(
	SUBMIT_DCI_CLAIM = "sdk/zhixin-api/dci/dci_claim"
	QUERY_DCI_RESULT = "sdk/zhixin-api/dci/get_dci_claim_result"
)

/**
确权相关接口
*/

type DciClaim struct {
	AgentAppId        string `json:"agentAppId"`
	DciName           string `json:"dciName"`
	ProposerEmail     string                       `json:"proposerEmail"`
	DciType           constants.DciType            `json:"dciType"`
	DciCreateProperty constants.DciCreateProperty  `json:"dciCreateProperty"`
	DciCreateTime     string                       `json:"dciCreateTime"`
	DciUrl            string                       `json:"dciUrl"`
	DciHash           string                       `json:"dciHash"`
	Signature         string                       `json:"signature"`
	AuthorList        []DciAuthor                  `json:"authorList"`
	RightInfoList     []DciRight                   `json:"rightInfoList"`
	RightSignatureDic map[string]map[string]string `json:"rightSignatureDic"`
	RequestType string `json:"requestType"`
	RedirectUrl string `json:"redirectUrl"`
	RepresentAppId string `json:"-"`
}

type DciClaimResp struct {
	TaskId string `json:"taskId"`
}

type dciClaimResult struct {
	Data interface{} `json:"data"`
	RetCode int `json:"retCode"`
	RetMsg string `json:"retMsg"`
}

type DciClaimQuery struct {
	TaskId      string `json:"taskId"`
	RequestType string `json:"requestType"`
	RedirectUrl string `json:"redirectUrl"`
}

type DciClaimQueryResp struct {
	Status   int `json:"status"`
	TortSearchList interface{} `json:"tortSearchList"`
	RecordTimestamp int `json:"recordTimestamp"`
	DciId         string `json:"dciId"`
	Url           string `json:"url"`
	Msg           string `json:"msg"`
}

type DciAuthor struct {
	AuthorIdCard string `json:"authorIdCard"`
	AuthorName string `json:"authorName"`
	AuthorType constants.AuthorType `json:"authorType"`
}

type DciRight struct {
	DciKey string `json:"dciKey"`
	Type constants.RightType `json:"type"`
	RighterInfoList []DciRighter `json:"righterInfoList"`
}

type DciRighter struct {
	RighterEmail string `json:"righterEmail"`
	RighterGainedWay constants.GainedWay `json:"righterGainedWay"`
	RighterIdCard string `json:"righterIdCard"`
	RighterName  string `json:"righterName"`
	RighterType  constants.RighterType `json:"righterType"`
	Sk string `json:"-"`
}

/**
提交确权申请
 */
func (zxl *zxlImpl) SubmitDciClaim(dci DciClaim, sk string, timeout time.Duration) (DciClaimResp, error) {
	var resp DciClaimResp
	if len(dci.AuthorList) > 5 {
		return resp, errors.New("作者数量超限额错误")
	}

	u, err := net.ResolveIPAddr("ip",dci.DciUrl)
	if err == nil {
		if IsInnerIp(u.IP.String()) {
			return resp, errors.New("url不合规，请检查")
		}
	} else {
		return resp, errors.New("url不合规，请检查")
	}

	content, err := zxl.getContent(dci.DciUrl)
	if err != nil {
		return resp, err
	}
	dciHash := hex.EncodeToString(sm3.SumSM3([]byte(content)))
	dci.DciHash = dciHash

	authorJson, err := json.Marshal(dci.AuthorList)
	if err != nil{
		return resp, err
	}

	signStr := strings.Join([]string{dci.ProposerEmail,
		dci.DciName,
		string(dci.DciType),
		dci.DciCreateTime,
		string(dci.DciCreateProperty),
		dci.DciUrl,
		dci.DciHash,
		string(authorJson)}, "_")


	signature, err := zxl.Sign(sk, []byte(signStr))
	if err != nil {
		return resp, nil
	}
	dci.Signature = signature
	rightSignatureDic := make(map[string]map[string]string)

	for i, right := range dci.RightInfoList {

		if right.Type != constants.RIGHT_TYPE_ALL {
			return resp, errors.New("right type is incorrect")
		}

		right.DciKey = dci.DciHash
		dci.RightInfoList[i] = right

		signMap := make(map[string]string)
		righterInfoJson, err := json.Marshal(right.RighterInfoList)
		if err != nil {
			return resp, nil
		}
		rightSignStr := strings.Join([]string{right.DciKey, string(right.Type), string(righterInfoJson)}, "_")

		for _, righter := range right.RighterInfoList {
			sign, err := zxl.Sign(righter.Sk, []byte(rightSignStr))
			if err != nil {
				return resp, err
			}
			signMap[righter.RighterEmail] = sign
		}
		rightSignatureDic[string(right.Type)] = signMap
	}

	// 如果是代理用户确权，就需要填充 AgentAppId 字段
	if dci.RepresentAppId != "" {
		dci.AgentAppId = zxl.appId
	}

	dci.RightSignatureDic = rightSignatureDic
	dci.RequestType = "POST"
	dci.RedirectUrl = SUBMIT_DCI_CLAIM

	paramBytes, _ := json.Marshal(dci)

	sendRetBytes, err := sendDciRequest(zxl.appId, zxl.appKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return resp, errors.New("提交确权结果错误：" + err.Error())
	}
	json.Unmarshal(sendRetBytes, &resp)

	return resp, nil
}

/**
查询确权结果
 */
func (zxl *zxlImpl) QueryDciClaimResult(dciQuery DciClaimQuery, timeout time.Duration) (DciClaimQueryResp, error) {
	var resp DciClaimQueryResp

	dciQuery.RequestType = "GET"
	dciQuery.RedirectUrl = QUERY_DCI_RESULT
	paramBytes, _ := json.Marshal(dciQuery)

	sendRetBytes, err := sendDciRequest(zxl.appId, zxl.appKey, "POST", defConf.ServerAddr+defConf.ContentCapture, paramBytes, timeout)
	if err != nil {
		return resp, errors.New("查询确权结果错误：" + err.Error())
	}

	json.Unmarshal(sendRetBytes, &resp)
	return resp, nil
}

func (zxl *zxlImpl) getContent(url string) (string, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(body), nil
}

func sendDciRequest(appId, appKey, method, url string, body []byte, timeout time.Duration) ([]byte, error) {
	var byteReader io.Reader = nil
	if body != nil {
		byteReader = bytes.NewReader(body)
	}

	cli := buildHttpClient(defConf.IsProxy, timeout)
	
	req, err := http.NewRequest(method, url, byteReader)
	if err != nil {
		return nil, errors.New("NewRequest error:" + err.Error())
	}
	req.Header.Add("appId", appId)

	// 不直接传递 appKey，和时间戳结合使用
	signatureTime := strconv.FormatInt(time.Now().UnixNano() / 1e6, 10)
	signature := hex.EncodeToString(sm3.SumSM3([]byte(appId+","+appKey+","+signatureTime)))

	req.Header.Add("signatureTime", signatureTime)
	req.Header.Add("signature", signature)
	req.Header.Add("content-type", "application/json")
	resp, err := cli.Do(req)
	if err != nil {
		return nil, errors.New("cli.Do error:" + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		if resp.StatusCode == 400 || resp.StatusCode == 500 || resp.StatusCode == 401 {
			data, _ := ioutil.ReadAll(resp.Body)
			var commonData CommonRet
			_ = json.Unmarshal(data, &commonData)
			return nil, errors.New("http response error info : " + commonData.Message)
		}
		return nil, errors.New("cli.Do error bad status : " + resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)

	var commonData dciClaimResult
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return nil, errors.New("returned data format error:" + string(data))
	}
	if commonData.RetCode != 0 {
		return nil, errors.New("http response error info : " + retErrMsg(strconv.Itoa(commonData.RetCode)))
	}
	retBytes, _ := json.Marshal(&commonData.Data)
	return retBytes, nil
}
