package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestUUID(t *testing.T) {
	tmpUid := uuid.New()
	fmt.Println(tmpUid.String())
	tmpUid = uuid.New()
	fmt.Println(tmpUid.String())
	tmpUid = uuid.New()
	fmt.Println(tmpUid.String())
	tmpUid = uuid.New()
	fmt.Println(tmpUid.String())
	tmpUid = uuid.New()
	fmt.Println(tmpUid.String())
	tmpUid = uuid.New()
	fmt.Println(tmpUid.String())

	idStr := strings.ReplaceAll(tmpUid.String(), "-", "")
	fmt.Println(idStr)
}
