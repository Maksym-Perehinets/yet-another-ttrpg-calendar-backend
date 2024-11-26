package server

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/cmd/api/service"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/shared/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// UpdateUserHandler updates user info if user who is updating is the same as the user being updated
// else will log suspicious activity and return unauthorized
func (s *Server) UpdateUserHandler(c *gin.Context) {
	user, ok := c.Get("user_id")
	if !ok {
		log.Printf("Error getting user from context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user id from context"})
		return
	}
	log.Printf("Updating user for %s by %s", c.Param("id"), user)
	updateUser := c.Param("id")

	if updateUser != user {
		log.Printf("Error updating user: User %s tried to update %s user data", user, updateUser)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You can't update someone else's info"})
		return
	}

	var updateFields request.UserUpdate

	if err := c.ShouldBindJSON(&updateFields); err != nil {
		log.Printf("Error updating user: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Updating user: %s with fields: %v", updateUser, updateFields)

	err := service.UpdateUserService(s.db, updateUser, updateFields)
	if err != nil {
		log.Printf("Error updating user: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User with id: "+(updateUser)+" updated successfully")
}
