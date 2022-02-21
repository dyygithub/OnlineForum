package api

import (
	"github.com/gin-gonic/gin"
	"path"
	"redrock_homework/dao"
	"redrock_homework/model"
	"redrock_homework/service"
	"redrock_homework/tool"
	"strconv"
	"time"
)

//获取帖子列表的接口函数
func PostList()gin.HandlerFunc{
	return func(c *gin.Context) {
		page,_:=strconv.Atoi(c.Query("page"))//表示要返回第几页的帖子
		size:=c.Query("size")//表示每一页帖子有多少个帖子
		sizeint,_:=strconv.Atoi(size)
		var array []model.Post
		arraypost,err:=service.List(page,array,sizeint)
		data:=gin.H{
			"post":arraypost,
		}
		if err!=nil{
			c.JSON(200,gin.H{
				"status":"-1",
				"info":err,
			})
		}else{
			c.JSON(200,gin.H{
				"status":"10000",
				"info":"success",
				"data":data,
			})
		}
	}
}
//获取某一个帖子的内容
func PostSingle() gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id:=c.Param("post_id")
		postid,_:=strconv.Atoi(post_id)
		post,err:=service.Single(postid)
		data:=gin.H{
			"data":post,
		}
		if err!=nil{
			c.JSON(200,gin.H{
				"status":-1,
				"info":err,
			})
		}else{
			c.JSON(200,gin.H{
				"status":10000,
				"info":"success",
				"data":data,
			})
		}
	}
}
//创建一个帖子
func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		content:=c.PostForm("content")
		topic_id:=c.PostForm("topic_id")
		file,_:=c.FormFile("pictures")
		dst:=path.Join("./files",file.Filename)
		c.SaveUploadedFile(file,dst)
		pictures:="127.0.0.1:8080/"+dst
		value,_:=c.Get("claims")
		claims:=value.(*tool.Myclaims)
		user_id:=claims.User_id
		avatar:=claims.Avatar
		nickname:=claims.Nickname
		pubulish_time:=time.Now().Format("200601021504")
		n:=service.CreatePost(pubulish_time,content,pictures,topic_id,user_id,avatar,nickname)
		post_id:=service.Getpost_id(content)
		data:=gin.H{
			"data":post_id,
		}
		if n==0{
			c.JSON(200,gin.H{
				"status":-1,
				"info":"创建失败",
			})
		}else {
			c.JSON(200,gin.H{
				"status":10000,
				"info":"success",
				"data":data,
			})
		}
	}
}
//更新一个帖子
func UpdatePost()gin.HandlerFunc{
	return func(c *gin.Context) {
		post_id:=c.Param("post_id")
		content:=c.PostForm("content")
		topic_id:=c.PostForm("topic_id")
		pictures:=c.PostForm("pictures")
		sql1:=`update post set content=? where post_id=?`
		sql2:=`update post set topic_id=? where post_id=?`
		sql3:=`update post set pictures=? where post_id=?`
		n1:=dao.Updatedb(sql1,content,post_id)
		n2:=dao.Updatedb(sql2,topic_id,post_id)
		n3:=dao.Updatedb(sql3,pictures,post_id)
		if!(n1!=0&&n2!=0&&n3!=0){
			c.JSON(200,gin.H{
				"status":-1,
				"info":"更新失败",
			})
		}else {
			c.JSON(200,gin.H{
				"status":10000,
				"info":"success",
			})
		}
	}
}
//删除一个帖子
func DeletePost()gin.HandlerFunc{
	return func(c *gin.Context) {
		post_id,_:=strconv.Atoi(c.Param("post_id"))
		if service.Deletepost(post_id)==0{
			c.JSON(200,gin.H{
				"status":-1,
				"info":"删除失败",
			})
		}else{
			//删除帖子下面对应的评论
			service.Deletecomment(post_id)
			//删除帖子对应得收藏数据
			service.Deletecollection(post_id)
			c.JSON(200,gin.H{
				"status":10000,
				"info":"success",
			})
		}
	}
}
//搜索帖子
func Searchpost()gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Query("key")
		page,_ :=strconv.Atoi(c.Query("page"))
		size := c.Query("size")
		sizeint, _ := strconv.Atoi(size)
		var array []model.Post
		arraypost, err := service.Searchpost(page,array,sizeint,key)
		data := gin.H{
			"post": arraypost,
		}
		if err != nil {
			c.JSON(200, gin.H{
				"status": "-1",
				"info":   err,
			})
		} else {
			c.JSON(200, gin.H{
				"status": "10000",
				"info":   "success",
				"data":   data,
			})
		}
	}
}
