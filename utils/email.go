package utils

import (
	"fmt"
	"net/smtp" //used to send emails using SMTP (Simple Mail Transfer Protocol).

	"github.com/akhilbabu26/multi-brand_backend_2/config"
)

func SendOTPEmail(toEmail, otp string)error{
	cfg := config.AppConfig.SMTP

	auth := smtp.PlainAuth(
		"",
		cfg.Email,
		cfg.Password,
		cfg.Host,
	)

	subject := "Subject: email verification OTP\r\n"
	body := fmt.Sprintf("Your OTP is: %s", otp)

	msg := []byte(subject + "\r\n" + body)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	return smtp.SendMail(
		addr,
		auth,
		cfg.Email,
		[]string{toEmail},
		msg,
	)
}