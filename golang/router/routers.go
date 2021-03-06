package routers

import (
	"github.com/gin-gonic/gin"
	"zimg/controller"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	zimg := controller.ZImgController{}
	r.POST("/", zimg.Index)
	r.GET("/:md5", zimg.Get)

	return r

}
