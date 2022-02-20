package tool

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"redrock_homework/model"
	"time"
)
type Myclaims struct {
	model.UserAccount
	model.Userinfo
	Power string
	jwt.StandardClaims
}
var mysecret = []byte("woshidyy")
//生成token
func Gentoken(user model.UserAccount,userinfo model.Userinfo)(string,error) {
     claims:= Myclaims{
		 user,
		 userinfo,
		 "admin",
		 jwt.StandardClaims{
			 NotBefore: time.Now().Unix(),
			 ExpiresAt: time.Now().Unix() +60*5,
			//签发5分钟后过期
			 Issuer: "dyy",//签发人
		 },
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenstr,err:=token.SignedString(mysecret)
	if err!=nil{
		fmt.Println(err)
	}
	return tokenstr,nil
}
//生成refresh_token
func Genretoken(user model.UserAccount,userinfo model.Userinfo)(string,error) {
	claims := Myclaims{
		user,
		userinfo,
		"refresh",
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*12*7,
			//签发7天后过期
			Issuer: "dyy", //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstr, err := token.SignedString(mysecret)
	if err != nil {
		fmt.Println(err)
	}
	return tokenstr, nil
}
//解析token
func Parsetoken(tokenstr string)(*Myclaims,error){
	token,err:=jwt.ParseWithClaims(tokenstr,&Myclaims{},func(token *jwt.Token)(interface{},error){
		return mysecret,nil
	})
	if err!=nil{
		fmt.Println(" token parse err:", err)
		return nil, err
	}
	claims:=token.Claims.(*Myclaims)
	return claims,nil
}
////刷新token
//func Refreshtoken(tokenstr string)string{
//	jwt.TimeFunc=func()time.Time{
//		return time.Unix(0,0)
//	}
//	token,err := jwt.ParseWithClaims(tokenstr,&Myclaims{}, func(token *jwt.Token) (interface{}, error) {
//		return mysecret,nil
//	})
//	if err!=nil{
//		 panic(err)
//	}
//	claims,ok:=token.Claims.(*Myclaims)
//	if !ok{
//		panic("token is valid")
//	}
//	jwt.TimeFunc=time.Now
//	claims.StandardClaims.ExpiresAt =time.Now().Add(5*time.Minute).Unix()
//	return Gentoken()
//
//}