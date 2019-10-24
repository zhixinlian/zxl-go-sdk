package zxl_go_sdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func generateUid() (string,error) {
	tmpUid, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	idStr := strings.ReplaceAll(tmpUid.String(), "-", "")
	return idStr, nil
}

func sendRequest(appId, appKey, method, url string, body []byte) ([]byte, error) {
	var byteReader io.Reader = nil
	if body != nil {
		byteReader = bytes.NewReader(body)
	}

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	cli := http.Client{Transport: tr}

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