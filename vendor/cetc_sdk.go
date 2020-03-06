package zxl_go_sdk

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm2"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"io/ioutil"
	"strings"
	"time"
)

type cetcSDKImpl struct {
	AppId  string
	AppKey string
}

// 生成公司钥对
// 返回值  pk公钥（string），sk私钥（string），err错误信息（error）
func (sdk *cetcSDKImpl) GenerateKeyPair() (pk string, sk string, err error) {
	privateKey, _ := sm2.GenerateKey(rand.Reader)
	pk = sm2.EncodePubKey(&privateKey.PublicKey)
	sk = sm2.EncodePrivKey(privateKey)
	return
}

func (sdk *cetcSDKImpl) Sign(prvKey string, data []byte) (string, error) {
	sk, err := sm2.DecodePrivKey(prvKey)
	if err != nil {
		return "", errors.New("Sign (DecodePriKey) error:" + err.Error())
	}
	signBytes, err := sk.Sign(rand.Reader, data, nil)
	if err != nil {
		return "", errors.New("Sign (Sign) error:" + err.Error())
	}
	return hex.EncodeToString(signBytes), nil
}

func (sdk *cetcSDKImpl) Verify(pubKey string, sign string, data []byte) (bool, error) {
	signBytes, err := hex.DecodeString(sign)
	if err != nil {
		return false, errors.New("Verify (DecodeString) error:" + err.Error())
	}
	pk, err := sm2.DecodePubKey(pubKey)
	if err != nil {
		return false, errors.New("Verify (DecodePubKey) error:" + err.Error())
	}
	return pk.Verify(data, signBytes)
}

func (sdk *cetcSDKImpl) EvidenceSave(evHash, extendInfo, sk, pk string) (EvSaveResult, error) {
	uid, err := generateUid()
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc generateUid) error:" + err.Error())
	}
	ed := CetcEvidenceReq{EvId: uid, EvHash: evHash, ExtendInfo: extendInfo,
		Time: time.Now().Format("2006-01-02 15:04:05")}
	rawStr := []byte(strings.Join([]string{sdk.AppId, ed.EvHash, ed.ExtendInfo, ed.EvId, ed.Time}, ","))
	signStr, err := sdk.Sign(sk, rawStr)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc Sign) error:" + err.Error())
	}
	ed.Sign = signStr

	bodyData, _ := json.Marshal(&ed)
	respBytes, err := sendRequest(sdk.AppId, sdk.AppKey, "POST", defConf.ServerAddr+defConf.EvidenceSave, bodyData)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc sendRequest) error:" + err.Error())
	}
	var saveResp CetcEvidenceResp
	err = json.Unmarshal(respBytes, &saveResp)
	if err != nil {
		return nil, errors.New("EvidenceSave (cetc Unmarshal) error:" + err.Error())
	}
	saveResp.EvHash = evHash
	saveResp.EvId = uid
	return &saveResp, nil
}

func (sdk *cetcSDKImpl) CalculateHash(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.New("CalculateHash (ReadFile) error:" + err.Error())
	}
	dataHash := sm3.SumSM3(data)
	return hex.EncodeToString(dataHash), nil
}
