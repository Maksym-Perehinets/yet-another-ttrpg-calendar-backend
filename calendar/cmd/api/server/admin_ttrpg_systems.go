package server

import (
	"fmt"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/shared/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) CreateSystemHandler(c *gin.Context) {
	var system request.CreateTTRPGSystem

	if err := c.BindJSON(&system); err != nil {
		log.Printf("Error binding TTRPG System: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := s.db.CreateTTRPGSystem(system.ToModel())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": fmt.Sprintf("TTRPG System created successfully with id: %s",
		strconv.FormatUint(uint64(id), 10))})

}
