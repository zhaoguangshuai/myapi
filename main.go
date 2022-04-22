package main

import (
	"fmt"                      // init 执行顺序描述
	"gohub/app/cmd"            // 该包依赖 3 2 1 包 所以执行顺序为 1 2 3 的init 4 和 5 的 init 按照顺序执行就可以了
	make2 "gohub/app/cmd/make" // 4
	"gohub/bootstrap"          // 3
	btsConig "gohub/config"    // 5
	"gohub/pkg/config"         // 2
	"gohub/pkg/console"        // 1
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConig.Initialize()
}

func main() {
	fmt.Println("main.go", 111111)
	// todo 配置信息没有初始化，获取为空
	fmt.Println("main.go", "app.name =>", config.Get("app.name"))
	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"), // todo 配置信息没有初始化，获取不到 Use 始终为空
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			fmt.Println("main.go", "cmd.env", cmd.Env)
			config.InitConfig(cmd.Env)
			// todo 配置信息初始化后，才可以获取到
			fmt.Println("main.go", "app.name =>", config.Get("app.name"))
			fmt.Println("main.go", 777777)
			// 初始化 Logger
			bootstrap.SetupLogger()
			fmt.Println("main.go", 888888)
			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
		},
	}
	fmt.Println("main.go", 333333)
	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make2.CmdMake,
	)

	// 配置默认运行 Web 服务
	fmt.Println("rootCmd.use=>", rootCmd.Use)
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
