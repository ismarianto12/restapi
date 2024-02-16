package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var inventory []Item

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, inventory)
}

func AddItem(c *gin.Context) {
	var newItem Item
	// Bind the JSON body to the newItem struct
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		// panic(err.Error())
	}

	inventory = append(inventory, newItem)
	c.JSON(http.StatusCreated, gin.H{"message": "Item added successfully"})
}

func Update(c *gin.Context) {
	var Payload Item
	if err := c.Bind(&Payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error()})
	}
	//	var name = Payload.Name
	//	var price = Payload.Price
	var id = Payload.ID
	c.JSON(http.StatusOK, gin.H{
		"data payload": id,
		"messages":     "data berhasil di update",
	})
	// if err: =
}
func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	if err := c.Bind(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error()})
	}
	log.Println(file.Filename)
	c.SaveUploadedFile(file, "./uploads")
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
