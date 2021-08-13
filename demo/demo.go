package demo

import (
	"fmt"
	zxl "github.com/zhixinlian/zxl-go-sdk"
)

func main() {
	config := zxl.ZxlConfig{
		AppId: "",
		AppKey: "",
		IsProxy: true,
		ServerAddr: "https://sdk.zxinchain.com",
		ProxyHost: "127.0.0.1",
		ProxyPort: "8080",
	}

	zxlClient, err := zxl.CreateZxlClientWithConfig(config)

	if err != nil {
		panic(err)
	}

	hash, err := zxlClient.CalculateHash("zxinchain")

	fmt.Println(hash)
}
