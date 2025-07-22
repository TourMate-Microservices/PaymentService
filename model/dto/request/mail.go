package request

import "log"

type MailBody struct {
	Email         string
	Subject       string
	Username      string
	TransactionId int
}

type SendMailRequest struct {
	Body         MailBody    `json:"mail_body"`
	TemplatePath string      `json:"template_path"`
	Logger       *log.Logger `json:"logger"`
}
