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

func HexToStr(c *gin.Context) {
	h := string(c.PostForm("hex"))
	s := util.HexStr(h)
	c.JSON(200, gin.H{
		"message": s,
	})
}

func StrToHex(c *gin.Context) {
	s := string(c.PostForm("str"))
	h := util.StrHex([]byte(s))
	c.JSON(200, gin.H{
		"message": h,
	})
}
