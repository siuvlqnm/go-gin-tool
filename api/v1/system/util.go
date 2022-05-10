package system

import (
	"github.com/gin-gonic/gin"
	response "github.com/siuvlqnm/go-gin-tool/common"
	"github.com/siuvlqnm/go-gin-tool/internal/util"
)

func Md5Encode(c *gin.Context) {
	b := []byte(c.PostForm("str"))
	m := util.MD5V(b)
	response.OkWithDetailed("md5 encode success", m, c)
}

func HexToStr(c *gin.Context) {
	h := string(c.PostForm("hex"))
	s := util.HexStr(h)
	response.OkWithDetailed("hex to str success", s, c)
}

func StrToHex(c *gin.Context) {
	s := string(c.PostForm("str"))
	h := util.StrHex([]byte(s))
	response.OkWithDetailed("str to hex success", h, c)
}
