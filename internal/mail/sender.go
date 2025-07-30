package mail

import "log"

// Заглушка – продавец сам подключит SendGrid / SES
func SendReset(to, token string) {
	resetURL := "http://localhost:8080/reset?token=" + token
	log.Printf("[MAIL] Send reset link to %s: %s", to, resetURL)
}