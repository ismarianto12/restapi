package controller

import "github.com/gin-gonic/gin"

type GetData struct {
	paramData string `form:field_a`
}

func GetSiswa(c *gin.Context) {
	var paramData GetData
	// c.Params = append(c.Params, paramData)
	c.Bind(&paramData)
	c.JSON(200, gin.H{
		"data": paramData,
	})
}

func DetailSiswa(c *gin.Context) {
	var paramData GetData
	c.Bind(&paramData)
	c.JSON(200, gin.H{
		"data": paramData,
	})
}

func StoreSIswa(c *gin.Context) {

}
