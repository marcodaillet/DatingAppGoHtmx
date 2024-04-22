package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func showRegistrationPage(c *gin.Context) {
	c.File("./static/register.html")
}

func registerHandler(c *gin.Context) {
	// Validate input fields
	username := strings.TrimSpace(c.PostForm("username"))
	email := strings.TrimSpace(c.PostForm("email"))
	firstName := strings.TrimSpace(c.PostForm("firstName"))
	lastName := strings.TrimSpace(c.PostForm("lastName"))
	password := strings.TrimSpace(c.PostForm("password"))

	if err := validateInputs(username, email, firstName, lastName, password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initialize database connection
	db, err := InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Check if email is already registered
	if emailExists(db, email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	// Hash password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	confirmationToken := tokenGenerator()

	_, err = db.Exec("INSERT INTO users (username, email, firstName, lastName, password, confirmationToken) VALUES ($1, $2, $3, $4, $5, $6)", username, email, firstName, lastName, hashedPassword, confirmationToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Send registration confirmation email with token link
	recipient := email
	subject := "Confirm Your Registration"
	body := fmt.Sprintf("Please click the following link to confirm your Tinder registration: <strong>localhost:8080/confirm?token=%s</strong>", confirmationToken)
	SendEmail(recipient, subject, body)

	// Registration successful
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful. Please check your email for confirmation."})
}

func validateInputs(username, email, firstName, lastName, password string) error {
	if username == "" || email == "" || firstName == "" || lastName == "" || password == "" {
		return errors.New("all fields are required")
	}

	if !isValidEmail(email) {
		return errors.New("invalid email address")
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func emailExists(db *sql.DB, email string) bool {
	var existingEmail string
	err := db.QueryRow("SELECT email FROM users WHERE email = $1", email).Scan(&existingEmail)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		// Handle other errors if necessary
		return true
	}
	return true
}
