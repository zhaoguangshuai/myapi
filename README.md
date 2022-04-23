**项目说明**
本项目构建一个论坛API为主题，从零开始构建一个高性能，功能齐全的API程序框架。

**程序结构**

```go
.├── app                            // 程序具体逻辑代码
│   ├── cmd                         // 命令
│   │   ├── cache.go                
│   │   ├── cmd.go
│   │   ├── key.go
│   │   ├── make                    // make 命令及子命令
│   │   │   ├── make.go
│   │   │   ├── make_apicontroller.go
│   │   │   ├── make_cmd.go
│   │   │   ├── make_factory.go
│   │   │   ├── make_migration.go
│   │   │   ├── make_model.go
│   │   │   ├── make_policy.go
│   │   │   ├── make_request.go
│   │   │   ├── make_seeder.go
│   │   │   └── stubs               // make 命令的模板
│   │   │       ├── apicontroller.stub
│   │   │       ├── cmd.stub
│   │   │       ├── factory.stub
│   │   │       ├── migration.stub
│   │   │       ├── model
│   │   │       │   ├── model.stub
│   │   │       │   ├── model_hooks.stub
│   │   │       │   └── model_util.stub
│   │   │       ├── policy.stub
│   │   │       ├── request.stub
│   │   │       └── seeder.stub
│   │   ├── migrate.go
│   │   ├── play.go
│   │   ├── seed.go
│   │   └── serve.go
│   ├── http                        // http 请求处理逻辑
│   │   ├── controllers             // 控制器，存放 API 和视图控制器
│   │   │   ├── api                 // API 控制器，支持多版本的 API 控制器
│   │   │   │   └── v1              // v1 版本的 API 控制器
│   │   │   │       ├── users_controller.go
│   │   │   │       └── ...
│   │   └── middlewares             // 中间件
│   │       ├── auth_jwt.go
│   │       ├── guest_jwt.go
│   │       ├── limit.go
│   │       ├── logger.go
│   │       └── recovery.go
│   ├── models                      // 数据模型
│   │   ├── user                    // 单独的模型目录
│   │   │   ├── user_hooks.go       // 模型钩子文件
│   │   │   ├── user_model.go       // 模型主文件
│   │   │   └── user_util.go        // 模型辅助方法
│   │   └── ...
│   ├── policies                    // 授权策略目录
│   │   ├── category_policy.go
│   │   └── ...
│   └── requests                    // 请求验证目录（支持表单、标头、Raw JSON、URL Query）
│       ├── validators              // 自定的验证规则
│       │   ├── custom_rules.go
│       │   └── custom_validators.go
│       ├── user_request.go
│       └── ...
├── bootstrap                       // 程序模块初始化目录
│   ├── app.go  
│   ├── cache.go
│   ├── database.go
│   ├── logger.go
│   ├── redis.go
│   └── route.go
├── config                          // 配置信息目录
│   ├── app.go
│   ├── captcha.go
│   ├── config.go
│   ├── database.go
│   ├── jwt.go
│   ├── log.go
│   ├── mail.go
│   ├── pagination.go
│   ├── redis.go
│   ├── sms.go
│   └── verifycode.go
├── database                        // 数据库相关目录
│   ├── database.db                 // sqlite 数据文件（加入到 .gitignore 中）
│   ├── factories                   // 模型工厂目录
│   │   ├── user_factory.go
│   │   └── ...
│   ├── migrations                  // 数据库迁移目录
│   │   ├── 2021_12_21_102259_create_users_table.go
│   │   ├── 2021_12_21_102340_create_categories_table.go
│   │   └── ...
│   └── seeders                     // 数据库填充目录
│       ├── users_seeder.go
│       ├── ...
├── pkg                             // 内置辅助包
│   ├── app
│   ├── auth
│   ├── cache
│   ├── captcha
│   ├── config
│   └── ...
├── public                          // 静态文件存放目录
│   ├── css
│   ├── js
│   └── uploads                     // 用户上传文件目录
│       └── avatars                 // 用户上传头像目录
├── routes                          // 路由
│   ├── api.go
│   └── web.go
├── storage                         // 内部存储目录
│   ├── app
│   └── logs                        // 日志存储目录
│       ├── 2021-12-28.log
│       ├── 2021-12-29.log
│       ├── 2021-12-30.log
│       └── logs.log
└── tmp                             // air 的工作目录
├── .env                            // 环境变量文件
├── .env.example                    // 环境变量示例文件
├── .gitignore                      // git 配置文件
├── .air.toml                       // air 配置文件
├── .editorconfig                   // editorconfig 配置文件
├── go.mod                          // Go Module 依赖配置文件
├── go.sum                          // Go Module 模块版本锁定文件
├── main.go                         // Gohub 程序主入口
├── Makefile                        // 自动化命令文件
├── readme.md                       // 项目 readme
```


**所有路由**
| 请求方法 | API地址      | 说明      |
|:--------:| :-------------:| :-------------:|
| POST | /api/v1/auth/login/using-phone |短信 + 手机号登录 |
| POST | /api/v1/auth/login/using-password |手机号、用户名、邮箱 + 密码 |
| POST |/api/v1/auth/login/refresh-token |刷新 Token |
| POST |/api/v1/auth/password-reset/using-email |邮件密码重置 |
| POST | /api/v1/auth/password-reset/using-phone |短信验证码密码重置 |
| POST | /api/v1/auth/signup/using-phone |使用手机号注册 |
| POST | /api/v1/auth/signup/using-email |使用邮箱注册|
| POST | /api/v1/auth/signup/phone/exist |手机号是否已注册 |
| POST | /api/v1/auth/signup/email/exist |email 是否已支持 |
| POST |/api/v1/auth/verify-codes/phone |发送短信验证码 |
| POST |/api/v1/auth/verify-codes/email |发送邮件验证码 |
| POST | /api/v1/auth/verify-codes/captcha |获取图片验证码 |
| GET | /api/v1/user |获取当前用户 |
| GET | /api/v1/users |用户列表 |
| PUT | /api/v1/users |修改个人资料 |
| PUT | /api/v1/users/email |修改邮箱 |
| PUT | /api/v1/users/phone |修改手机号 |
| PUT | /api/v1/users/password |修改密码 |
| PUT |/api/v1/users/avatar |上传头像 |
| GET |/api/v1/categories |分类列表 |
| POST |/api/v1/categories |创建分类 |
| PUT |	/api/v1/categories/:id |更新分类 |
| DELETE |/api/v1/categories/:id |删除分类 |
| GET |/api/v1/topics |话题列表 |
| POST |/api/v1/topics |创建话题 |
| PUT |/api/v1/topics/:id |更新话题 |
| DELETE |/api/v1/topics/:id | 删除话题 |
|GET |/api/v1/topics/:id |获取话题 |
| GET |/api/v1/links |友情链接列表 |

**第三方依赖**
使用到的开源库：

- gin —— 路由、路由组、中间件
- zap —— 高性能日志方案
- gorm —— ORM 数据操作
- cobra —— 命令行结构
- viper —— 配置信息
- cast —— 类型转换
- redis —— Redis 操作
- jwt —— JWT 操作
- base64Captcha —— 图片验证码
- govalidator —— 请求验证器
- limiter —— 限流器
- email —— SMTP 邮件发送
- aliyun-communicate —— 发送阿里云短信
- ansi —— 终端高亮输出
- strcase —— 字符串大小写操作
- pluralize —— 英文字符单数复数处理
- faker —— 假数据填充
- imaging —— 图片裁切

**自定义的包**

- app —— 应用对象
- auth —— 用户授权
- cache —— 缓存
- captcha —— 图片验证码
- config —— 配置信息
- console —— 终端
- database —— 数据库操作
- file —— 文件处理
- hash —— 哈希
- helpers —— 辅助方法
- jwt —— JWT 认证
- limiter —— API 限流
- logger —— 日志记录
- mail —— 邮件发送
- migrate —— 数据库迁移
- paginator —— 分页器
- redis —— Redis 数据库操作
- response —— 响应处理
- seed —— 数据填充
- sms —— 发送短信
- str —— 字符串处理
- verifycode —— 数字验证码

**所有命令**

```go
go run main.go -h
Default will run "serve" command, you can use "-h" flag to see all subcommands

Usage:
   [command]

Available Commands:
  cache       Cache management
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  key         Generate App Key, will print the generated Key
  make        Generate file and code
  make        Generate file and code
  migrate     Run database migration
  play        Likes the Go Playground, but running at our application context
  seed        Insert fake data to the database
  serve       Start web server

Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file
  -h, --help         help for this command

Use " [command] --help" for more information about a command.
```
make 命令：

```go
go run main.go make -h
Generate file and code

Usage:
   make [command]

Available Commands:
  apicontroller Create api controller，exmaple: make apicontroller v1/user
  cmd           Create a command, should be snake_case, exmaple: make cmd buckup_database
  factory       Create model's factory file, exmaple: make factory user
  migration     Create a migration file, example: make migration add_users_table
  model         Crate model file, example: make model user
  policy        Create policy file, example: make policy user
  request       Create request file, example make request user
  seeder        Create seeder file, example:  make seeder user

Flags:
  -h, --help   help for make

Global Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file

Use " make [command] --help" for more information about a command.
```
migrate 命令：

```go
go run main.go migrate -h
Run database migration

Usage:
   migrate [command]

Available Commands:
  down        Reverse the up command
  fresh       Drop all tables and re-run all migrations
  refresh     Reset and re-run all migrations
  reset       Rollback all database migrations
  up          Run unmigrated migrations

Flags:
  -h, --help   help for migrate

Global Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file

Use " migrate [command] --help" for more information about a command.
```
