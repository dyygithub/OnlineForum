package service

import (
	"fmt"
	"redrock_homework/dao"
	"redrock_homework/model"
)

//获取帖子列表的服务函数
func List(page int,array []model.Post,size int)([]model.Post,error) {
	sqlstr := `select post_id,publish_time,content,pictures,topic_id,user_id,avatar,nickname,praise_count,comment_count from post `
	u:= model.Post{}
	rows, err := dao.QueryMore(sqlstr)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err:=rows.Scan(&u.Post_id, &u.Publish_time, &u.Content, &u.Pictures, &u.Topic_id, &u.User_id, &u.Avatar, &u.Nickname, &u.Praise_count, &u.Comment_count)
		if err != nil {
			return nil, err
		}
		array=append(array,u)
	}
	return array[(page-1)*size:page*size],nil
}
//获取某个帖子内容的服务函数
func Single(post_id int)(model.Post,error){
	u:=model.Post{}
	sqlstr:=`select post_id,publish_time,content,pictures,topic_id,user_id,avatar,nickname,praise_count,comment_count from post where post_id=?`
	row:=dao.Queryone(sqlstr,post_id)
	err:=row.Scan(&u.Post_id, &u.Publish_time, &u.Content, &u.Pictures, &u.Topic_id, &u.User_id, &u.Avatar, &u.Nickname, &u.Praise_count, &u.Comment_count)
	if err!=nil{
		return u,err
	}
	return u,nil
}
//创建一个帖子
func CreatePost(publish_time string,content string,pictures string,topic_id string,user_id string,avatar string,nickname string)int64{
	sqlstr:="insert into post (publish_time,content,pictures,topic_id,user_id,avatar,nickname)values(?,?,?,?,?,?,?)"
	ret,err:=dao.Db.Exec(sqlstr,publish_time,content,pictures,topic_id,user_id,avatar,nickname)
	if err!=nil{
		fmt.Println(err)
		return 0
	}
	n,err:=ret.RowsAffected()
	if err!=nil{
		fmt.Println(err)
		return 0
	}
	return n
}
//返回新创建帖子的post—id
func Getpost_id(content string)int {
	sqlstr:=`select post_id from post where content=?`
	var post_id int
	dao.Queryone(sqlstr,content).Scan(&post_id)
	return post_id
}
//删除一条帖子
func Deletepost(post_id int)int64{
	sqlstr:=`delete from post where post_id=?`
    return dao.Deletedb(sqlstr,post_id)
}
//删除帖子要删除其下面的评论
func Deletecomment(post_id int)  {
	sqlstr1:=`delete from responses1 where post_id=?`
	sqlstr2:=`delete from responses2 where post_id=?`
	dao.Deletedb(sqlstr1,post_id)
	dao.Deletedb(sqlstr2,post_id)
}
//删除对于这个帖子的收藏数据
func Deletecollection(post_id int){
	sqlstr:=`delete from colletion where post_id=?`
	dao.Deletedb(sqlstr,post_id)
}
//搜索帖子
func Searchpost(page int ,array []model.Post,size int,key string)([]model.Post,error){
	sqlstr := "select post_id,publish_time,content,pictures,topic_id,user_id,avatar,nickname,praise_count,comment_count from post where content like CONCAT('%',?,'%')"
	u:= model.Post{}
	rows, err := dao.QueryMore(sqlstr,key)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err:=rows.Scan(&u.Post_id, &u.Publish_time, &u.Content, &u.Pictures, &u.Topic_id, &u.User_id, &u.Avatar, &u.Nickname, &u.Praise_count, &u.Comment_count)
		if err != nil {
			return nil, err
		}
		array=append(array,u)
	}
	return array[(page-1)*size:page*size],nil
}