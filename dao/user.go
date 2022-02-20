package dao

import (
	"fmt"
	"redrock_homework/model"
)

//对数据库进行user相关的查询操作
func Queryuser(sqlstr string,element string)model.UserAccount{
	u:=model.UserAccount{}
	row:=Queryone(sqlstr,element)
	err:=row.Scan(&u.Username,&u.Password)
	if err!=nil{
		fmt.Println(err)
	}
	return  u
}
