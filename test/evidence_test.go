package test

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	zxl_go_sdk "github.com/zhixinlian/zxl-go-sdk"

	"github.com/zhixinlian/zxl-go-sdk/sm/sm3"
)

func TestCetc(t *testing.T) {
	/** 2020-07-24 add start */
	zxl, err := zxl_go_sdk.NewZxlImpl("", "")
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
	result1, err := zxl.GetEvidenceStatus("", time.Second*2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result1)
	zxl1, err := zxl_go_sdk.NewZxlImpl("", "")
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
	result4, err := zxl1.GetContentStatus("", time.Second*2)
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
	zxl, err := zxl_go_sdk.NewZxlImpl("", "")
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
	zxl, err := zxl_go_sdk.NewZxlImpl("xxxxxxxxxx0xxxx", "appTestTX")
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
	zxl, err := zxl_go_sdk.NewZxlImpl("xxxxxxxxxx0xxxx", "appTestTX")
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
	zxl, err := zxl_go_sdk.NewZxlImpl("xxxxxxxxxx1xxxx", "appTestTX")
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
	zxl, err := zxl_go_sdk.NewZxlImpl("", "appTestTX")
	hashStr, err := zxl.CalculateStrHash("fwejfoiwfjoweifjowf")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(hashStr)
}

/**代理用户注册的相关调用示例*/
func TestAgentUser(t *testing.T) {
	zxl, err := zxl_go_sdk.NewZxlImpl("", "")
	if err != nil {
		fmt.Println(err)
	}
	evHash, err := zxl.CalculateStrHash("test上链")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(zxl.RepresentSave(evHash, "测试上链", "", "", 0))
}

//func TestContentCaptureVideoWithDuration(t *testing.T) {
//	type fields struct {
//		AppId  string
//		AppKey string
//	}
//	type args struct {
//		webUrls  string
//		timeout  time.Duration
//		duration int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    string
//		wantErr bool
//	}{
//		{name: "1", fields: fields{AppId: "", AppKey: ""}, args: args{webUrls: "", timeout: 10, duration: 500}, want: "", wantErr: true},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			zxl, _ := zxl_go_sdk.NewZxlImpl("", "")
//			got, err := zxl.ContentCaptureVideoWithDuration(tt.args.webUrls, tt.args.timeout, tt.args.duration)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ContentCaptureVideoWithDuration() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("ContentCaptureVideoWithDuration() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
