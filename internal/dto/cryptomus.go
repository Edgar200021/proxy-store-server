package dto

import "github.com/shopspring/decimal"

type CurrencyCode string

const (
	USD CurrencyCode = "USD"
	EUR CurrencyCode = "EUR"
	RUB CurrencyCode = "RUB"
)

type PaymentStatus string

const (
	PAID                 PaymentStatus = "paid"
	PAID_OVER            PaymentStatus = "paid_over"
	WRONG_AMOUNT         PaymentStatus = "wrong_amount"
	PROCESS              PaymentStatus = "process"
	CONFIRM_CHECK        PaymentStatus = "confirm_check"
	WRONG_AMOUNT_WAITING PaymentStatus = "wrong_amount_waiting"
	CHECK                PaymentStatus = "check"
	FAIL                 PaymentStatus = "fail"
	CANCEL               PaymentStatus = "cancel"
	SYSTEM_FAIL          PaymentStatus = "system_fail"
	REFUND_PROCESS       PaymentStatus = "refund_process"
	REFUND_FAIL          PaymentStatus = "refund_fail"
	REFUND_PAID          PaymentStatus = "refund_paid"
	LOCKED               PaymentStatus = "locked"
)

type CreateCryptomusInvoiceRequest struct {
	Amount   decimal.Decimal `json:"amount"`
	Currency CurrencyCode    `json:"currency"`
	OrderId  string          `json:"order_id"`
}

type CreateCryptomusInvoiceResponse struct {
	State  int16 `json:"state"`
	Result struct {
		Url     string `json:"url"`
		OrderId string `json:"order_id"`
	}
}

type CryptomusPaymentInfoResponse struct {
	State  int16 `json:"state"`
	Result struct {
		Uuid    string          `json:"uuid"`
		OrderId string          `json:"order_id"`
		Amount  decimal.Decimal `json:"amount"`
		Status  PaymentStatus   `json:"status"`
	}
}
