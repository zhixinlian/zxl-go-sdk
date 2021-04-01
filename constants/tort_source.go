package constants


/**
 * 数据源头
 */
type TortSource int

const (
	PROPERTY    TortSource = 1  //房产类监测，type 为图片时可用
	ANY         TortSource = 2  //原创类监测，type 为图片时可用
	SHORT_VIDEO     TortSource = 21 //短视频监测，type 为视频时可用
	LONG_VIDEO      TortSource = 22 //长视频监测，type 为长视频时可用
)
