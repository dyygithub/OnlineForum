package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock_homework/service"
	"redrock_homework/tool"
	"strings"
)

//jwt中间件，对token进行检查
func JwtAuth()gin.HandlerFunc{
	return func(c *gin.Context) {
		authheader :=c.Request.Header.Get("Authorization")
		if authheader ==""{
			c.JSON(http.StatusOK,gin.H{
				"status":-1,
				"info":"无权访问,请求并未携带token",
			})
			c.Abort()//结束后续操作
			return
		}
		//log.Print("token:",authheader)
		//按空格拆分
		parts :=strings.SplitN(authheader," ",2)
		if!(len(parts)==2&&parts[0]=="Bearer"){
			c.JSON(200,gin.H{
				"status":-1,
				"info":"请求头中的auth格式有误",
			})
			c.Abort()
			return
		}
		//解析token包含的信息
		claims,err:=tool.Parsetoken(parts[1])
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"status":-1,
				"info":"无效的token",
			})
			c.Abort()
			return
		}
		//鉴别是token还是refreshtoken
		if claims.Power!="admin"{
			c.JSON(200,gin.H{
				"status":-1,
				"info":"没有刷新的权限",
			})
			c.Abort()
			return
		}
		err=CheckUserInfo(claims)
		if err!=nil{
			c.JSON(200,gin.H{
				"status":-1,
				"info":"用户名或密码错误",
			})
			c.Abort()
			return
		}
		c.Set("claims",claims)
		c.Next()
	}
}

//检查用户名信息是否正确
func CheckUserInfo(claims *tool.Myclaims)error{
	username:=claims.Username
	password:=claims.Password
	if service.IsTrue(username,password){
		return nil
	}
	return errors.New("用户名或密码错误")
}

