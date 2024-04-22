package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.StaticFS("/static", gin.Dir("static", true))

    router.GET("/", showLoginPage)
    router.GET("/login", showLoginPage)          // Added for login page
    router.GET("/register", showRegistrationPage) // Added for redirection
    router.GET("/confirm", confirmationHandler)  // Added for confirmation page

    router.POST("/register", registerHandler)
    router.POST("/login", loginHandler) // Added for login form submission

    router.Run(":8080")
}
