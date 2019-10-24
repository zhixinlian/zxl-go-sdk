package main

import (
	"fmt"
	uuid "github.com/zhixinlian/zxl-go-sdk/satori/go.uuid"
	"strings"
)

func main() {
	tmpUid, err := uuid.NewV1()
	fmt.Println(err)
	fmt.Println(tmpUid.String())
	tmpUid, err = uuid.NewV1()
	fmt.Println(err)
	fmt.Println(tmpUid.String())
	tmpUid, err = uuid.NewV1()
	fmt.Println(err)
	fmt.Println(tmpUid.String())
	tmpUid, err = uuid.NewV1()
	fmt.Println(err)
	fmt.Println(tmpUid.String())
	tmpUid, err = uuid.NewV1()
	fmt.Println(err)
	fmt.Println(tmpUid.String())
	tmpUid, err = uuid.NewV1()
	fmt.Println(err)
	fmt.Println(tmpUid.String())

	idStr := strings.ReplaceAll(tmpUid.String(), "-", "")
	fmt.Println(idStr)
}
