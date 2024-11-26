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
