package mail

func GenerateMail(to, subject, body string) Mail {
	return Mail{
		To:      []string{to},
		Subject: subject,
		Body:    body,
	}
}

func GenerateMails(to []string, subject, body string) Mail {
	return Mail{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}
