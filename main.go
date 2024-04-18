package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFS("/static", gin.Dir("static", true))

	router.GET("/", showRegistrationPage)

	router.POST("/register", registerHandler)

	router.Run(":8080")
}
