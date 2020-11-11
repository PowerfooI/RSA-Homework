package main

import (
	"golangv/server"
)

func main() {
	s := server.NewServer()

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
