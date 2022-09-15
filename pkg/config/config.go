package config

type config struct {
	Filenme string
	Error   error
}

func (c *config) Init() {
	options, err := configParser(c.Filenme)
	if err != nil {
		c.Error = err
		return
	}
	if (options != Options{}) {
		Cfg.Set(&options)
	}
}

var Configure = getConfig()

func getConfig() *config {
	return &config{
		Filenme: "configs/application.yml",
		Error:   nil,
	}
}
