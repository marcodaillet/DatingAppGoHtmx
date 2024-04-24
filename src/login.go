package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// loginHandler handles login form submissions
func loginHandler(c *gin.Context) {
	// Parse username and password from request body
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Initialize database connection
	db, err := InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Query the database to verify the username and password
	var hashedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE username = $1", loginData.Username).Scan(&hashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Verify password
	if !comparePasswords(hashedPassword, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate a session token (dummy implementation)
	sessionToken := tokenGenerator()

	// Set the session token as a cookie in the response
	c.SetCookie("session_token", sessionToken, 0, "/", "", false, true)

	// Successful login
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "session_token": sessionToken})
}
