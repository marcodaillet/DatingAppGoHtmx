// email.go
package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// SendEmail sends an email using the specified SMTP server
func SendEmail(recipient, subject, body string) error {
	fmt.Println(recipient)
	// Read SMTP configuration from environment variables
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	sender := os.Getenv("SENDER_EMAIL")

	// Set up authentication information
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpServer)

	// Connect to the server, authenticate, and send the email
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, sender, []string{recipient}, []byte("Subject: "+subject+"\r\n"+"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"+body))
	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}
	return nil
}
