package api

import (
	"github.com/gin-gonic/gin"
	"redrock_homework/service"
	"redrock_homework/tool"
	"strconv"
)


//点赞操作
func Praise()gin.HandlerFunc{
	return func(c *gin.Context) {
		model,_:=strconv.Atoi(c.PostForm("model"))
		target_id,_:=strconv.Atoi(c.PostForm("target_id"))
		//当model为1时，点赞帖子
		//当model为2时，点赞一级评论
		//当model为3时，点赞二级评论
		service.Praise(model,target_id)
		c.JSON(200,gin.H{
			"status":10000,
			"info":"success",
		})
	}
}
//获取关注列表
func Getfocuslist()gin.HandlerFunc{
	return func(c *gin.Context) {
		claims,_:=c.Get("claims")
		claimsstr:=claims.(*tool.Myclaims)
		username:=claimsstr.Username
		u:=service.Getfocuslist(username)
		data:=gin.H{
			"focus":u,
		}
		c.JSON(200,gin.H{
			"status":10000,
			"info":"success",
			"data":data,
		})
	}
}
//关注操作
func Focus()gin.HandlerFunc{
	return func(c *gin.Context) {
		user_id:=c.PostForm("user_id")
		claims,_:=c.Get("claims")
		claimsstr:=claims.(*tool.Myclaims)
		username:=claimsstr.Username
		if username==service.Queryusername(user_id){
			c.JSON(200,gin.H{
				"status":-1,
				"info":"不能关注自己",
			})
		}else{
			service.Collectpost(user_id,username)
			c.JSON(200,gin.H{
				"status":10000,
				"info":"success",
			})
		}
	}
}
//获取用户收藏列表
func Getcollection()gin.HandlerFunc{
	return func(c *gin.Context) {
		claims,_:=c.Get("claims")
		claimsstr:=claims.(*tool.Myclaims)
		username:=claimsstr.Username
		u:=service.Getcollection(username)
		data:=gin.H{
			"focus":u,
		}
		c.JSON(200,gin.H{
			"status":10000,
			"info":"success",
			"collections":data,
		})
	}
}
//收藏帖子
func Collectpost()gin.HandlerFunc{
	return func(c *gin.Context) {
		post_id,_:=strconv.Atoi(c.PostForm("post_id"))
		claims,_:=c.Get("claims")
		claimsstr:=claims.(*tool.Myclaims)
		username:=claimsstr.Username
		service.Collect(post_id,username)
		c.JSON(200,gin.H{
			"status":10000,
			"info":"success",
		})
	}
}