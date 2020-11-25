package zxl_go_sdk

import (
	"bytes"
	"container/list"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func generateUid() (string, error) {
	tmpUid := uuid.NewV1()
	idStr := strings.ReplaceAll(tmpUid.String(), "-", "")
	//idStr := ""
	return idStr, nil
}
func addTrust(pool *x509.CertPool, path string) {
	//aCrt, err := ioutil.ReadFile(path)
	//if err != nil {
	//	fmt.Println("ReadFile err:", err)
	//	return
	//}
	aCrt := []byte("-----BEGIN CERTIFICATE-----\nMIICTzCCAbgCCQDx5kTlifTTZDANBgkqhkiG9w0BAQsFADBsMQswCQYDVQQGEwJj\nbjELMAkGA1UECAwCY2QxCzAJBgNVBAcMAmNkMQwwCgYDVQQKDAN6eGwxDDAKBgNV\nBAsMA3p4bDETMBEGA1UEAwwKYWNjZXNzLmNvbTESMBAGCSqGSIb3DQEJARYDYWFh\nMB4XDTE5MDcyNTAyMTM1OVoXDTI5MDcyMjAyMTM1OVowbDELMAkGA1UEBhMCY24x\nCzAJBgNVBAgMAmNkMQswCQYDVQQHDAJjZDEMMAoGA1UECgwDenhsMQwwCgYDVQQL\nDAN6eGwxEzARBgNVBAMMCmFjY2Vzcy5jb20xEjAQBgkqhkiG9w0BCQEWA2FhYTCB\nnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAwwcACPiwpUsDewlQhElDYLNWPz5B\nZ9DDPW5lUiNJey1oTLf2JP+B4O+BF0H+cg9JXtLuiGfiC6OdJG2e3SjbE/+wSAp0\nSBaeeG6hgKfxJClIzXPNxdayCvxwWf5Z3R3b+XRXceHR/hvHgmlCGTZ+E7Bu5mi4\n2UMPNItP694jcQcCAwEAATANBgkqhkiG9w0BAQsFAAOBgQBWNvgT7ut+lMEBm/Vw\nGCOVdyr7lSCxiS/lg31/zwWWuWtqhdAqljPmaWEtihkNjDVJpHS8ur6yuTwdCNcy\nbPo53O5/bIIIVKf7TMr/neEK7TbuTAf7CA9noMqC7K3vDSC8xdlCSAMO9N96QOk/\nJjCnk7d8N539fQPt80FMKkqhVw==\n-----END CERTIFICATE-----")
	pool.AppendCertsFromPEM(aCrt)
}
func buildHtppClient(isProxy bool, timeout time.Duration) *http.Client {
	pool := x509.NewCertPool()

	//cliCrt, err := tls.LoadX509KeyPair("D:\\certificate\\go\\server.crt", "D:\\certificate\\go\\server.key")
	//if err != nil {
	//	fmt.Println("Loadx509keypair err:", err)
	//	return nil
	//}
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
				RootCAs: pool,
				//Certificates:       []tls.Certificate{cliCrt},
				InsecureSkipVerify: false}, Proxy: proxy}
		client := &http.Client{Transport: transport, Timeout: timeout}
		return client
	} else {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: nil,
				//Certificates:       []tls.Certificate{cliCrt},
				InsecureSkipVerify: false}, Proxy: proxy}
		client := &http.Client{Transport: transport, Timeout: timeout}
		return client
	}

}
func sendRequest(appId, appKey, method, url string, body []byte, timeout time.Duration) ([]byte, error) {
	var byteReader io.Reader = nil
	if body != nil {
		byteReader = bytes.NewReader(body)
	}

	//tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	//
	//cli := http.Client{Transport: tr, Timeout: timeout}
	cli := buildHtppClient(defConf.IsProxy, timeout)

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
	errorCodeLst.PushBack(&TxErrorCodeType{"-101", "此编号未查到对应信息"})
	errorCodeLst.PushBack(&TxErrorCodeType{"3108", "服务不可用,调用后端服务失败"})
	for e := errorCodeLst.Front(); e != nil; e = e.Next() {
		if strings.Index(msg, (e.Value).(*TxErrorCodeType).Code) == -1 {
			continue
		}
		return (e.Value).(*TxErrorCodeType).Msg
	}
	return ""
}

//新增腾讯中间件请求操作发送
func sendTxMidRequest(appId, appKey, method, url string, body []byte, timeout time.Duration) ([]byte, error) {
	var byteReader io.Reader = nil
	if body != nil {
		byteReader = bytes.NewReader(body)
	}

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, DisableKeepAlives: true}

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
		if resp.StatusCode == 400 || resp.StatusCode == 500 || resp.StatusCode == 401 {
			data, _ := ioutil.ReadAll(resp.Body)
			var commonData CommonRet
			_ = json.Unmarshal(data, &commonData)
			return nil, errors.New("http response error info : " + commonData.Message)
		}
		return nil, errors.New("cli.Do error bad status : " + resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var commonData TxRetCommonData
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return nil, errors.New("returned data format error:" + string(data))
	}
	if commonData.RetCode != 0 {
		return nil, errors.New("http response error info : " + retErrMsg(strconv.Itoa(commonData.RetCode)))
	}
	retBytes, _ := json.Marshal(&commonData.Detail)
	return retBytes, nil
}
