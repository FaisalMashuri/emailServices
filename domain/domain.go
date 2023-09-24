package domain

type EmailDomain interface {
	SendEmail(emailReciever, otp string) error
}
