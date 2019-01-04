package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	s.GET("/:msg", func(c *gin.Context) {
		log.Println(c.Param("msg"))
		c.String(200, c.Param("msg"))
	})
	s.Run(":3000")
}