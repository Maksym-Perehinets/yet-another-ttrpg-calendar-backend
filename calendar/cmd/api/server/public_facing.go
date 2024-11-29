package server

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/cmd/api/server/services"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/shared/transformers"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (s *Server) GetLocationsHandler(c *gin.Context) {
	log.Printf("Request to get locations from %s", c.ClientIP())
	if c.Query("page") == "" || c.Query("limit") == "" {
		c.JSON(400, "Both page and limit are required")
		return
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(400, "Invalid page")
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(400, "Invalid page")
		return
	}

	data, err := services.GetLocationsService(s.db, page, limit)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, data)
}

func (s *Server) GetLocationHandler(c *gin.Context) {
	log.Printf("Request to get location from %s", c.ClientIP())
	if c.Param("id") == "" {
		c.JSON(400, gin.H{"error": "Id is required"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	data, err := s.db.GetLocation(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := transformers.ToLocationResponse(data)

	c.JSON(200, response)
}
