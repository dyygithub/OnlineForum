package service

import (
	"fmt"
	"redrock_homework/dao"
	"redrock_homework/model"
)
//点赞操作
func Praise(model int,target_id int){
	var praise_count int
	if model==1{
		sqlstr1:=`select praise_count from post where post_id=?`
		dao.Queryone(sqlstr1,target_id).Scan(&praise_count)
		sqlstr2:=`update post set praise_count=? where post_id=?`
		praise_count++
		dao.Updatedb(sqlstr2,praise_count,target_id)
	}
	if model ==2{
		sqlstr1:=`select praise_count from responses1 where comment_id=?`
		dao.Queryone(sqlstr1,target_id).Scan(&praise_count)
		sqlstr2:=`update responses1 set praise_count=? where comment_id=?`
		praise_count++
		dao.Updatedb(sqlstr2,praise_count,target_id)
	}
	if model==3{
		sqlstr1:=`select praise_count from responses2 where comment_id=?`
		dao.Queryone(sqlstr1,target_id).Scan(&praise_count)
		sqlstr2:=`update responses2 set praise_count=? where comment_id=?`
		praise_count++
		dao.Updatedb(sqlstr2,praise_count,target_id)
	}
}
//获取用户关注列表
func Getfocuslist(focuser string)[]model.Focus{
	sqlstr:=`select user_id,avatar,nickname,introduction from focus where focuser=?`
	u:=model.Focus{}
	arr:=make([]model.Focus,0)
	rows,err:=dao.QueryMore(sqlstr,focuser)
	defer rows.Close()
	if err!=nil{
		fmt.Println(err)
	}
	for rows.Next(){
		err:=rows.Scan(&u.User_id,&u.Avatar,&u.Nickname,&u.Introduction)
		if err!=nil{
			fmt.Println(err)
		}
		arr=append(arr,u)
	}
	return arr
}
//关注别的用户
func Collectpost(user_id string,focuser string){
	u:=model.Focus{}
	sqlstr1:=`select id,avatar,nickname,introduction from userinfo where id=?`
	dao.Queryone(sqlstr1,user_id).Scan(&u.User_id,&u.Avatar,&u.Nickname,&u.Introduction)
	sqlstr2:=`insert into focus (focuser,user_id,avatar,nickname,introduction)values(?,?,?,?,?)`
	dao.Insertdb(sqlstr2,focuser,u.User_id,u.Avatar,u.Nickname,u.Introduction)
}
//通过关注的user_id 查询user_name
func Queryusername(user_id string) string{
	sqlstr:=`select username from userinfo where id=?`
	var username string
	dao.Queryone(sqlstr,user_id).Scan(&username)
	return username
}
//获取用户收藏列表
func Getcollection(collecter string)[]model.Collection{
	sqlstr:=`select post_id,publish_time,user_id,avatar,nickname from collection where collecter=?`
	u:=model.Collection{}
	arr:=make([]model.Collection,0)
	rows,err:=dao.QueryMore(sqlstr,collecter)
	defer rows.Close()
	if err!=nil{
		fmt.Println(err)
	}
	for rows.Next(){
		err:=rows.Scan(&u.Post_id,&u.Publish_time,&u.User_id,&u.Avatar,&u.Nickname)
		if err!=nil{
			fmt.Println(err)
		}
		arr=append(arr,u)
	}
	return arr
}
//收藏一个帖子
func Collect(post_id int,collecter string){
	u:=model.Collection{}
	sqlstr1:=`select post_id,publish_time,user_id,avatar,nickname from post where post_id=?`
	dao.Queryone(sqlstr1,post_id).Scan(&u.Post_id,&u.Publish_time,&u.User_id,&u.Avatar,&u.Nickname)
	sqlstr2:=`insert into collection (collecter,post_id,publish_time,user_id,avatar,nickname)values(?,?,?,?,?,?)`
	dao.Insertdb(sqlstr2,collecter,u.Post_id,u.Publish_time,u.User_id,u.Avatar,u.Nickname)
}