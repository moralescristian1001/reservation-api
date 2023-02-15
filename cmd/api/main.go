package main

import (
	"fmt"
	"reservation-api/internal/app"
)

func recoverPanicOnStartup() {
	if r := recover(); r != nil {
		fmt.Println(fmt.Sprintf("panic '%v'", r))
	}
}

func main() {
	defer recoverPanicOnStartup()
	app.Start()
}
