package payment

var (
	StatusPending = "PENDING"
	StatusPaid    = "PAID"
	StatusFailed  = "FAILED"
)

type Payment struct {
	ID            string `json:"id"`
	CorrelationID string `json:"correlationId"`
	Status        string `json:"status"`
}
