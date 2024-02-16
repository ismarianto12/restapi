package controller

import (
	"github.com/apidingin/models" // Import your controller package
	// Import your controller package
	"github.com/gin-gonic/gin"
	// Import your controller package
)

func ListPost(c *gin.Context) {
	var user []models.User
	_, err := dbmap.Select(&user, "select * from user")
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}
