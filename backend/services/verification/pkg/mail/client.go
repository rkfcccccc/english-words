package mail

import (
	"fmt"
	"net/smtp"
)

type Client struct {
	host, port, username, displayName string
	auth                              smtp.Auth
}

func NewClient(host, port, username, password, displayName string) *Client {
	return &Client{host, port, username, displayName, smtp.PlainAuth("", username, password, host)}
}

func (m *Client) Send(recipient, subject, content string) error {
	addr := m.host + ":" + m.port

	body := fmt.Sprintf("Content-Type: text/html; charset=\"UTF-8\";\r\n"+
		"From: \"%s\" <%s>\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n%s\r\n",
		m.displayName, m.username, recipient, subject, content,
	)

	err := smtp.SendMail(addr, m.auth, m.username, []string{recipient}, []byte(body))

	if err != nil {
		return fmt.Errorf("smtp.SendMail: %v", err)
	}

	return nil
}
