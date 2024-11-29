package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) adminHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
