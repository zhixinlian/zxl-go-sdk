package test

import (
	"fmt"
	zxl_go_sdk "github.com/zhixinlian/zxl-go-sdk/v2"
	"github.com/zhixinlian/zxl-go-sdk/v2/constants"
	"testing"
	"time"
)

var appId = ""
var appKey = ""
var pk = ""
var sk = ""

func TestEpRegister(t *testing.T) {
	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	user := zxl_go_sdk.AgentUser{
		EpName:          "",
		Pwd:             "123456aaa",
		RepresentEmail:  "xxx@admin.com",
		Representative:  "",
		CreditCode:      "",
		Contact:         "",
		Idcard:          "",
		Mobile:          "",
		PlatformName:    "小说平台",
		PlatformUrl:     "https://www.baldu.com",
		BusinessType:    constants.BUSINESS_COPYRIGHT,
		UserType:        constants.USER_LEGAL_PERSON,
		CardFrontFile:   "./idcard_front.jpeg",
		CardBackendFile: "./idcard_front.jpeg",
		LicenseFile:     "./license.jpeg",
	}
	result, err := zxlSDK.RegisterUser(user, 5*time.Second)

	if err != nil {
		fmt.Println("注册用户错误: " + err.Error())
		return
	}

	fmt.Printf("注册结果：%+v", result)
}

func TestQueryEpInfo(t *testing.T) {
	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	result, err := zxlSDK.SelectEpInfo("xx@admin.com", 5*time.Second)
	if err != nil {
		fmt.Println("查询企业用户错误：" + err.Error())
		return
	}

	fmt.Printf("查询结果： %+v", result)
}

func TestRegisterPerson(t *testing.T) {
	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	user := zxl_go_sdk.AgentUser{
		PersonName:      "",
		Pwd:             "123456aaa",
		RepresentEmail:  "xxx@admin.com",
		Idcard:          "",
		Mobile:          "",
		UserType:        constants.USER_NATURAL_PERSON,
		CardFrontFile:   "./idcard_front.jpeg",
		CardBackendFile: "./idcard_front.jpeg",
	}
	result, err := zxlSDK.RegisterUser(user, 5*time.Second)

	if err != nil {
		fmt.Println("注册用户错误: " + err.Error())
		return
	}

	fmt.Printf("注册结果：%+v", result)
}

func TestQueryPerson(t *testing.T) {
	zxlSDK, err := zxl_go_sdk.NewZxlImpl(appId, appKey)
	if err != nil {
		fmt.Println("初始化 SDK 错误")
	}

	result, err := zxlSDK.SelectEpInfo("xx@admin.com", 5*time.Second)

	if err != nil {
		fmt.Printf("查询用户错误： %+v", err)
		return
	}
	fmt.Printf("查询结果：%+v", result)
}
