package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("static/*.html")

	router.Static("/static", "./static/")

	authMiddleware := func(c *gin.Context) {
		sessionToken, err := c.Cookie("session_token")
		if err != nil || sessionToken == "" {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		c.Next()
	}

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

	router.GET("/home", authMiddleware, func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"Title":        "Home Page",
			"ContentBlock": "home_content",
		})
	})

	router.GET("/confirm", confirmationHandler) // Added for confirmation page
	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler) // Added for login form submission
	router.POST("/disconnect", disconnectHandler)

	router.Run(":8080")
}
