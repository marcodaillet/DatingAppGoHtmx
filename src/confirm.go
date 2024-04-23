package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func confirmationHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.Redirect(http.StatusFound, "/login?confirmed=false&error=Confirmation token is missing")
		return
	}

	db, err := InitDB()
	if err != nil {
		c.Redirect(http.StatusFound, "/login?confirmed=false&error=Failed to connect to database")
		return
	}
	defer db.Close()

	var confirmed bool
	err = db.QueryRow("SELECT confirmed FROM users WHERE confirmationtoken = $1", token).Scan(&confirmed)
	if err == sql.ErrNoRows {
		c.Redirect(http.StatusFound, "/login?confirmed=false&error=Invalid confirmation token")

		return
	} else if err != nil {
		c.Redirect(http.StatusFound, "/login?confirmed=false&error=Failed to verify confirmation token")
		return
	}

	_, err = db.Exec("UPDATE users SET confirmed = true WHERE confirmationtoken = $1", token)
	if err != nil {
		c.Redirect(http.StatusFound, "/login?confirmed=false&error=Failed to confirm registration")
		return
	}

	c.Redirect(http.StatusFound, "/login?confirmed=true")
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Registration confirmed successfully"})
}
