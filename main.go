package main

import (
	"fmt"
	. "sanjos/auth/controller"
)

func main() {

	fmt.Println("== Simple Auth Service ==")
	InitRepository()
	// TODO : serve your web server here
	InsertUserData()
}
