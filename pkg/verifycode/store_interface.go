package verifycode

// Store 验证码模块不应该跟特定的存储服务做绑定，所以我们创建了
//Store interface 。未来如果要使用其他内存服务器（例如 Memcached），可以很方便切换
// 创建接口 Store
type Store interface {
	// Set 保存验证码就
	Set(id, value string) bool

	// Get 获取验证码
	Get(id string, clear bool) string

	// Verify 检查验证码
	Verify(id, answer string, clear bool) bool
}
