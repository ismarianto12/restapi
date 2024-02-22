package controller

import (
	"fmt"
	"net/http"

	"github.com/apidingin/config" // Import your controller package
	"github.com/apidingin/models" // Import your controller package
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Level    string `json:"password"`
}

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

func Tambahdata(c *gin.Context) {
	var dbmap = config.InitDb()
	var login Login

	if err := c.BindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := dbmap.Exec(`insert into login set username = ?, password = ?, level = ?`, login.Username, login.Password, login.Level)
	if err == nil {
		c.JSON(200, gin.H{"response": "data insert success fully"})
	} else {
		c.JSON(400, gin.H{"response": err.Error()})
	}
}

func Tesloping(c *gin.Context) {
	var numbers []string
	for i := 1; i <= 10000000000; i++ {
		data := fmt.Sprintf("data ke %d", i)
		numbers = append(numbers, data)
	} //ternyata golang tak sanggup handle ke 1 jta
	c.JSON(200, gin.H{"response": numbers})
}
