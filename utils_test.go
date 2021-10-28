package zxl_go_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIp(t *testing.T) {
	a := assert.New(t)

	domain := "www.test.com"
	result := checkIp(domain)
	a.Equal(result, false)


	ip := "128.167.14.65"
	result = checkIp(ip)
	a.Equal(result, true)

	wrongIp := "128.78.6"
	result = checkIp(wrongIp)
	a.Equal(result, false)
}

func TestIsInnerIpFromUrl(t *testing.T) {
	a := assert.New(t)

	domain := "www.test.com"
	result := isInnerIpFromUrl(domain)
	a.Equal(result, false)

	url := "https://www.test.com?query=1"
	result = isInnerIpFromUrl(url)
	a.Equal(result, false)

	innerIp := "http://10.12.56.78?query=1"
	result = isInnerIpFromUrl(innerIp)
	a.Equal(result, true)
}