package service

import (
	"fmt"
	"redrock_homework/dao"
	"redrock_homework/model"
)

//注册相关的服务函数
//1.验证注册账号是否重复
func IspeatUsername(username string)bool  {
	sqlstr:=`select username,password from useraccount where username=?`
	if dao.Queryuser(sqlstr,username).Username==""{
		return true
	}else{
		return false
	}

}
//2.将注册的账号插入数据库
func Insert (username string,password string)int64{
	sqlstr:=`insert into useraccount(username,password) values(?,?)`
	n:=dao.Insertdb(sqlstr,username,password)
	return n
}
//登录相关的服务函数
//验证密码是否正确
func IsTrue(username string,password string)bool{
	sqlstr:=`select username,password from useraccount where username=?`
	 if dao.Queryuser(sqlstr,username).Password==password&&dao.Queryuser(sqlstr,username).Password!=""{
		 return true
	 }else{
		 return false
	 }
}
//查询用户信息
func QueryInfo(id string)model.Userinfo{
	sqlstr:=`select id,avatar,nickname,introduction,phone ,qq,gender,email,birthday from userinfo where id=?`
	u:=model.Userinfo{}
	row:=dao.Queryone(sqlstr,id)
	err:=row.Scan(&u.User_id,&u.Avatar,&u.Nickname,&u.Introduction,&u.Phone,&u.Qq,&u.Gender,&u.Email,&u.Birthday)
	if err!=nil{
		fmt.Println(err)
	}
	return u
}
//更改用户信息
func UpdateInfo(sqlstr string,newdata string,olddata string)int64  {
	return dao.Updatedb(sqlstr,newdata,olddata)
}