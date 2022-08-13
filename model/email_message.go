package model

type EmailMessage struct {
	From    string
	Subject string
	Body    string
	To      []string
}
