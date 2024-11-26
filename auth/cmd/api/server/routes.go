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
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.GET("/health", s.healthHandler)
			auth.POST("/sign-in", s.RegisterHandler)
			auth.POST("/login", s.LoginHandler)
			//r.GET("/refresh", s.RefreshHandler)
			//r.GET("/logout", s.LogoutHandler)
		}

		admin := v1.Group("/auth/admin")
		admin.Use(commonAuth.AuthMiddleware("admin"))
		{
			admin.GET("/health", s.adminHealthHandler)
			admin.GET("/users", s.GetUsersHandler)
			admin.DELETE("/users/:id", s.DeleteUserHandler)
			admin.POST("/users/:id/permissions", s.ChangeRoleHandler)
		}

		authorized := v1.Group("/user")
		authorized.Use(commonAuth.AuthMiddleware(""))
		{
			authorized.PATCH("/:id", s.UpdateUserHandler)
		}
	}
	return r
}
