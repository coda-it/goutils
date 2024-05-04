package mailer

import (
	"github.com/coda-it/goutils/logger"
	"net/smtp"
)

// IMailer - interface for mailer
type IMailer interface {
	AddRecipient(string)
	SendEmail(string, string, string)
	BulkEmail(string, string)
}

// Mailer - email notifier
type Mailer struct {
	Sender      string
	Password    string
	SMTPPort    string
	SMTPAuthURL string
	recipients  []string
}

// New - creates new instance of Mailer
func New(recipients []string, sender string, password string, SMTPPort string, SMTPAuthURL string) *Mailer {
	return &Mailer{
		Sender:      sender,
		Password:    password,
		SMTPPort:    SMTPPort,
		SMTPAuthURL: SMTPAuthURL,
		recipients:  recipients,
	}
}

// AddRecipient - adds new recipient of mailer
func (m *Mailer) AddRecipient(email string) {
	m.recipients = append(m.recipients, email)
}

func composeMessage(from string, to string, subject string, body string) string {
	return "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body
}

// SendEmail - send email to subscriber
func (m *Mailer) SendEmail(subject string, body string, recipient string) {
	msg := composeMessage(m.Sender, recipient, subject, body)
	smtpAuth := smtp.PlainAuth("", m.Sender, m.Password, m.SMTPAuthURL)

	err := smtp.SendMail(m.SMTPAuthURL+":"+m.SMTPPort, smtpAuth, m.Sender, []string{recipient}, []byte(msg))

	if err != nil {
		logger.Log("error sending email to " + recipient + ": " + err.Error())
		return
	}

	logger.Log("email sent to " + recipient)
}

// BulkEmail - sends email to all users
func (m *Mailer) BulkEmail(subject string, body string) {
	for _, r := range m.recipients {
		m.SendEmail(subject, body, r)
	}
}
