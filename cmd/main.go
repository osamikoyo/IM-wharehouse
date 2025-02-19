package main

import (
	"fmt"

	"github.com/osamikoyo/IM-wharehouse/internal/app"
)

func main() {
	app,err := app.New()
	if err != nil{
		fmt.Println(err)
	}

	if err = app.Run(); err != nil{
		fmt.Println(err)
	}
}
