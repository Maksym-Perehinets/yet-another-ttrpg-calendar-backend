package server

import (
	commonAuth "github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/auth"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	cal := r.Group("/v1/calendar")
	{
		cal.GET("/", s.HelloWorldHandler)

		// Locations
		cal.GET("/locations", s.GetLocationsHandler)
		cal.GET("/location/:id", s.GetLocationHandler)
	}

	users := r.Group("/v1/calendar/authn")
	users.Use(commonAuth.AuthMiddleware(""))
	{
		users.GET("/", s.HelloWorldHandler)
	}

	admin := r.Group("/v1/calendar/admin")
	admin.Use(commonAuth.AuthMiddleware("admin"))
	{
		admin.GET("/health", s.adminHealthHandler)

		// Locations
		admin.POST("/location", s.CreateLocationHandler)
		admin.DELETE("/location/:id", s.DeleteLocationHandler)
		admin.PATCH("/location/:id", s.UpdateLocationHandler)
	}

	return r
}
