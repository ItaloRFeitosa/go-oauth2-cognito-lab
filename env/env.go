package env

import "os"

var (
	PaymentURL = os.Getenv("PAYMENT_URL")
	OrderURL   = os.Getenv("ORDER_URL")
)
