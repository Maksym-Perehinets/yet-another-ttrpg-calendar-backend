package server

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/cmd/api/server"
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
		AllowCredentials: true, // Enable cookies/auth
	}))

	cal := r.Group("/v1/calendar")
	{
		cal.GET("/", s.HelloWorldHandler)
		//cal.GET("/health", s.healthHandler)
	}

	users := r.Group("/v1/calendar/authn")
	users.Use(server.AuthMiddleware(""))
	{
		users.GET("/", s.HelloWorldHandler)
	}

	admin := r.Group("/v1/calendar/admin")
	admin.Use(server.AuthMiddleware("admin"))
	{
		admin.GET("/health", s.HelloWorldHandler)
	}

	return r
}
