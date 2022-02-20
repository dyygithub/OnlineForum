package api

import (
	"github.com/gin-gonic/gin"
	"redrock_homework/model"
)

//获取所有主题
func Gettopic()gin.HandlerFunc{
	return func(c *gin.Context) {
		var u1 = model.Topic{
			Topic_id: "1",
			Logo_url: "www.life.com",
			Topic_name: "生活杂谈",
			Introduction: "分享你的生活",
		}
		var u2 =model.Topic{
			Topic_name: "2",
			Logo_url: "www.jishu.com",
			Topic_id: "技术",
			Introduction: "快来和大神学习技术吧",
		}
		var topic[]model.Topic
		topic=append(topic,u1,u2)
		data:=gin.H{
			"post":topic,
		}
		c.JSON(200,gin.H{
			"status":10000,
			"info":"success",
			"data":data,
		})
	}
}