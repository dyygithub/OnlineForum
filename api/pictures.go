package api

import (
	"github.com/gin-gonic/gin"
	"path"
)

func Showpictures()gin.HandlerFunc{
	return func(c *gin.Context) {

		//获取图片
		c.File("/files/f1.jpg")
	}
}

func Uppictures()gin.HandlerFunc{
	return func(c *gin.Context) {
		f,err:=c.FormFile("f1")
		if err!=nil{
			c.JSON(200,gin.H{
				"error":err.Error(),
			})
		}else{

			dst:=path.Join("./files","f5.","jpg")
			c.SaveUploadedFile(f,dst)
			c.JSON(200,gin.H{
				"status":"ok",
			})
		}
	}
}
