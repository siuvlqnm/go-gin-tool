package system

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/go-gin-tool/internal/util"
)

func Md5Encode(c *gin.Context) {
	b := []byte(c.PostForm("str"))
	m := util.MD5V(b)
	c.JSON(200, gin.H{
		"message": m,
	})
}
