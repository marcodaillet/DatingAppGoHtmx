package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showRegistrationPage(c *gin.Context) {
	c.File("./static/register.html")
}

func registerHandler(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	password := c.PostForm("password")

	db, err := InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	var existingEmail string
	err = db.QueryRow("SELECT email FROM users WHERE email = $1", email).Scan(&existingEmail)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email uniqueness"})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	_, err = db.Exec("INSERT INTO users (username, email, firstName, lastName, password) VALUES ($1, $2, $3, $4, $5)", username, email, firstName, lastName, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful. Please check your email for confirmation."})
}
