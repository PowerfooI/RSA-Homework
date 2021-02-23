package main

import (
	"backend/server"
)

func main() {
	s := server.NewServer()

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
