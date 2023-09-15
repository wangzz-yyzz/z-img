package utils

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"z-img/config"
)

// GetMd5ByImgName read img and return the md5
func GetMd5ByImgName(name string) string {
	// read the file
	f, err := os.ReadFile(fmt.Sprintf("%s%s%s", config.ImgPath, "/", name))
	if err != nil {
		log.Println(err)
	}

	// cal the md5
	sig := md5.Sum(f)
	return fmt.Sprintf("%x", sig)
}
