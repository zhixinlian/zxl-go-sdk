package constants

/**
接入平台业务类型
*/

type BusinessType int

const (
	BUSINESS_FINANCE   BusinessType = 1 // 金融类
	BUSINESS_COPYRIGHT BusinessType = 2 // 版权类
	BUSINESS_OTHER     BusinessType = 3 // 其他类
	BUSINESS_DEFAULT   BusinessType = 4 // 未填写
)
