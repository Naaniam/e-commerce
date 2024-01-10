package service

import (
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

var store = sessions.NewCookieStore([]byte("secret"))

// EmailConfig contains SMTP server details.
type EmailConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

func SendMailHandler(c echo.Context) error {
	// Retrieve user email from the session
	session, err := store.Get(c.Request(), "member-session")
	if err != nil {
	return c.JSON(http.StatusInternalServerError, "Error Occured: "+err.Error())
	}

	userEmail, ok := session.Values["email"].(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "Error Occured")
	}

	session, err = store.Get(c.Request(), "response-session")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error Occured: "+err.Error())
	}
	response, ok := session.Values["response"]
	if !ok {
		return c.JSON(http.StatusInternalServerError, "Error Occured")
	}

	SendConfirmationEmail(userEmail, response)
	// Display the user's email in the mail sending handler
	return c.JSON(http.StatusOK, "Message: Check order detials through the Mail sent to you to the Email: "+userEmail)
}

func SendConfirmationEmail(toEmail string, response interface{}) error {
	// Replace these values with your actual SMTP server details
	emailConfig := EmailConfig{
		Username: "sivaharitha.s@mitrahsoft.com",
		Password: "fkhsznntenaexgjq",
		Port:     "587",
		Host:     "smtp.gmail.com",
	}

	auth := smtp.PlainAuth("", emailConfig.Username, emailConfig.Password, emailConfig.Host)

	// Email template
	subject := "Confirmation Email"
	body := fmt.Sprintf("Thanks for ordering\n%s", response)

	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", toEmail, subject, body)

	// Connect to the SMTP server and send the email
	err := smtp.SendMail(emailConfig.Host+":"+emailConfig.Port, auth, emailConfig.Username, []string{toEmail}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
