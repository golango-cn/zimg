package controller

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/gin-gonic/gin"
	"zimg/core"
	"zimg/core/common"
)

type ZImgController struct {
}

func (*ZImgController) Index(ctx *gin.Context) {

	formFile, err := ctx.FormFile("userfile")
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	f, err := formFile.Open()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	defer f.Close()

	zimg := &core.ZImg{File: f}
	md5, err := zimg.Upload()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 200,
		"md5":  md5,
	})

}

func (*ZImgController) Get(ctx *gin.Context) {

	url := fmt.Sprintf("%s%s", common.ServerSetting.ZImgServer, ctx.Request.RequestURI)

	requ := httplib.Get(url)
	b, _ := requ.Bytes()

	ctx.Writer.Write(b)

}
