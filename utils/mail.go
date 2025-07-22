package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"text/template"
	"tourmate/payment-service/constant/env"
	"tourmate/payment-service/constant/noti"
	"tourmate/payment-service/model/dto/request"

	"gopkg.in/mail.v2"
)

func SendMail(req request.SendMailRequest) error {
	var errLogMsg string = fmt.Sprintf(noti.MAIL_ERR_MSG, "Utils.Mail - SendMail")

	template, err := template.ParseFiles(req.TemplatePath)
	if err != nil {
		req.Logger.Println(errLogMsg + err.Error())
		return errors.New(noti.INTERNALL_ERR_MSG)
	}

	var body bytes.Buffer
	if err := template.Execute(&body, req.Body); err != nil {
		req.Logger.Println(errLogMsg + err.Error())
		return errors.New(noti.INTERNALL_ERR_MSG)
	}

	serviceEmail := os.Getenv(env.SERVICE_EMAIL)
	securityPass := os.Getenv(env.SECURITY_PASS)
	host := os.Getenv(env.HOST)
	port, err := strconv.Atoi(os.Getenv(env.MAIL_PORT))
	if err != nil {
		port = 587
	}

	m := mail.NewMessage()
	m.SetHeader("From", "tourmate_prn232@gmail.com")
	m.SetHeader("To", req.Body.Email)
	m.SetHeader("Subject", req.Body.Subject)
	m.SetBody("text/html", body.String())

	dialer := mail.NewDialer(host, port, serviceEmail, securityPass)

	if err := dialer.DialAndSend(m); err != nil {
		req.Logger.Println(errLogMsg + err.Error())
		return errors.New(noti.GENERATE_MAIL_WARN_MSG)
	}

	return nil
}
