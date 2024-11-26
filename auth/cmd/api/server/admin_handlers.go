package server

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/cmd/api/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (s *Server) adminHealthHandler(c *gin.Context) {
	log.Printf("Health check from %s", c.Request.RemoteAddr)
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) GetUsersHandler(c *gin.Context) {
	log.Printf("Geting all users and their roles for %s", c.Request.RemoteAddr)

	users, err := service.GetUsersService(s.db)
	if err != nil {
		log.Printf("Error getting users: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (s *Server) DeleteUserHandler(c *gin.Context) {
	log.Printf("Geting user and their role for %s", c.Request.RemoteAddr)

	id, err := service.DeleteUserService(s.db, c.Param("id"))
	if err != nil {
		log.Printf("Error getting user: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User with id: "+(id)+" deleted successfully")
}

func (s *Server) ChangeRoleHandler(c *gin.Context) {
	user, ok := c.Get("user_id")
	if !ok {
		log.Printf("Error getting user from context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user id from context"})
		return
	}
	log.Printf("Changing user role for %s by %s", c.Query("id"), user)

	if c.Query("role") == user {
		log.Printf("Error changing user role: User tried to change his own role")
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't change your own role"})
		return
	}

	id, err := service.ChangeRoleService(s.db, c.Param("id"), c.Query("role"))
	if err != nil {
		log.Printf("Error changing user role: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User with id: "+(id)+" role changed to "+c.Query("role")+" successfully")
}
