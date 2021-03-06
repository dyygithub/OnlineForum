package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock_homework/api"
	"redrock_homework/dao"
)
func main(){
	r:=gin.Default()
	//连接数据库
	dao.Initdb()
	//加载静态文件。
	r.Static("/files","./files")
	{
		//注册
		r.POST("/user/register",api.Register())
		//登录并且获取token
		r.GET("/user/token",api.Login())
		//通过第三方授权登录
		r.GET("oauth/redirect",api.Oauth())
		//重定向到要借助的第三方网站，这里指的是github
		r.GET("/redirect/github", func(c *gin.Context) {
			 c.Redirect(http.StatusMovedPermanently,"https://github.com" +
				 "/login/oauth/authorize?client_id=552c51bfa7dad98e1d8e" +
				 "&redirect_url=http://127.0.0.1:8080/oauth/redirect")
		})
		//刷新token
		r.GET("/user/token/refresh",api.Refresh())
		//修改密码
		r.PUT("/user/password",api.JwtAuth(),api.Correct())
		//获取用户信息
		r.GET("/user/info/:id",api.QueryUserInfo())
		//更改用户信息
		r.PUT("/user/info",api.UpdateUserInfo())
	}
	{
		//显示帖子列表
		r.GET("/post/list",api.PostList())
		//显示某个帖子
		r.GET("/post/single/:post_id",api.PostSingle())
		//发布一个帖子
		r.POST("/post/single",api.JwtAuth(),api.CreatePost())
		//更新一个帖子
		r.PUT("/post/single/:post_id",api.JwtAuth(),api.UpdatePost())
		//删除一个帖子
		r.DELETE("/post/single/:post_id",api.JwtAuth(),api.DeletePost())
		//搜索帖子
		r.GET("/post/search",api.Searchpost())
	}
	//获取所有主题
	r.GET("/topic/list",api.Gettopic())
	{
		//获取某个帖子/评论下的评论
		r.GET("/comment",api.Getcomment())
		//发布一条评论
		r.POST("/comment",api.JwtAuth(),api.Publishcomment())
		//更新一条评论
		r.PUT("/comment/:comment_id",api.JwtAuth(),api.UpdatePost())
		//删除一条评论
		r.DELETE("/comment/:comment_id",api.JwtAuth(),api.Deletecomment())
	}
	{
		//点赞操作
		r.PUT("/operate/praise",api.JwtAuth(),api.Praise())
		//获取关注列表
		r.GET("/operate/focus/list",api.JwtAuth(),api.Getfocuslist())
		//关注
		r.PUT("/operate/focus",api.JwtAuth(),api.Focus())
		//获取收藏列表
		r.GET("/operate/collect/list",api.JwtAuth(),api.Getcollection())
		//收藏
		r.PUT("operate/collect",api.JwtAuth(),api.Collectpost())
	}
	r.Run(":8080")
}
