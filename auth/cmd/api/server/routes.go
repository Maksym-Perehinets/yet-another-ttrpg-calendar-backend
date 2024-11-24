package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	auth := r.Group("/v1/auth")
	{
		auth.GET("/health", s.healthHandler)
		auth.PUT("/sign-in", s.RegisterHandler)
		auth.POST("/login", s.LoginHandler)
		//r.GET("/refresh", s.RefreshHandler)
		//r.GET("/logout", s.LogoutHandler)
	}

	admin := r.Group("/v1/github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/admin")
	admin.Use(AuthMiddleware("admin"))
	{
		admin.GET("/health", s.adminHealthHandler)
		admin.GET("/users", s.GetUsersHandler)
		admin.DELETE("/users/:id", s.DeleteUserHandler)
		admin.PUT("/change-permission", s.ChangeRoleHandler)
	}
	return r
}
