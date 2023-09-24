package models

type EventEmitter struct {
	Message string
	Payload interface{}
}

type EmailEmitter struct {
	Email string
	Otp   string
}
