package mail

func GenerateMail(to, subject, body string) Mail {
	return Mail{
		To:      []string{to},
		Subject: subject,
		Body:    body,
	}
}

func GenerateMails(to []string, subject, body string) []Mail {
	numMails := len(to)/batch + 1
	mails := make([]Mail, numMails)

	for i := 0; i < numMails; i++ {
		start := i * batch
		end := start + batch
		if end > len(to) {
			end = len(to)
		}
		mails[i] = Mail{
			To:      to[start:end],
			Subject: subject,
			Body:    body,
		}
	}
	return mails
}
