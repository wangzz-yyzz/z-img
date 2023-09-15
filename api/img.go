package api

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path/filepath"
	"z-img/config"
	"z-img/database"
	"z-img/utils"
)

// GetImgById show the img by id(md5), if file not exist, return the 404 code
func GetImgById(c *gin.Context) {
	imgId := c.Query("id")

	writer := c.Writer

	// try to open file
	var file *os.File
	defer file.Close()
	for _, value := range database.ExtList {
		file, _ = os.Open(fmt.Sprintf("%s/%s%s", config.ImgPath, imgId, value))
		// find the file
		if file != nil {
			// read and show the pic
			buff, _ := io.ReadAll(file)
			_, _ = writer.Write(buff)
			c.JSON(200, gin.H{
				"msg": "pong",
			})
			return
		}
	}

	// can not find the file
	c.JSON(404, gin.H{
		"message": "wrong id",
	})
}

// PostImg handle the post pic.
// check if there is a same pic
// check if the pic is valid
// rename the pic and save it
func PostImg(c *gin.Context) {
	// get the posted pic
	f, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	log.Println("posting: ", f.Filename)

	// check the type
	typeOk := false
	typeIndex := 0
	for index, value := range database.ExtList {
		if filepath.Ext(f.Filename) == value {
			typeOk = true
			typeIndex = index
			break
		}
	}
	if !typeOk {
		c.JSON(250, gin.H{
			"msg": "not supported type yet",
		})
		return
	}

	// open the file and cal the md5
	file, _ := f.Open()
	buff, _ := io.ReadAll(file)
	sig := md5.Sum(buff)
	md5Code := fmt.Sprintf("%x", sig)

	// if there is a same pic, return the id
	for _, v := range database.Md5List {
		if v == md5Code {
			c.JSON(200, gin.H{
				"msg": "same pic",
				"id":  md5Code,
			})
			return
		}
	}

	// else, save the md5 code
	database.Md5List = append(database.Md5List, md5Code)
	log.Println("md5: ", md5Code)

	// save the file
	newName := fmt.Sprintf("%s%s", md5Code, database.ExtList[typeIndex])
	_ = c.SaveUploadedFile(f, config.ImgPath+"/"+newName)

	// close the file
	defer file.Close()
	log.Println("save done")

	c.JSON(200, gin.H{
		"msg": "pong",
		"id":  md5Code,
	})
}

// RemoveImgById remove the img by id(md5), if file not exist, return the 404 code
func RemoveImgById(c *gin.Context) {
	// get param id
	imgId := c.Query("id")

	// check if the file exist
	for _, value := range database.ExtList {
		err := os.Remove(fmt.Sprintf("%s/%s%s", config.ImgPath, imgId, value))
		if err == nil {
			// remove the md5 code
			database.Md5List = utils.DeleteVal(database.Md5List, imgId)
			c.JSON(200, gin.H{
				"msg": "pong",
			})
			return
		}
	}

	// can not find the file
	c.JSON(404, gin.H{
		"message": "wrong id",
	})
}
