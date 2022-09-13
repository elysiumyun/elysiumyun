package config

// options
type Options struct {
	App App `yaml:"app,omitempty"`
}

// app env config
type App struct {
	Addr     string `yaml:"addr"`
	Port     string `yaml:"port"`
	TimeZone string `yaml:"timezone"`
}
