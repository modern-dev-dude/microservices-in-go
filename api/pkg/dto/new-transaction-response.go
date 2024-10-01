package dto

type NewTransactionResponse struct {
	TransactionId  string  `json:"transaction_id"`
	CurrentBalance float64 `json:"current_balance"`
}
