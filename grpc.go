package main

import (
	context "context"
	"errors"
	"strings"
)

func sendEmail(s Session) error {
	com := strings.Split(s.to, "@")
	conn, ok := receivers[com[1]]
	if !ok {
		return errors.New("Email not exist")
	}
	c := NewSMTPClient(conn)
	email := Email{
		From: s.from,
		To:   s.to,
		Data: s.data,
	}
	resp, err := c.SendEmail(context.Background(), &email)
	if err != nil {
		return err
	}
	if !resp.Status {
		return errors.New("Send Email failed")
	}
	return nil
}
