package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/go-gin-tool/api/v1/system"
)

func Router() {
	r := gin.Default()

	r.POST("md5", system.Md5Encode)

	r.Run()
}
