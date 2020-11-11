package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golangv/rsa"
)

type Server struct {
	g *gin.Engine
}

func NewServer() *Server {
	s := new(Server)
	s.g = gin.New()
	s.g.Use(cors.Default())

	s.g.GET("keys", s.getKeys)
	s.g.POST("encode", s.encodeMsg)
	s.g.POST("decode", s.decodeMsg)
	s.g.POST("sign", s.signMsg)
	s.g.POST("verify", s.verifySignature)

	return s
}

func (s *Server) Run() error {
	rsa.SetUp()
	return s.g.Run()
}
