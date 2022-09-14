package app

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/elysiumyun/elysium/internal/pkg/usage"
	"github.com/elysiumyun/elysium/pkg/info"
)

type app struct {
	success int
	failure int
}

var App = GetApp()

func GetApp() *app {
	return &app{
		success: 0,
		failure: 137,
	}
}

func (app *app) Run() (int, error) {
	var err error

	// process app flags
	if len(os.Args) > 1 {
		var argv = os.Args[1:]
		var argc = len(os.Args) - 1
		err = app.flags(argc, argv)
	} else {
		err = elysium()
	}
	if err != nil {
		return app.failure, err
	}
	return app.success, nil

}

func (app *app) flags(argc int, argv []string) error {
	var err error
	switch argv[0] {
	case "-c", "config", "--config":
		log.Println("^-^")
	case "-h", "--help", "help":
		usage.Usage()
	case "-v", "--version", "version":
		fmt.Println("App Version Info")
		info.Version.Print()
		fmt.Println("App Runtime Info")
		info.Runtime.Print()
	default:
		err = errors.New("please check usage")
	}
	return err
}
