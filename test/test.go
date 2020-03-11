package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func main() {
	tmpUid := uuid.NewV1()
	fmt.Println(tmpUid.String())
	tmpUid = uuid.NewV1()
	fmt.Println(tmpUid.String())
	tmpUid = uuid.NewV1()
	fmt.Println(tmpUid.String())
	tmpUid = uuid.NewV1()
	fmt.Println(tmpUid.String())
	tmpUid = uuid.NewV1()
	fmt.Println(tmpUid.String())
	tmpUid = uuid.NewV1()
	fmt.Println(tmpUid.String())

	idStr := strings.ReplaceAll(tmpUid.String(), "-", "")
	fmt.Println(idStr)
}
