package model

type Post struct {
	Post_id int`json:"post_id"`
	Publish_time string `json:"publish_time"`
	Content string `json:"content"`
	Pictures string `json:"pictures"`
	Topic_id string `json:"topic_id"`
	User_id string `json:"user_id"`
	Avatar string `json:"avatar"`
	Nickname string `json:"nickname"`
	Praise_count int64 `json:"praise_count"`
	Comment_count int64 `json:"comment_count"`
}
