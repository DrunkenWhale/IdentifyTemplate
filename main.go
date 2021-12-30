package main

import (
	"IdentifyTemplate/web"
	"fmt"
)

func main() {
	app, err := web.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	err1 := app.Run(":2333")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}
