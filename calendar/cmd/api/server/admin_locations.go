package server

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/shared/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) CreateLocationHandler(c *gin.Context) {
	var location request.CreateLocation

	if err := c.BindJSON(&location); err != nil {
		log.Printf("Error binding location: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := s.db.CreateLocation(&models.Locations{
		Name:        location.Name,
		Description: location.Description,
		City:        location.City,
		Street:      location.Street,
		LinkToSite:  location.LinkToSite,
		Price:       location.Price,
		PricingType: location.PricingType,
		OpenAt:      location.OpenAt,
		CloseAt:     location.CloseAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Location created successfully with id: "+strconv.FormatUint(uint64(id), 10))
}

func (s *Server) DeleteLocationHandler(c *gin.Context) {
	log.Printf("Request to delete location from %s", c.ClientIP())
	if c.Param("id") == "" {
		c.JSON(400, "Id is required")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "Invalid id")
		return
	}

	err = s.db.DeleteLocation(id)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, "Location deleted")
}

func (s *Server) UpdateLocationHandler(c *gin.Context) {
	var location request.Update

	if err := c.BindJSON(&location); err != nil {
		log.Printf("Error binding location: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "Invalid id")
		return
	}

	err = s.db.UpdateLocation(id, location)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, "Location updated")
}
