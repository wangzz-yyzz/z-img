package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"z-img/config"
	"z-img/utils"
)

var Md5List []string
var TypeList = [7]string{"jpg", "png", "jpeg", "gif", "ico", "webp", "bmp"}
var ExtList = [7]string{".jpg", ".png", ".jpeg", ".gif", ".ico", ".webp", ".bmp"}

// get all the pic and the md5 to the pics, then rename all the pic with md5
func init() {
	// get all file in img
	log.Println("getting all the pic")
	var index = 0
	var old = 0
	err := filepath.Walk(config.ImgPath, func(path string, info os.FileInfo, err error) error {
		// keep the file end with jpg or png only
		typeOk := false
		typeIndex := 0
		for index, value := range ExtList {
			// if the type is valid, change typeOk to true
			if value == filepath.Ext(path) {
				typeOk = true
				typeIndex = index
			}
		}

		// if the type is invalid, skip it
		if !typeOk {
			return nil
		}

		// get the md5
		sig := utils.GetMd5ByImgName(info.Name())
		Md5List = append(Md5List, sig)

		// get the ext
		var ext string
		ext = TypeList[typeIndex]

		// get the new name
		var newName = fmt.Sprintf("%s.%s", sig, ext)
		err = os.Rename(path, fmt.Sprintf("%s/%s", config.ImgPath, newName))

		// find the existing pic
		if newName == info.Name() {
			old += 1
		}
		index += 1
		return nil
	})

	if err != nil {
		log.Println(err)
	}
	log.Printf("init database done, init %d pics total, old %d", index, old)

}
