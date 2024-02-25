package controller

import (
	"fmt"

	"github.com/apidingin/config" // Import your controller package
	"github.com/gin-gonic/gin"
)

type Kelas struct {
	kelas_id   int    `json:name="kelas_id"`
	nama_kelas string `json:name="nama_kelas"`
	kelas      string `json:name="kelas"`
	kode       string `json:name="kode"`
	mapel      string `json:name="mapel"`
}

// func KelasData(c *gin.Context) {
// 	var param Kelas
// 	var dbmap = config.InitDb
// 	data := dbmap.QueryRow("select * from kelas where kelas_id = ?", param.kelas_id).Scan(&param)
// 	if data != nil {
// 		return c.JSON(200, gin.H{
// 			"messages": "Data Kelas Kosong",
// 		})
// 	} else {
// 		return c.JSON(200, gin.H{
// 			"messages": param,
// 		})
// 	}
// }

func KelasData(c *gin.Context) {
	var param Kelas
	var column []string

	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	dbmap := config.InitDb()
	err := dbmap.SelectOne(&param, "select * from kelas where id = ?", param.kelas_id)
	column := err.Column()

	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": column})
		return
	}
	c.JSON(200, gin.H{
		"messages": "Request Successful",
		"data":     param,
	})
}
func updateKelas(c *gin.Context) {
	var status int
	var param Kelas
	dbmap := config.InitDb()
	dbmap.Update(`update kelas set kelas = ?, kode = ?, mapel= ?`,
		param.kelas, param.kode, param.mapel,
	)
	// if data.Is(data, data.ErrorNoRows) {
	// return c.JSON(200, gin.H{
	// 	"messages": "data berhasil diupdate",
	// })

	// }
	return c.JSON(200, gin.H{
		"messages": "data berhasil diupdate",
	})
	// if data.Error != nil {

	// }else{

	// }
}
