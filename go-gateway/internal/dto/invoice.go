package dto

import "github.com/eulincon/imersao22/go-gateway/internal/domain"

const (
	StatusPending  = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

type CreateInvoiceInput struct {
	APIKey      string
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	PaymentType string  `json:"payment_type"`
	CardNumber  string  `json:"card_number"`
	CVV         string  `json:"cvv"`
	ExpiryMonth int     `json:"expiry_month"`
	ExpiryYear  int     `json:"expiry_year"`
	CardHolder  string  `json:"card_holder"`
}

type InvoiceOutput struct {
	ID             string
	AccountId      string  `json:"account_id"`
	Amount         float64 `json:"amount"`
	Status         string  `json:"status"`
	Description    string  `json:"description"`
	PaymentType    string  `json:"payment_type"`
	CardLastDigits string  `json:"card_last_digits"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

func ToInvoice(input CreateInvoiceInput, accountID string) (*domain.Invoice, error) {
	card := domain.CreditCard{
		Number:         input.CardNumber,
		CVV:            input.CVV,
		ExpiryMonth:    input.ExpiryMonth,
		ExpiryYear:     input.ExpiryYear,
		CardholderName: input.CardHolder,
	}
	return domain.NewInvoice(
		accountID,
		input.Description,
		input.PaymentType,
		input.Amount,
		card,
	)
}

func FromInvoice(invoice *domain.Invoice) *InvoiceOutput {
	return &InvoiceOutput{
		ID:             invoice.ID,
		AccountId:      invoice.AccountID,
		Amount:         invoice.Amount,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentType:    invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      invoice.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
