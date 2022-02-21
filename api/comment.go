package api

import (
	"github.com/gin-gonic/gin"
	"path"
	"redrock_homework/service"
	"redrock_homework/tool"
	"strconv"
	"time"
)

//获取某一帖子或者是评论的评论
func Getcomment()gin.HandlerFunc{
	return func(c *gin.Context) {
		model,_:=strconv.Atoi(c.Query("model"))//表示要返回评论的类型，1为一级评论，2为二级评论
		target_id,_:=strconv.Atoi(c.Query("target_id"))//一级评论时为post_id,二级评论时为目标评论的id
		page,_:=strconv.Atoi(c.Query("page"))
		size,_:=strconv.Atoi(c.Query("size"))
		if model==1{
			array1,_:=service.Getcomment(model,target_id,page,size)
			data:=gin.H{
				"responses":array1,
			}
			c.JSON(200,gin.H{
				"data":data,
				"status":10000,
				"info":"success",
			})
		}else{
			_,array2:=service.Getcomment(model,target_id,page,size)
			data:=gin.H{
				"responses":array2,
			}
			c.JSON(200,gin.H{
				"data":data,
				"status":10000,
				"info":"success",
			})
		}

	}
}
//发布一条评论
func Publishcomment()gin.HandlerFunc{
	return func(c *gin.Context) {
		model,_:=strconv.Atoi(c.PostForm("model"))
		target_id,_:=strconv.Atoi(c.PostForm("target_id"))
		content:=c.PostForm("content")
		file,_:=c.FormFile("photo")
		dst:=path.Join("./files",file.Filename)
		_=c.SaveUploadedFile(file,dst)
		photo:="127.0.0.1:8080/"+dst
		claims,_:=c.Get("claims")
		claimsstr:=claims.(*tool.Myclaims)
		nickname:=claimsstr.Nickname
		user_id:=claimsstr.User_id
		avatar:=claimsstr.Avatar
		publish_time:=time.Now().Format("200601021504")
		praise_count:=0
		post_id:=service.Attainpost_id(target_id)//当model为2时，通过target_id获得post_id
		n:=service.Publishcomment(model,post_id,target_id,content,photo,nickname,user_id,avatar,publish_time,praise_count)
		comment_id:=service.Getcomment_id(content,model)
		data:=gin.H{
			"comment_id":comment_id,
		}
		if n==0{
			c.JSON(200,gin.H{
				"status":-1,
				"info":"发表评论失败",
			})
		}else{
			//发布评论后相关帖子评论数将发生改变
			if model==1{
				service.Add(target_id)
			}else {
				service.Add(post_id)
			}
			c.JSON(200,gin.H{
				"status":10000,
				"info":"success",
				"data":data,
			})
		}
	}
}
//更新一条评论
func Updatecomment()gin.HandlerFunc{
	return func(c *gin.Context) {
		comment_id, _ := strconv.Atoi(c.Param("comment_id"))
		content := c.PostForm("content")
		//通过comment_id,用户的相关信息判断要修改的评论是一级评论还是二级评论
		claims, _ := c.Get("claims")
		claimsstr := claims.(*tool.Myclaims)
		user_id := claimsstr.User_id
		nickname := claimsstr.Nickname
		avatar := claimsstr.Avatar
		photo := c.PostForm("photo")
		//如果查出来为一级评论
		if service.Confirm1(comment_id, nickname, user_id, avatar) != "" {
			if service.Updatecomment1(photo, content, comment_id) == 0 {
				c.JSON(200, gin.H{
					"status": -1,
					"info":   "更新失败",
				})
			} else {
				c.JSON(200, gin.H{
					"status": 10000,
					"info":   "success",
				})
			}

		}
		//如果查出来为二级评论
		if service.Confirm2(comment_id, nickname, user_id, avatar) != "" {
			if service.Updatecomment2(content, comment_id) == 0 {
				c.JSON(200, gin.H{
					"status": -1,
					"info":   "更新失败",
				})
			} else {
				c.JSON(200, gin.H{
					"status": 10000,
					"info":   "success",
				})
			}
		}
		if service.Confirm1(comment_id, nickname, user_id, avatar)== ""&&service.Confirm2(comment_id, nickname, user_id, avatar)==""{
			c.JSON(200,gin.H{
				"status":-1,
				"info":"要更新的评论不存在",
			})
		}
	}
}
//删除评论
func Deletecomment()gin.HandlerFunc{
	return func(c *gin.Context) {
		comment_id,_:=strconv.Atoi(c.Param("comment_id"))
		claims,_:=c.Get("claims")
		claimsstr:=claims.(*tool.Myclaims)
		user_id:=claimsstr.User_id
		nickname:=claimsstr.Nickname
		avatar:=claimsstr.Avatar
		//如果查出来为一级评论
		if service.Confirm1(comment_id,nickname,user_id,avatar)!=""{
			service.Deletecomment1(comment_id)
			c.JSON(200,gin.H{
				"status":10000,
				"info":"success",
			})
		}
		//如果查出来是二级评论
		if service.Confirm2(comment_id,nickname,user_id,avatar)!=""{
			service.Deletecomment2(comment_id)
			c.JSON(200,gin.H{
				"status":10000,
				"info":"success",
			})
		}
		if service.Confirm1(comment_id, nickname, user_id, avatar)== ""&&service.Confirm2(comment_id, nickname, user_id, avatar)==""{
			c.JSON(200,gin.H{
				"status":-1,
				"info":"没有要删除的评论",
			})
		}
	}
}