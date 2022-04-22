module gohub

go 1.17

require (
	github.com/KenmyZhang/aliyun-communicate v0.0.0-20180308134849-7997edc57454 // 调用阿里云发送短信的安装包
	github.com/gin-gonic/gin v1.7.7 // web 框架 gin
	github.com/go-redis/redis/v8 v8.11.5 // 操作 redis 的安装包
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible // 发送邮件的安装包
	github.com/mojocn/base64Captcha v1.3.5 // 生成图形验证码安装包
	github.com/spf13/cast v1.4.1 // 修改 .env 配置信息的类型使用
	github.com/spf13/viper v1.11.0 // 解析 .env 配置信息使用
	github.com/thedevsaddam/govalidator v1.9.10 // 表单验证器包
	go.uber.org/zap v1.21.0 // Uber 开源的日志工具 zap
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // 数据加密包（密码加密）
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // lumberjack 是一套滚动日志的实现方案，帮助我们管理日志文件
	gorm.io/driver/mysql v1.3.3 // mysql 驱动
	gorm.io/driver/sqlite v1.3.1 // sqlite 驱动
	gorm.io/gorm v1.23.4 // 数据库操作包
)

require (
	github.com/bxcodec/faker/v3 v3.8.0 // 假数据工厂包
	github.com/gertd/go-pluralize v0.2.0 // 字符串辅助方法 单数转复数 复数转单数
	github.com/golang-jwt/jwt v3.2.2+incompatible // 接口授权辅助包
	github.com/iancoleman/strcase v0.2.0 // 字符串辅助方法 驼峰字符串转下划线 下划线转驼峰字符串
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // 命令行高亮打印日志
	github.com/spf13/cobra v1.4.0 // go 命令行依赖包
	github.com/ulule/limiter/v3 v3.10.0 // 限流器依赖包
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pelletier/go-toml/v2 v2.0.0-beta.8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
