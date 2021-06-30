package main

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func main() {
	tmpUid, _ := uuid.NewUUID()
	fmt.Println(tmpUid.String())
	tmpUid, _ = uuid.NewUUID()
	fmt.Println(tmpUid.String())
	tmpUid, _ = uuid.NewUUID()
	fmt.Println(tmpUid.String())
	tmpUid, _ = uuid.NewUUID()
	fmt.Println(tmpUid.String())
	tmpUid, _ = uuid.NewUUID()
	fmt.Println(tmpUid.String())
	tmpUid, _ = uuid.NewUUID()
	fmt.Println(tmpUid.String())

	idStr := strings.ReplaceAll(tmpUid.String(), "-", "")
	fmt.Println(idStr)
}
