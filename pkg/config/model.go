package config

// options
type Options struct {
	App   `yaml:"app"`   // App 运行设置
	Mysql `yaml:"mysql"` // Mysql数据库配置
}

// app env config
type App struct {
	Addr     string `yaml:"addr"`     // 地址
	Port     string `yaml:"port"`     // 端口
	TimeZone string `yaml:"timezone"` // 时区
}

type Mysql struct {
	Addr     string `yaml:"addr"`     // 地址
	Port     string `yaml:"port"`     // 端口
	User     string `yaml:"user"`     // 用户
	Passwd   string `yaml:"passwd"`   // 密码
	Database string `yaml:"database"` // 库名
	Charset  string `yaml:"charset"`  // 编码
}
