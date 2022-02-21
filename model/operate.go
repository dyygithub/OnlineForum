package model

type Focus struct {
	focuser string`json:"collecter"`
    User_id string`json:"user_id"`
	Avatar string`json:"avater"`
	Nickname string`json:"nickname"`
	Introduction string`json:"introduction"`
}
type Collection struct {
	collecter string`json:"collecter"`
    Post_id int`json:"post_id"`
	Publish_time string`json:"publish_time"`
	User_id string`json:"user_id"`
	Avatar string`json:"avatar"`
	Nickname string`json:"nickname"`
}
