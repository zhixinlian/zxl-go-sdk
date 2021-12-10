package constants

/**
作品创作属性
**/
type DciCreateProperty string

const (
	DCI_CREATE_PROPERTY_ORIGINAL   DciCreateProperty = "原创" // 1
	DCI_CREATE_PROPERTY_ADAPT      DciCreateProperty = "改编" // 2
	DCI_CREATE_PROPERTY_TRANSLATE  DciCreateProperty = "翻译" // 3
	DCI_CREATE_PROPERTY_COLLECTION DciCreateProperty = "汇编" // 4
	DCI_CREATE_PROPERTY_COMMENT    DciCreateProperty = "注释" //5
	DCI_CREATE_PROPERTY_TIDY       DciCreateProperty = "整理" // 6
	DCI_CREATE_PROPERTY_OTHER      DciCreateProperty = "其他" //7
)
