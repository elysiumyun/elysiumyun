package config

import (
	"os"
	"sync"

	"github.com/elysiumyun/elysium/pkg/utils"
)

var (
	addr     string = os.Getenv("ElysiumAddr")
	port     string = os.Getenv("ElysiumPort")
	timezone string = os.Getenv("ElysiumTZ")
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

func getCfg() *cfg {
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

var Cfg *cfg = getCfg()
