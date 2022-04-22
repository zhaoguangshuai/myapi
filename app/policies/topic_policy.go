// Package policies 用户授权
package policies

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/topic"
	"gohub/pkg/auth"
)

func CanModifyTopic(c *gin.Context, _Topic topic.Topic) bool {
	return auth.CurrentUID(c) == _Topic.UserID
}
