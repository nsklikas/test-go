package main

import (
	"test-go-server/cmd/server"
)

func main() {
	s := server.New()
	s.Run()
}
