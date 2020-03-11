package zxl_go_sdk

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func generateUid() (string, error) {
	tmpUid := uuid.NewV1()
	idStr := strings.ReplaceAll(tmpUid.String(), "-", "")
	return idStr, nil
}
func addTrust(pool *x509.CertPool, path string) {
	aCrt, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
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
