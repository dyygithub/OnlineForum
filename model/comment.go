package model

//一级评论的model
type Responses1 struct {
	Comment_id int`json:"comment_id"`
	Post_id int`json:"post_id"`
	Publish_time string`json:"publish_time"`
	Content string`json:"content"`
	Pictures string`json:"pictures"`//一级评论才有
	User_id string`json:"user_id"`//一级评论的用户的id
	Avatar string`json:"avatar"`
	Nickname string`json:"nickname"`//一级评论的用户昵称
	Praise_count int`json:"praise_count"`
}
//二级评论的model
type Responses2 struct {
	Comment_id int`json:"comment_id"`
	Post_id int`json:"post_id"`
	target_id int`json:"taget_id"`//表示该评论所依赖评论的id
	Publish_time string`json:"publish_time"`
	Content string`json:"content"`
	Reply_user_id string`json:"reply_user_id"`//二级评论的用户id
	Avatar string`json:"avatar"`
	Reply_nickname string`json:"reply_nickname"`//二级评论的用户昵称
	Praise_count int`json:"praise_count"`
}
