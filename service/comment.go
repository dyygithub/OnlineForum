package service

import (
	"fmt"
	"redrock_homework/dao"
	"redrock_homework/model"
)
//获取某一条帖子或评论的评论
func Getcomment(model_ int,target_id int,page int,size int)([]model.Responses1,[]model.Responses2) {
	var array1 []model.Responses1
	var array2 []model.Responses2
	if model_ == 1 {
		sql := `select comment_id,post_id,publish_time,content,pictures,user_id,avatar,nickname,praise_count from responses1 where and post_id=?`
		rows, err := dao.QueryMore(sql, target_id)
		defer rows.Close()
		if err != nil {
			fmt.Println(err)
		}
		u1 := model.Responses1{}
		for rows.Next(){
			err := rows.Scan(&u1.Comment_id, &u1.Post_id, &u1.Publish_time, &u1.Content, &u1.Pictures, &u1.User_id, &u1.Avatar, &u1.Nickname, &u1.Praise_count)
			if err != nil {
				fmt.Println(err)
			}
			array1 = append(array1, u1)
		}
		return array1[(page-1)*size:page*size],nil
	} else {
		sql := `select comment_id,post_id,publish_time,content,reply_user_id,avatar,reply_nickname,praise_count,from responses2 where and taget_id=?`
		rows, err := dao.QueryMore(sql, target_id)
		defer rows.Close()
		if err != nil {
			fmt.Println(err)
		}
		u2 := model.Responses2{}
		for  rows.Next() {
			err := rows.Scan(&u2.Comment_id, &u2.Post_id, &u2.Publish_time, &u2.Content, &u2.Reply_user_id, &u2.Avatar, &u2.Reply_nickname, &u2.Praise_count)
			if err != nil {
				fmt.Println(err)
			}
			array2 = append(array2,u2)
		}
		return nil,array2[(page-1)*size:page*size]
	}
}
//发表二级评论时通过一级评论的id得到post_id
func Attainpost_id(target_id int) int{
	sqlstr:=`select post_id from responses1 where comment_id=?`
	u:=model.Responses2{}
	err:=dao.Queryone(sqlstr,target_id).Scan(&u.Post_id)
	if err!=nil{
		fmt.Println(err)
	}
	return u.Post_id
}
//将发布的评论的相关信息插入数据库
func Publishcomment(model int,post_id int,target_id int,content string,photo string,nickname string,user_id string,avatar string,publish_time string,praise_count int)int64{
	if model==1{
		sqlstr:=`insert into reseponses1 (post_id,content,pictures,nickname,user_id,avatar,publish_time,praise_count)values(?,?,?,?,?,?,?,?)`
		return dao.Insertdb(sqlstr,target_id,content,photo,nickname,user_id,avatar,publish_time,praise_count)
	}else{
		sqlstr:=`insert into reseponses2 (post_id,target_id,content,reply_nickname,reply_user_id,avatar,publish_time,praise_count)values(?,?,?,?,?,?,?,?)`
		return dao.Insertdb(sqlstr,post_id,target_id,content,nickname,user_id,avatar,publish_time,praise_count)
	}
}
//获取刚发布评论的comment_id
func Getcomment_id(content string,model_ int)int{
	sqlstr1:=`select comment_id from response1 where content=?`
	sqlstr2:=`select comment_id from response2 where content=?`
	comment_id:=0
	if model_==1{
		dao.Queryone(sqlstr1,content).Scan(&comment_id)
	}else{
		dao.Queryone(sqlstr2,content).Scan(&comment_id)
	}
	return comment_id
}
//发布评论后相关帖子的评论数加一
func Add(post_id int){
	sqlstr1:=`select comment_count from post where post_id=?`//先查到原来的评论数
	u:=model.Post{}
	err:=dao.Queryone(sqlstr1,post_id).Scan(&u.Comment_count)
	if err!=nil{
		fmt.Println(err)
	}
	n:=u.Comment_count+1
	sqrstr2:=`update post set comment_count=? where post_id=?`//将该帖子的评论数改变
	dao.Updatedb(sqrstr2,n,post_id)
}
//确定评论是一级评论还是二级评论
func Confirm1(comment_id int,nickname string,user_id string,avatar string)string{
	sqlstr:=`select publish_time from responses1 where comment_id=?and nickname=? and user_id=? and avatar=?`
	u:=model.Responses1{}
	dao.Queryone(sqlstr,comment_id,nickname,user_id,avatar).Scan(&u.Publish_time)
	return u.Publish_time
}
func Confirm2(comment_id int,nickname string,user_id string,avatar string)string{
	sqlstr:=`select publish_time from responses2 where comment_id=?and nickname=? and user_id=? and avatar=?`
	u:=model.Responses2{}
	dao.Queryone(sqlstr,comment_id,nickname,user_id,avatar).Scan(&u.Publish_time)
	return u.Publish_time
}
func Updatecomment1(photo string,content string,comment_id int)int64{
	sqlstr:=`update responses1 set photo=?,content=? where comment_id=?`
	return dao.Updatedb(sqlstr,photo,content,comment_id)
}
func Updatecomment2(content string,comment_id int)int64{
	sqlstr:=`update responses2 set content=? where comment_id=?`
	return dao.Updatedb(sqlstr,content,comment_id)
}
//删除一级评论要删除其对应的二级评论
//删除一级评论还需要删除与其关联的二级评论
func Deletecomment1(comment_id int){
	sqlstr1:=`delete from responses1 where comment_id=?`
	sqlstr2:=`delete from responses2 where target_id=?`
	dao.Deletedb(sqlstr1,comment_id)
	dao.Deletedb(sqlstr2,comment_id)
}
//删除二级评论就只需要删除该评论
func Deletecomment2(comment_id int){
	sqlstr:=`delete from responses2 where comment_id=?`
	dao.Deletedb(sqlstr,comment_id)
}