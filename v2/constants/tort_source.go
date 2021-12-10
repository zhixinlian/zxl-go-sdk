package constants


/**
 * 数据源头
 */
type TortSource int

const (
	PROPERTY           TortSource = 1  //房产类监测，type 为图片时可用
	ANY                TortSource = 2  //原创类监测，type 为图片时可用
	SHORT_VIDEO        TortSource = 21 //短视频监测，type 为视频时可用
	LONG_VIDEO         TortSource = 22 //长对短监测:监测范围和短视频一样，type 为长视频时可用
	LONG_TO_LONG_VIDEO TortSource = 23 //长对长监测:全网视频小网站，type为视频时可用
	NEWS_TEXT          TortSource = 41 //新闻咨询类，监测范围全网
)
