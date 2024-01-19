/*
 Package test test-ceping_test
    - @File:  ceping_test
    - @Description: // 清洁性测评专用
    - @Author: suxiongye
    - @Date: 2024/1/18 15:53
    - @Copyright: Tencent
*/

package test

import (
	"flag"
	"fmt"
	zxlgosdk "github.com/zhixinlian/zxl-go-sdk/v2"
	"strconv"
	"testing"
)

/**
查询取证结果
*/
func TestGetObtainResult(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	var argLen = 4
	argList := flag.Args() // flag.Args() 返回 -args 后面的所有参数，以切片表示，每个元素代表一个参数
	if len(argList) < argLen {
		t.Fatalf("缺少参数，参数数量要求至少为：%d 个", argLen)
	}
	// 测试环境
	var serverAddr = "https://testsdk.zxinchain.com"
	var env = argList[0]
	if env == "prod" {
		serverAddr = "https://sdk.zxinchain.com"
	}
	// 代理商
	var appid = argList[1]
	var appKey = argList[2]
	var orderNo = argList[3]

	var config = &zxlgosdk.ZxlConfig{
		AppId:      appid,
		AppKey:     appKey,
		ServerAddr: serverAddr,
	}
	zxlSdk, err := zxlgosdk.CreateZxlClientWithConfig(*config)

	evIdData, err := zxlSdk.GetEvidenceStatus(orderNo, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("evIdData is %+v\n", evIdData)
}

/**
查询取证结果
*/
func TestSendMobileObtain(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	var argLen = 8
	argList := flag.Args() // flag.Args() 返回 -args 后面的所有参数，以切片表示，每个元素代表一个参数
	if len(argList) < argLen {
		t.Fatalf("缺少参数，参数数量要求至少为：%d 个", argLen)
	}
	// 测试环境
	var serverAddr = "https://testsdk.zxinchain.com"
	var env = argList[0]
	if env == "prod" {
		serverAddr = "https://sdk.zxinchain.com"
	}
	// 代理商
	var appid = argList[1]
	var appKey = argList[2]
	var shareUrl = argList[3]
	var appName = argList[4]
	var title = argList[5]
	var remark = argList[6]
	var durationStr = argList[7]

	var duration, err = strconv.Atoi(durationStr)
	if err != nil {
		t.Fatalf("录制时间长度要求舒服整数。%s", durationStr)
	}

	var config = &zxlgosdk.ZxlConfig{
		AppId:      appid,
		AppKey:     appKey,
		ServerAddr: serverAddr,
	}
	zxlSdk, err := zxlgosdk.CreateZxlClientWithConfig(*config)

	evIdData, err := zxlSdk.EvidenceObtainMobile(shareUrl, appName, title, remark, duration, 0)
	if err != nil {
		t.Fatalf("发起手机取证错误。%s", err.Error())
	}
	fmt.Println(fmt.Sprintf("取证任务Id： %s", evIdData))
}
