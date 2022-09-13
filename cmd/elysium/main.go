package main

import (
	"os"

	"github.com/elysiumyun/elysium/internal/app"
	"github.com/elysiumyun/elysium/pkg/info"
)

func init() {
	info.PrintAll()
}

func main() {
	var err error
	var code int
	// bootstrap
	code, err = app.App.Run()
	defer os.Exit(code)
	if err != nil {
		panic(err)
	}
}
