package config

// Redis 配置
type redisConfig struct {
	Addr     string
	Password string
	Db       int
}

var Redis = redisConfig{
	Addr:     "192.168.0.104:6379",
	Password: "123456",
	Db:       5,
}

// 发送邮件配置
type emailConfig struct {
	Host     string // smtp服务器地址
	Port     int    // smtp服务器端口
	Username string // 邮箱用户名
	From     string // 发件人地址
	Password string // 邮箱密钥

}

var Emial = emailConfig{
	Host:     "smtp.163.com",
	Port:     465,
	Username: "阿新网",
	From:     "xu756top@163.com",
	Password: "CZVXGXVNKUIXNDHX",
}
