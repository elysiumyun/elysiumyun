package main

import (
	"log"
	"os"

	"github.com/elysiumyun/elysium/internal/app"
	"github.com/elysiumyun/elysium/pkg/info"
)

func init() {
	if os.Getenv("ElysiumMode") == "debug" {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
	}
	info.Version.Print()
}

func main() {
	// bootstrap
	code, err := app.App.Run()
	defer os.Exit(code)
	if err != nil {
		log.Println(err)
	}
}
