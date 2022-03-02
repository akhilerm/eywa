package upi

import (
	"fmt"
)

const (
	UPIScheme     = "upi://"
	UPIURLAddress = "pay"
	separator     = "&"
	INR           = "INR"
	//upiRegex      = "[a-zA-Z0-9-_]{4,20}@[a-zA-Z]{2,64}"
)

type PaymentRequest struct {
	payeeAddress    string
	payeeName       string
	transactionNote string
	amount          string
	// currently only INR
	currency string
}

// NewPaymentRequest creates a new payment request
// currently the currency is set at INR
func NewPaymentRequest() *PaymentRequest {
	return &PaymentRequest{
		currency: INR,
	}
}

func (p *PaymentRequest) WithPayeeAddress(pa string) *PaymentRequest {
	p.payeeAddress = pa
	return p
}

func (p *PaymentRequest) WithPayeeName(pn string) *PaymentRequest {
	p.payeeName = pn
	return p
}

func (p *PaymentRequest) WithAmount(am string) *PaymentRequest {
	p.amount = am
	return p
}

func (p *PaymentRequest) WithCurrency(cu string) *PaymentRequest {
	p.currency = cu
	return p
}

func (p *PaymentRequest) WithTransactionNote(tn string) *PaymentRequest {
	p.transactionNote = tn
	return p
}

// GenerateUPIPaymentURL generates a payment URL as specified in
// https://developers.google.com/pay/india/api/android/in-app-payments.
// More parameters and its usage can be found in the above doc.
func (p PaymentRequest) GenerateUPIPaymentURL() (string, error) {
	if !p.IsValid() {
		return "", fmt.Errorf("invalid payment request: %v", p)
	}
	return UPIScheme + UPIURLAddress + "?" +
		"pa=" + p.payeeAddress + separator +
		"pn=" + p.payeeName + separator +
		"am=" + p.amount + separator +
		"cu=" + p.currency + separator +
		"tn=" + p.transactionNote, nil
}

// IsValid returns if the payment request is valid
func (p PaymentRequest) IsValid() bool {
	//re := regexp.MustCompile(upiRegex)
	return !(len(p.payeeName) == 0 ||
		p.currency != INR ||
		len(p.transactionNote) == 0 ||
		len(p.amount) == 0 ||
		len(p.payeeAddress) == 0)
}
