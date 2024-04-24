package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("static/*.html")

	router.Static("/static", "./static/")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Title":        "Login Page",
			"ContentBlock": "login_content",
		})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Title":        "Login Page",
			"ContentBlock": "login_content",
		})
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"Title":        "Registration Page",
			"ContentBlock": "register_content",
		})
	})

	router.GET("/confirm", confirmationHandler) // Added for confirmation page
	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler) // Added for login form submission

	router.Run(":8080")
}
