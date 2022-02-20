package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redrock_homework/dao"
	"redrock_homework/model"
	"redrock_homework/service"
	"redrock_homework/tool"
)
//注册相关的接口函数
func Register()gin.HandlerFunc{
	return func(c *gin.Context){
		username:=c.PostForm("username")
		password:=c.PostForm("password")
        if !service.IspeatUsername(username){
			c.JSON(400,gin.H{
				"status":-1,
				"info":"注册账号重复",
			})
			return
		}
		if service.Insert(username,password)!=0{
			data:=gin.H{"refresh_token": "null","token": "null"}
			c.JSON(200,gin.H{
				"status": 10000,
				"info": "success",
				"data": data,
			})
		}else{
			c.JSON(200,gin.H{
				"status":-1,
				"info":"注册失败",
			})
		}


	}
}
//登录相关的接口函数
func Login()gin.HandlerFunc{
	return func(c *gin.Context){
		username:=c.Query("username")
		password:=c.Query("password")
		if !service.IsTrue(username,password){
			c.JSON(200,gin.H{
				"status":-1,
				"info":"密码错误",
			})
			return
		}
		fmt.Println("登录成功")
		u:=model.UserAccount{
			Username: username,
			Password: password,
		}
		////将用户信息放入token中
		sqlstr:=`select id,avatar,nickname,introduction,phone ,qq,gender,email,birthday from userinfo where username=?`
		row:=dao.Queryone(sqlstr,username)
		userinfo:=model.Userinfo{}
		err:=row.Scan(&userinfo.User_id,&userinfo.Avatar,&userinfo.Nickname,&userinfo.Introduction,&userinfo.Phone,&userinfo.Qq,&userinfo.Gender,&userinfo.Email,&userinfo.Birthday)
		if err!=nil{
			fmt.Println(err)
		}
		token,_:=tool.Gentoken(u,userinfo)
		refresh_token,_:=tool.Genretoken(u,userinfo)
		data:=gin.H{"refresh_token": refresh_token,"token": token}
		c.JSON(200,gin.H{
			"status":10000,
			"info":"success",
			"data":data,
		})
	}
}
//刷新token的接口函数,本质上是利用refre_token创建一个新的token和新的refresh_token
func Refresh()gin.HandlerFunc{
	return func(c *gin.Context) {
		refresh_tokenstr:=c.Query("refresh_token")
		claims,err:=tool.Parsetoken(refresh_tokenstr)
		if err!=nil{
			fmt.Println("解析refresh_token失败")
		}
		if claims.Power!="refresh"{
			c.JSON(200,gin.H{
				"status":-1,
				"info":"没有刷新的权限",
			})
			return
		}
		//将用户信息放入token中
		useraccount:=claims.UserAccount
		username:=useraccount.Username
		sqlstr:=`select id,avatar,nickname,introduction,phone ,qq,gender,email,birthday from userinfo where username=?`
		row:=dao.Queryone(sqlstr,username)
		userinfo:=model.Userinfo{}
		err=row.Scan(&userinfo.User_id,&userinfo.Avatar,&userinfo.Nickname,&userinfo.Introduction,&userinfo.Phone,&userinfo.Qq,&userinfo.Gender,&userinfo.Email,&userinfo.Birthday)
		if err!=nil{
			fmt.Println(err)
		}
		token,err:=tool.Gentoken(useraccount,userinfo)
		if err!=nil{
			fmt.Println("刷新token失败")
		}
		refersh_token,err:=tool.Genretoken(useraccount,userinfo)
		if err!=nil{
			fmt.Println("刷新refresh_token失败")
		}
		data:=gin.H{
			"token":token,
			"refresh_token":refersh_token,
		}
		c.JSON(200,gin.H{
			"status":10000,
			"info":"success",
			"data":data,
		})
	}
}
//修改密码的接口函数
func Correct()gin.HandlerFunc{
	return func(c *gin.Context) {
		claims,_:=c.Get("claims")//从中间件中或取解析token返回的数据
		Password:=claims.(*tool.Myclaims).Password
		old_password:= c.Query("old_password")
		new_password:= c.Query("new_password")
		if old_password!=Password{
			c.JSON(200,gin.H{
				"status":-1,
				"info":"输入的旧密码有错",
			})
		}else{
			sqlstr:=`update useraccount set password=? where password=?`
			n:=dao.Updatedb(sqlstr,new_password,old_password)//返回的是更改数据的数目
			if n==0{
				c.JSON(200,gin.H{
					"status":-1,
					"info":"修改密码失败",
				})
			}else{
				c.JSON(200,gin.H{
					"status":10000,
					"info":"success",
				})
			}
		}

	}
}
//查看用户信息的接口函数
func QueryUserInfo()gin.HandlerFunc{
	return func(c *gin.Context) {
		id:=c.Param("id")
		fmt.Println(service.QueryInfo(id))
		data:=gin.H{"data":service.QueryInfo(id)}
		c.JSON(200,gin.H{
			"status":10000,
			"info":"success",
			"data":data,
		})

	}
}
//更改用户的信息的接口函数
func UpdateUserInfo()gin.HandlerFunc{
	return func(c *gin.Context) {
		nickname:=c.PostForm("nickname")//传入nickname，进行更改
		introduction:=c.PostForm("introduction")//传入introduction，进行更改
		newnickname:=c.PostForm("newnickname")
		newintroduction:=c.PostForm("newintroduction")
		sqlstr1:=`update userinfo set nickname=? where nickname=?`
		sqlstr2:=`update userinfo set introduction=? where introduction=?`
		n1:=dao.Updatedb(sqlstr1,newnickname,nickname)
		n2:=service.UpdateInfo(sqlstr2,newintroduction,introduction)
		if n1==0&&n2==0 {
			c.JSON(200,gin.H{
				"status":-1,
				"info":"更改失败",
			})

		}else{
			c.JSON(200,gin.H{
				"status":10000,
				"info":"success",
			})
		}
	}
}


