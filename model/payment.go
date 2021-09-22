package model

type SavePaymentResponse struct {
	Error string   `json:"error"`
	Data  *Payment `json:"data"`
}

type GetAllPaymentsResponse struct {
	Error string    `json:"error"`
	Data  []Payment `json:"data"`
}

type Payment struct {
	Amount      float64    `json:"amount"`
	Category    string     `json:"category"`
	Date        CustomTime `json:"date"`
	Description string     `json:"description"`
	PaymentType string     `json:"payment_type"`
}
