package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gateway struct {
	Comm      string `form:"comm" json:"comm" uri:"comm" xml:"comm" binding:"required"`
	ModuleKey string `form:"moduleKey" json:"moduleKey" uri:"moduleKey" xml:"moduleKey" binding:"required"`
}

func HanleGateway(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(json)
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

func main() {
	route := gin.Default()
	route.POST("cgi-bin/gateway.fcg", HanleGateway)
	route.Run(":8000")
}
