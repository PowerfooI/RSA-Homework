package server

import (
	"fmt"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer()

	err := s.Run()
	if err != nil {
		fmt.Println(err)
	}
}
