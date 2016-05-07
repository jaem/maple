package main

import (
	"github.com/jaem/maple"
)

func main() {
	config := maple.NewConfig()
	maple.InitDatabase(config)
	maple.StartServer(config)


}