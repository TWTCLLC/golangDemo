package main

import (
	"demo/config"
	"demo/controller"
)

func main() {
	config.ConfigInit()
	controller.WebServiceStart()
}
