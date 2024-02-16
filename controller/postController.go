package controller

import (
	"fmt"
	"net/http"

	"github.com/apidingin/models" // Import your controller package
	// Import your controller package
	// Import your controller package
	"github.com/gin-gonic/gin"
	// Import your controller package
	"github.com/apidingin/config" // Import your controller package
)

func ListPost(c *gin.Context) {
	var user []models.User
	var dbmap = config.InitDb()
	_, err := dbmap.Select(&user, `select username,
	password,
	level from login`)
	fmt.Println(err)
	if err == nil {
		c.JSON(200, gin.H{
			"response": http.StatusOK, "messages": "success fully response", "data": user})
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}
