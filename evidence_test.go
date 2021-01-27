package zxl_go_sdk

import (
	"encoding/hex"
	"fmt"
	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
	"testing"
	"time"
)

func TestCetc(t *testing.T) {
	/** 2020-07-24 add start */
	zxl, err := NewZxlImpl("200514000200001", "9d7f4ba445a54ed2b041e142d5ea12f3")
	if err != nil {
		fmt.Println("错误")
	}
	//图片取证
	result, err := zxl.EvidenceObtainPic("https://www.baidu.com", "图片", "go_sdk_test", time.Second*2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	//根据orderNo查询取证任务状态及结果
	result1, err := zxl.GetEvidenceStatus("1595578460075584005313706", time.Second*2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result1)
	zxl1, err := NewZxlImpl("200515000110001", "0e4bce1b0ef8471fb9140b849e776f48")
	if err != nil {
		fmt.Println("错误")
	}
	//下发截屏任务
	result2, err := zxl1.ContentCapturePic("https://www.baidu.com", time.Second*2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result2)
	//根据订单号查询截屏任务的状态及结果
	result4, err := zxl1.GetContentStatus("1595578461244000782838012", time.Second*2)
	if err != nil {
		fmt.Println("321321", err.Error())
	}
	fmt.Println(result4)
	/** 2020-07-24 add end */
	//zxl, err := NewZxlImpl("190725000110077", "appTest77")
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//pk, sk, err := zxl.GenerateKeyPair()
	//if err != nil {
	//	t.Error(err.Error())
	//}
	////err = zxl.BindUserCert(pk, sk)
	////if err != nil {
	////	t.Error(err)
	////}
	//err = zxl.UpdateUserCert(pk, sk, 0)
	//if err != nil {
	//	t.Error(err.Error())
	//	return
	//}
	//
	//hashData := sm3.SumSM3([]byte("123123123"))
	//evHash := hex.EncodeToString(hashData)
	//fmt.Print("111111")
	//result, err := zxl.EvidenceSave(evHash, "abc", sk, pk, 0)
	//fmt.Print("2222222")
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//
	//fmt.Println(result.EvId, result.TxHash, result.CreateTime)
	//
	//time.Sleep(time.Second * 10)
	//queryResult1, err := zxl.QueryWithTxHash(result.TxHash,0)
	//if err != nil || queryResult1[0].EvHash != evHash{
	//	t.Error("QueryWithTxHash error")
	//}
	//
	//queryResult2, err := zxl.QueryWithEvId(queryResult1[0].EvId,0)
	//if err != nil || queryResult2[0].EvHash != evHash{
	//	t.Error("QueryWithEvId error")
	//}
	//queryResult3, err := zxl.QueryWithEvHash(evHash, time.Second*2)
	//if err != nil || queryResult3[0].EvHash != evHash{
	//	t.Error("QueryWithEvHash error")
	//}
}

func TestTencent(t *testing.T) {
	zxl, err := NewZxlImpl("190725000110077", "appTest77")
	if err != nil {
		t.Error(err)
	}
	pk, sk, err := zxl.GenerateKeyPair()
	if err != nil {
		t.Error(err)
	}
	err = zxl.UpdateUserCert(pk, sk, time.Second*2)
	if err != nil {
		t.Error(err.Error())
	}

	hashData := sm3.SumSM3([]byte("123123123"))
	evHash := hex.EncodeToString(hashData)
	result, err := zxl.EvidenceSave(evHash, "abc", sk, pk, time.Second*2)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(result.TxHash, result.CreateTime, result.BlockHeight)

	time.Sleep(time.Second * 10)
	queryResult1, err := zxl.QueryWithTxHash(result.GetTxHash(), time.Second*2)
	if err != nil || queryResult1[0].EvHash != evHash {
		t.Error("QueryWithTxHash error")
	}

	queryResult2, err := zxl.QueryWithEvId(result.GetEvId(), time.Second*2)
	if err != nil || queryResult2[0].EvHash != evHash {
		t.Error("QueryWithEvId error")
	}
	queryResult3, err := zxl.QueryWithEvHash(evHash, time.Second*2)
	if err != nil || queryResult3[0].EvHash != evHash {
		t.Error("QueryWithEvHash error")
	}
}

func TestSign(t *testing.T) {
	zxl, err := NewZxlImpl("xxxxxxxxxx0xxxx", "appTestTX")
	if err != nil {
		t.Error(err)
	}
	pk, sk, err := zxl.GenerateKeyPair()
	fmt.Println(pk)
	fmt.Println(sk)
	if err != nil {
		t.Error(err)
	}

	rawData := []byte("abcdefghijklmnopqrstuvwxyz")
	signStr, err := zxl.Sign(sk, rawData)
	fmt.Println(signStr, err)

	success, err := zxl.Verify(pk, signStr, rawData)
	fmt.Println(success, err)
}

func TestCipher(t *testing.T) {
	zxl, err := NewZxlImpl("xxxxxxxxxx0xxxx", "appTestTX")
	if err != nil {
		t.Error(err)
	}
	rawData := []byte("123abc")
	pwd := "111"
	encryptedData, err := zxl.EncryptData(pwd, rawData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(encryptedData, err)
	decryptedData, err := zxl.DecryptData("asdf", encryptedData)
	fmt.Println(string(decryptedData), err)
}

func TestHash(t *testing.T) {
	zxl, err := NewZxlImpl("xxxxxxxxxx1xxxx", "appTestTX")
	if err != nil {
		t.Error(err)
	}
	hashStr, err := zxl.CalculateHash("G:\\channel.zip")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(hashStr)
}
func TestStrHash(t *testing.T) {
	zxl, err := NewZxlImpl("190725000110077", "appTestTX")
	hashStr, err := zxl.CalculateStrHash("fwejfoiwfjoweifjowf")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(hashStr)
}

/**代理用户注册的相关调用示例*/
func TestAgentUser(t *testing.T) {
	//var filePath = "E:\\至信链\\sdk接入文档\\inner.png"
	zxl, err := NewZxlImpl("201010000210001", "9487fe0f7d1f436fb0ef62ce6608c236")
	if err != nil {
		fmt.Println(err)
	}
	//var user = AgentUser{RepresentEmail: "990991011@qq.com", Pwd: "w836546028", CardFrontFile: filePath, LetterFile: filePath,
	//	CardBackendFile: filePath, LicenseFile: filePath, Representative: "李艳", EpName: "山东拼多多供应链管理有限公司", CreditCode: "91370105MA3P4F040J",
	//	Idcard: "511321198912037013", Contact: "刘飞", Title: "cto", Mobile: "18280097243", Category: 1}
	//registerFlag, err := zxl.RegisterUser(user, 0)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(registerFlag)
	//fmt.Println(zxl.SelectEpInfo("990991011@qq.com",0))
	///**注册审核完毕的appId及appKey -- 201201000110001 1d8b89be2eab4d4e8c9ddbe5843ca983**/
	////公私钥绑定
	//fmt.Println(zxl.BindRepresentUserCert("201201000110001","1d8b89be2eab4d4e8c9ddbe5843ca983",
	//	"04c0eb7f2b60a8752c8852ac255c966fee8aa5accc7a1f74b55d50bbc0faf3d4dbd98144624b5f002e0cd88662337e6e0a90f4875487fef04750b1124fc03cae61",
	//	"8b387ab63c4238758b762ae5269bb05173312296cb473e9b804b0bd274a37e61"))
	//公私钥更新
	//fmt.Println(zxl.UpdateRepresentUserCert("201201000110001","1d8b89be2eab4d4e8c9ddbe5843ca983",
	//	"04fe7cc425457346d563a1ecde041ae0eabf4abf8e19a20d7087bab1a60be5dcad328049398c9ea0454547c15d54a0053b014da11b87d90aa3ab5fa90dd9ca11f4",
	//	"8e7b64d470b04000a24be597496449e4573b0e566b8df30420f4c8e5b9821103"))
	//代理用户上链
	evHash, err := zxl.CalculateStrHash("test上链")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(zxl.RepresentSave(evHash, "测试上链", "8b387ab63c4238758b762ae5269bb05173312296cb473e9b804b0bd274a37e61", "201201000110001", 0))
	//fmt.Println(zxl.EvidenceSave(evHash,"ceshi","ba35a87f5550d0319c2518db920c40f0bb1b9df6e5aab2142c0c04203fbe8d09",
	//	"04e347499fc53813a00053612d77e4c2229a586e04a9b384c20be20a662f95b0a94a22ff3b417c779df27243f847a32a7a0c0b108a6745ae56ac8b3fdea0d36683",0))
}
