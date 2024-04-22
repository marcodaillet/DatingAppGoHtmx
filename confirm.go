package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func confirmationHandler(c *gin.Context) {
	c.File("./static/confirm.html")
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Confirmation token is missing"})
		return
	}

	db, err := InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to connect to database"})
		return
	}
	defer db.Close()

	var confirmed bool
	err = db.QueryRow("SELECT confirmed FROM users WHERE confirmationtoken = $1", token).Scan(&confirmed)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid confirmation token"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to verify confirmation token"})
		return
	}

	_, err = db.Exec("UPDATE users SET confirmed = true WHERE confirmationtoken = $1", token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to confirm registration"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Registration confirmed successfully"})
}
