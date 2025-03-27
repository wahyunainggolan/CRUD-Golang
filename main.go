package main

import (
	"crud/app"
)

func main() {
	var a app.App
	a.CreateConnection()
	a.Routes()
	a.Run()
}