package main

import (
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/emersion/go-smtp"
)

// SMTPHandlers implements SMTP server methods.
type SMTPHandlers struct{}

// Login handles a login command with username and password.
func (handlers *SMTPHandlers) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	return &Session{}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (handlers *SMTPHandlers) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return &Session{}, nil
}

// A Session is returned after successful login.
type Session struct {
	from string
	to   string
	data string
}

// Mail mail handler
func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	s.from = from
	return nil
}

// Rcpt rcpt handler
func (s *Session) Rcpt(to string) error {

	var exist bool

	for _, server := range conf.Server {
		if strings.HasSuffix(to, "@"+server.EmailDomain) {
			exist = true
		}
	}

	if !exist {
		return smtp.ErrAuthRequired
	}

	s.to = to
	return nil
}

// Data data handler
func (s *Session) Data(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	s.data = string(b)
	err = sendEmail(*s)
	log.Println(err)
	return nil
}

// Reset reset
func (s *Session) Reset() {}

// Logout logout
func (s *Session) Logout() error {
	return nil
}
