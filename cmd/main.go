package main

import (
	"github.com/jaem/maple"
)

func main() {
	config := maple.NewConfig()
	maple.InitDatabase(config)
	maple.StartServer(config)


	//user := accounts.User{ Name: "slene", Password: "hello" }
	//accounts.CreateUser(&user)

}