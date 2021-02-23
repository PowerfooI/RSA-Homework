package server

import (
	"backend/rsa"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	g *gin.Engine
}

func NewServer() *Server {
	s := new(Server)
	gin.SetMode("release")
	s.g = gin.New()
	s.g.Use(cors.Default())
	s.g.Static("/js", "./dist/js")
	s.g.Static("/css", "./dist/css")
	s.g.Static("/fonts", "./dist/fonts")
	s.g.StaticFile("/", "./dist/index.html")

	group := s.g.Group("/api")
	group.GET("/keys", s.getKeys)
	group.POST("/encode", s.encodeMsg)
	group.POST("/decode", s.decodeMsg)
	group.POST("/sign", s.signMsg)
	group.POST("/verify", s.verifySignature)

	//s.g.Any("/", func(c *gin.Context){
	//	c.Request.URL.Path = "/s/index.html"  //把请求的URL修改
	//	s.g.HandleContext(c)  //继续后续处理
	//})

	return s
}

func (s *Server) Run() error {
	rsa.SetUp()
	fmt.Println("* 请用浏览器打开 http://localhost:8123/ 使用RSA相关应用")
	return s.g.Run(":8123")
}
