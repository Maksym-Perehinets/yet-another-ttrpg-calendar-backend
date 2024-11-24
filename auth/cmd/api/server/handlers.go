package server

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/cmd/api/service"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/cmd/api/service/validate"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/shared/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (s *Server) healthHandler(c *gin.Context) {
	log.Printf("Health check from %s", c.Request.RemoteAddr)
	c.JSON(http.StatusOK, "OK")
}

func (s *Server) adminHealthHandler(c *gin.Context) {
	log.Printf("Health check from %s", c.Request.RemoteAddr)
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) RegisterHandler(c *gin.Context) {
	log.Printf("Register request from %s", c.Request.RemoteAddr)

	user := validate.NewRegisterRequest(c)

	// Validate email, username, and password
	if err := user.ValidateEmail(s.db); err != nil {
		log.Printf("Invalid email: %s", user.GetStruct().Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.ValidateUsername(s.db); err != nil {
		log.Printf("Invalid username: %s", user.GetStruct().Username)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.ValidatePassword(); err != nil {
		log.Printf("Invalid password for user: %s", user.GetStruct().Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, cookieExpiresIn, err := service.RegisterService(user.GetStruct())
	if err != nil {
		log.Printf("Error registering user: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.SetCookieHandler("Authorization", token, "/", cookieExpiresIn, c)
	c.JSON(http.StatusOK, "User created successfully")
}

func (s *Server) LoginHandler(c *gin.Context) {
	log.Printf("Login request from %s", c.Request.RemoteAddr)

	token, cookieExpiresIn, err := service.LoginService(s.db, c)
	if err != nil {
		log.Printf("Error logging in: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.SetCookieHandler("Authorization", token, "/", cookieExpiresIn, c)
	c.JSON(http.StatusOK, "User logged in successfully")
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

	id, err := service.ChangeRoleService(s.db, c.Query("id"), c.Query("role"))
	if err != nil {
		log.Printf("Error changing user role: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User with id: "+(id)+" role changed to "+c.Query("role")+" successfully")
}
