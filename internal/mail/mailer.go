package mail

import (
	"GithubReleaseNotificationAPI/internal/github"
	"fmt"
	"net/smtp"
)

type Service interface {
	SendConfirmationEmail(toEmail, repoName, confirmToken string) error
	SendReleaseNotification(toEmail string, repoName string, release *github.Release) error
}

type smtpService struct {
	host       string
	port       string
	user       string
	pass       string
	fromEmail  string
	appBaseURL string
}

func NewSMTPService(host, port, user, pass, fromEmail, appBaseURL string) Service {
	return &smtpService{
		host:       host,
		port:       port,
		user:       user,
		pass:       pass,
		fromEmail:  fromEmail,
		appBaseURL: appBaseURL,
	}
}

func (s *smtpService) SendConfirmationEmail(toEmail, repoName, confirmToken string) error {
	subject := fmt.Sprintf("Confirm subscription: %s", repoName)
	confirmLink := fmt.Sprintf("%s/api/confirm/%s", s.appBaseURL, confirmToken)

	body := fmt.Sprintf(`
		<p>Confirm subscription to <b>%s</b>:</p>
		<p>
			<a href="%s" style="background-color: #2ea44f; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; display: inline-block;">Confirm Subscription</a>
		</p>
	`, repoName, confirmLink)

	return s.send(toEmail, subject, body)
}

func (s *smtpService) SendReleaseNotification(toEmail string, repoName string, release *github.Release) error {
	subject := fmt.Sprintf("New Release for %s: %s", repoName, release.Tag)

	body := fmt.Sprintf(`
		<h3>New release available for <b>%s</b></h3>
		<p><b>Tag:</b> %s</p>
		<p><b>Name:</b> %s</p>
		<p><a href="%s" style="background-color: #0366d6; color: white; padding: 8px 16px; text-decoration: none; border-radius: 5px; display: inline-block;">View Release on GitHub</a></p>
	`, repoName, release.Tag, release.Name, release.URL)

	return s.send(toEmail, subject, body)
}

func (s *smtpService) send(toEmail, subject, body string) error {
	var auth smtp.Auth
	if s.user != "" && s.pass != "" {
		auth = smtp.PlainAuth("", s.user, s.pass, s.host)
	}

	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
		"\r\n"+
		"%s", s.fromEmail, toEmail, subject, body)

	address := fmt.Sprintf("%s:%s", s.host, s.port)
	err := smtp.SendMail(address, auth, s.fromEmail, []string{toEmail}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email to %s: %w", toEmail, err)
	}

	return nil
}
