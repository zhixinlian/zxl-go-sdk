package constants

/**
作品类型
**/

type DciType string

const (
	DCI_TYPE_TEXT             DciType = "文字作品"           // 1
	DCI_TYPE_MUSIC            DciType = "音乐作品"           // 2
	DCI_TYPE_ART_ARCH         DciType = "美术或建筑作品"        // 3
	DCI_TYPE_FILMING          DciType = "摄影作品"           // 4
	DCI_TYPE_AUDIOVISUAL      DciType = "视听作品"     // 5
	DCI_TYPE_PIC_MODEL        DciType = "设计图或示意图等图形模型作品" // 6
	DCI_TYPE_SOFTWARE         DciType = "计算机软件"          // 7
	DCI_TYPE_OTHER            DciType = "符合作品特征的其他智力成果"        //8
)
