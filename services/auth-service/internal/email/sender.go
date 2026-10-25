package email

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"
)

type Sender interface {
	SendPasswordReset(
		ctx context.Context,
		to string,
		resetLink string,
	) error
}

type SMTPSender struct {
	host     string
	port     string
	username string
	password string
	from     string
}

func NewSMTPSender(
	host,
	port,
	username,
	password,
	from string,
) *SMTPSender {
	return &SMTPSender{
		host:     host,
		port:     port,
		username: username,
		from:     from,
	}
}

func (s *SMTPSender) SendPasswordReset(
	ctx context.Context,
	to string,
	resetLink string,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	addr := s.host + ":" + s.port

	auth := smtp.PlainAuth(
		"",
		s.username,
		s.password,
		s.host,
	)

	body := makeBody(
		"Hello,\r\n\r\n"+
		"Use the following link to reset your password:\r\n\r\n"+
		"%s\r\n\r\n"+
		"This link expires in 15 minutes. "+
		"If you did not request this, ignore this email.\r\n",
		resetLink,
	)

	message := strings.Join([]string{
		"From: " + s.from,
		"To: " + to,
		"Subject: Reset your password",
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		"",
		body,
	}, "\r\n")

	return smtp.SendMail(
		addr,
		auth,
		s.from,
		[]string{to},
		[]byte(message),
	)
}

func makeBody(
	text string,
	resetLink string,
) string {
	return fmt.Sprintf(text, resetLink)
}