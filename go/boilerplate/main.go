package main

import (
	"boilerplate/server"
	"time"
)

func main() {

	time.Sleep(time.Second * 1)

	server.Start()
}
