package email

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"path/filepath"

	"loan-service/configs"
)

func SendEmailWithAttachment(to, subject, body, filePath string) error {
	cfg := configs.LoadConfig()
	from := cfg.SMTP.Email
	password := cfg.SMTP.Password
	smtpHost := cfg.SMTP.Host
	smtpPort := cfg.SMTP.Port

	// Read the attachment
	attachmentBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read attachment: %v", err)
	}

	// Base64 encode the attachment
	encodedAttachment := base64.StdEncoding.EncodeToString(attachmentBytes)
	attachmentName := filepath.Base(filePath)

	// Create the email headers
	subjectHeader := fmt.Sprintf("Subject: %s\n", subject)
	toHeader := fmt.Sprintf("To: %s\n", to)
	fromHeader := fmt.Sprintf("From: %s\n", from)
	mime := "MIME-version: 1.0;\nContent-Type: multipart/mixed; boundary=BOUNDARY\n\n"

	// Create the email body
	message := bytes.Buffer{}
	message.WriteString(toHeader)
	message.WriteString(fromHeader)
	message.WriteString(subjectHeader)
	message.WriteString(mime)
	message.WriteString("--BOUNDARY\n")
	message.WriteString("Content-Type: text/plain; charset=\"utf-8\"\n\n")
	message.WriteString(body + "\n\n")

	// Add the attachment
	message.WriteString("--BOUNDARY\n")
	message.WriteString("Content-Type: application/pdf\n")
	message.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\n", attachmentName))
	message.WriteString("Content-Transfer-Encoding: base64\n\n")
	message.WriteString(encodedAttachment + "\n")
	message.WriteString("--BOUNDARY--")

	// Send the email
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message.Bytes())
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
