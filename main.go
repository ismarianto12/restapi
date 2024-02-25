package main

import (
	// "fmt"
	// "log"
	// "net/http"
	// "github.com/apidingin/controller" // Import your controller package
	"github.com/apidingin/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/parameter", controller.GetItems)
	r.PUT("/update/item", controller.Update)
	r.POST("/upload", controller.UploadFile)
	r.POST("/addItem", controller.AddItem)
	r.POST("/user/store", controller.Tambahdata)
	r.GET("/user/list", controller.ListPost)
	r.GET("/api/looping", controller.Tesloping)
	r.GET("/api/getsiswa", controller.GetSiswa)
	r.GET("/api/array", controller.ParsingArray)
	r.GET("/api/kelas", controller.KelasData)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
