package constants

/**
确权权利类型
*/
type RightType string

const (
	RIGHT_TYPE_PUBLIC       RightType = "发表权"     // 1
	RIGHT_TYPE_SIGN         RightType = "署名权"     // 2
	RIGHT_TYPE_MODIFY       RightType = "修改权"     // 3
	RIGHT_TYPE_FULL_PROTECT RightType = "保护权"     // 4
	RIGHT_TYPE_COPY         RightType = "复制权"     // 5
	RIGHT_TYPE_DISTRIBUTION RightType = "发行权"     // 6
	RIGHT_TYPE_RENT         RightType = "出租权"     // 7
	RIGHT_TYPE_DISPLAY      RightType = "展览权"     // 8
	RIGHT_TYPE_SHOW         RightType = "表演权"     // 9
	RIGHT_TYPE_PLAY         RightType = "放映权"     // 10
	RIGHT_TYPE_BROADCAST    RightType = "广播权"     // 11
	RIGHT_TYPE_NET_TRANS    RightType = "信息网络创博权" // 12
	RIGHT_TYPE_FILMING      RightType = "摄制权"     // 13
	RIGHT_TYPE_ADAPT        RightType = "改编权"     // 14
	RIGHT_TYPE_TRANSLATE    RightType = "翻译权"     // 15
	RIGHT_TYPE_COLLECTION   RightType = "汇编权"     // 16
	RIGHT_TYPE_OTHER        RightType = "其他权利"    // 17
	RIGHT_TYPE_ALL        RightType = "所有"    // 18
)
