package main

import (
	"github.com/elysiumyun/elysium/internal/app"
)

func main() {
	if err := app.App(); err != nil {
		panic(err)
	}
}
