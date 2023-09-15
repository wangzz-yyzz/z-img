package utils

import (
	"fmt"
	"testing"
)

func TestGetMd5ByImgName(t *testing.T) {
	sig := GetMd5ByImgName("street.jpg")
	fmt.Printf("sig: %s,%T", sig, sig)
}
