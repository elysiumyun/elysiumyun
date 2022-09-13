package config

import (
	"os"
	"sync"

	"github.com/elysiumyun/elysium/pkg/utils"
)

type cfg struct {
	mutex   sync.RWMutex
	options Options
}

func (cfg *cfg) Get() *Options {
	cfg.mutex.TryRLock()
	defer cfg.mutex.RUnlock()
	return &cfg.options
}

func (cfg *cfg) Set(newOptions *Options) {
	cfg.mutex.TryLock()
	defer cfg.mutex.Unlock()
	cfg.options = *newOptions
}

func Cfg() *cfg {
	addr := os.Getenv("")
	port := os.Getenv("")
	timezone := os.Getenv("")
	return &cfg{
		options: Options{
			App: App{
				Addr:     utils.TCO(addr == "", "localhost", addr).(string),
				Port:     utils.TCO(port == "", "8080", port).(string),
				TimeZone: utils.TCO(timezone == "", "Asia/Shanghai", timezone).(string),
			},
		},
	}
}
