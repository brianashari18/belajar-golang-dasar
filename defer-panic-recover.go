package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("End App")
	if message := recover(); message != nil {
		fmt.Println(message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups error")
	}
}

func main() {
	runApp(false)
	fmt.Println("TRIGGER")
}
