package main

import (
	"fmt"
	"songs_api/server"
	"time"
)

func main() {
	fmt.Println("Hello")
	time.Sleep(30 * time.Second)
	server.Server()

}
