package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) getKeys(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (s *Server) encodeMsg(c *gin.Context) {

}

func (s *Server) decodeMsg(c *gin.Context) {

}

func (s *Server) signMsg(c *gin.Context) {

}
