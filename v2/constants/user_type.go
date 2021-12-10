package constants

/**
平台用户类型
*/
type UserType int

const (
	USER_LEGAL_PERSON   UserType = 1 // 法人主体
	USER_NATURAL_PERSON UserType = 2 // 自然人主体
)
