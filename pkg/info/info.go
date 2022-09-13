package info

import "fmt"

// app build context
var verStr string
var tagStr string
var uptStr string
var envStr string

var AppName string

func PrintAll() {
	fmt.Println("App Version Info")
	Version.Print()
	fmt.Println("App Runtime Info")
	Runtime.Print()
}
