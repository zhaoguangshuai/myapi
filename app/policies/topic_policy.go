package policies

import (
    "gohub/app/models/topic"
    "gohub/pkg/auth"

    "github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, topic topic.Topic) bool {
    return auth.CurrentUID(c) == topic.UserID
}

// func CanViewTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanCreateTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanUpdateTopic(c *gin.Context, topicModel topic.Topic) bool {}
// func CanDeleteTopic(c *gin.Context, topicModel topic.Topic) bool {}
