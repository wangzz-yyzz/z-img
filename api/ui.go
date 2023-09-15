package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"z-img/database"
)

// GetIndex render the index html
func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": database.Md5List,
	})
}

// GetMd5List return the md5 list
func GetMd5List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"list": database.Md5List,
	})
}
