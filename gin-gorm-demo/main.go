package main

import (
	"gin-gorm-demo/model"
	"gin-gorm-demo/router"
)

func main() {
	model.Connect()
	r := router.App()
	err := r.Run()
	if err != nil {
		return
	}

}
