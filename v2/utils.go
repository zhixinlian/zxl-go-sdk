package zxl_go_sdk

import (
	"bytes"
	"container/list"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/zhixinlian/zxl-go-sdk/v2/sm/sm3"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/**
 * 请求中携带的公共信息
 */
type commReqInfo struct {
	RequestId string
}

const REQUEST_ID = "traceId"

func generateUid() (string, error) {
	tmpUid, _ := uuid.NewUUID()
	idStr := strings.ReplaceAll(tmpUid.String(), "-", "")
	return idStr, nil
}
func addTrust(pool *x509.CertPool, path string) {
	aCrt := []byte("-----BEGIN CERTIFICATE-----\nMIICTzCCAbgCCQDx5kTlifTTZDANBgkqhkiG9w0BAQsFADBsMQswCQYDVQQGEwJj\nbjELMAkGA1UECAwCY2QxCzAJBgNVBAcMAmNkMQwwCgYDVQQKDAN6eGwxDDAKBgNV\nBAsMA3p4bDETMBEGA1UEAwwKYWNjZXNzLmNvbTESMBAGCSqGSIb3DQEJARYDYWFh\nMB4XDTE5MDcyNTAyMTM1OVoXDTI5MDcyMjAyMTM1OVowbDELMAkGA1UEBhMCY24x\nCzAJBgNVBAgMAmNkMQswCQYDVQQHDAJjZDEMMAoGA1UECgwDenhsMQwwCgYDVQQL\nDAN6eGwxEzARBgNVBAMMCmFjY2Vzcy5jb20xEjAQBgkqhkiG9w0BCQEWA2FhYTCB\nnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAwwcACPiwpUsDewlQhElDYLNWPz5B\nZ9DDPW5lUiNJey1oTLf2JP+B4O+BF0H+cg9JXtLuiGfiC6OdJG2e3SjbE/+wSAp0\nSBaeeG6hgKfxJClIzXPNxdayCvxwWf5Z3R3b+XRXceHR/hvHgmlCGTZ+E7Bu5mi4\n2UMPNItP694jcQcCAwEAATANBgkqhkiG9w0BAQsFAAOBgQBWNvgT7ut+lMEBm/Vw\nGCOVdyr7lSCxiS/lg31/zwWWuWtqhdAqljPmaWEtihkNjDVJpHS8ur6yuTwdCNcy\nbPo53O5/bIIIVKf7TMr/neEK7TbuTAf7CA9noMqC7K3vDSC8xdlCSAMO9N96QOk/\nJjCnk7d8N539fQPt80FMKkqhVw==\n-----END CERTIFICATE-----")
	pool.AppendCertsFromPEM(aCrt)
}
func buildHttpClient(isProxy bool, timeout time.Duration) *http.Client {
	pool := x509.NewCertPool()

	var proxy func(*http.Request) (*url.URL, error) = nil
	if isProxy {
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://" + defConf.ProxyHost + ":" + defConf.ProxyPort)
		}
	}
	if defConf.ServerCrtPath != "" {
		addTrust(pool, defConf.ServerCrtPath)
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            pool,
				InsecureSkipVerify: false}, Proxy: proxy}
		client := &http.Client{Transport: transport, Timeout: timeout}
		return client
	} else {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            nil,
				InsecureSkipVerify: false}, Proxy: proxy}
		client := &http.Client{Transport: transport, Timeout: timeout}
		return client
	}

}
func sendRequest(appId, appKey, method, url string, body []byte, timeout time.Duration) ([]byte, *commReqInfo, error) {
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
	signatureTime := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	signature := hex.EncodeToString(sm3.SumSM3([]byte(appId + "," + appKey + "," + signatureTime)))

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
		if resp.StatusCode == 400 || resp.StatusCode == 500 {
			data, _ := ioutil.ReadAll(resp.Body)
			var commonData CommonRet
			_ = json.Unmarshal(data, &commonData)
			return nil, &cri, errors.New("http response error info : " + commonData.Message)
		}
		return nil, &cri, errors.New("cli.Do error bad status : " + resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	var commonData CommonRet
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return nil, &cri, errors.New("returned data format error:" + string(data))
	}
	if commonData.Code != 200 {
		return nil, &cri, errors.New("http response error info : " + commonData.Message)
	}

	retBytes, _ := json.Marshal(&commonData.Data)
	return retBytes, &cri, nil
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
	errorCodeLst.PushBack(&TxErrorCodeType{"-101", "此编号未查到对应信息"})
	errorCodeLst.PushBack(&TxErrorCodeType{"3108", "服务不可用,调用后端服务失败"})
	errorCodeLst.PushBack(&TxErrorCodeType{"3116", "调用取证工具失败"})
	for e := errorCodeLst.Front(); e != nil; e = e.Next() {
		if strings.Index(msg, (e.Value).(*TxErrorCodeType).Code) == -1 {
			continue
		}
		return (e.Value).(*TxErrorCodeType).Msg
	}
	return msg
}

//新增腾讯中间件请求操作发送
func sendTxMidRequest(appId, appKey, method, url string, body []byte, timeout time.Duration) ([]byte, *commReqInfo, error) {
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
	signatureTime := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	signature := hex.EncodeToString(sm3.SumSM3([]byte(appId + "," + appKey + "," + signatureTime)))

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

	var commonData TxRetCommonData
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return nil, &cri, errors.New("returned data format error:" + string(data))
	}
	if commonData.RetCode != 0 {
		return nil, &cri, errors.New("http response error info : " + retErrMsg(strconv.Itoa(commonData.
			RetCode)) + "	" + commonData.RetMsg)
	}
	retBytes, _ := json.Marshal(&commonData.Detail)
	return retBytes, &cri, nil
}

func isInnerIpFromUrl(originUrl string) bool {
	u, err := url.Parse(originUrl)
	if err != nil {
		return true
	}

	h := strings.Split(u.Host, ":")

	// 先检查是否是 ip，只对内网的 ip 进行拦截，如果是域名， 直接放行
	if !checkIp(h[0]) {
		return false
	}
	addr, err := net.ResolveIPAddr("ip", h[0])
	if err != nil {
		return true
	}

	if isInnerIp(addr.IP.String()) {
		return true
	}

	return false
}

func isInnerIp(ipStr string) bool {
	inputIpNum := inetAton(ipStr)
	innerIpA := inetAton("10.255.255.255")
	innerIpB := inetAton("172.16.255.255")
	innerIpC := inetAton("192.168.255.255")
	innerIpD := inetAton("100.64.255.255")
	innerIpF := inetAton("127.255.255.255")

	return inputIpNum>>24 == innerIpA>>24 || inputIpNum>>20 == innerIpB>>20 ||
		inputIpNum>>16 == innerIpC>>16 || inputIpNum>>22 == innerIpD>>22 ||
		inputIpNum>>24 == innerIpF>>24
}

func checkIp(ipStr string) bool {
	address := net.ParseIP(ipStr)
	if address == nil {
		return false
	} else {
		return true
	}
}

func inetAton(ipStr string) int64 {
	bits := strings.Split(ipStr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}
