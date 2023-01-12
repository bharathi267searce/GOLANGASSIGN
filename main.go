package main

import (
	"fmt"

	// "github.com/1234bharathi/GOLANGASSIGN/Handlers"

	// "github.com/1234bharathi/GOLANGASSIGN/Support"
	application "github.com/1234bharathi/GOLANGASSIGN/Application"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	_ "github.com/lib/pq"
)

func main() {
	Support.DbConnect()
	func() {
		fmt.Println("Enter console for console interface or enter API to work with Postman API")
		var reply string
		_, err := fmt.Scanln(&reply)
		if err != nil {
			fmt.Println("error in reading input!!")
		}
		if reply == "console" {
			application.Console()
		} else {
			fmt.Print("nk")
			application.ServerFunctionCall()
		}
		return

	}()
	// router.HandleFunc("/", HomeLink)
	//product handlers

	// application.ServerFunctionCall()
}
