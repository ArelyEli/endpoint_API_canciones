package main

import (
	"fmt"
	"songs_api/server"
	"time"
	// "songs_api/database"
)

func main() {
	fmt.Println("Hello")
	time.Sleep(30 * time.Second)
	server.Server()

}
