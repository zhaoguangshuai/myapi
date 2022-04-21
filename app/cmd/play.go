package cmd

import (
	"github.com/spf13/cobra"
	"gohub/pkg/console"
	"gohub/pkg/redis"
	"time"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
	// 存放测试代码
	// 存进去 redis 中
	redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	// 从 redis 里获取
	console.Success(redis.Redis.Get("hello"))
}
