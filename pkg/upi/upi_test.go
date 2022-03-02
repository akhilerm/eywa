package upi

import "testing"

func TestPaymentRequest_IsValid(t *testing.T) {
	type fields struct {
		payeeAddress    string
		payeeName       string
		transactionNote string
		amount          string
		currency        string
	}
	tests := map[string]struct {
		fields fields
		want   bool
	}{
		"invalid upi id": {
			fields: fields{
				payeeAddress:    "a123-1@oksbi1",
				payeeName:       "Akhil Mohan",
				transactionNote: "Test transaction",
				amount:          "350",
				currency:        "INR",
			},
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p := PaymentRequest{
				payeeAddress:    tt.fields.payeeAddress,
				payeeName:       tt.fields.payeeName,
				transactionNote: tt.fields.transactionNote,
				amount:          tt.fields.amount,
				currency:        tt.fields.currency,
			}
			if got := p.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentRequest_GenerateUPIPaymentURL(t *testing.T) {
	type fields struct {
		payeeAddress    string
		payeeName       string
		transactionNote string
		amount          string
		currency        string
	}
	tests := map[string]struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p := PaymentRequest{
				payeeAddress:    tt.fields.payeeAddress,
				payeeName:       tt.fields.payeeName,
				transactionNote: tt.fields.transactionNote,
				amount:          tt.fields.amount,
				currency:        tt.fields.currency,
			}
			got, err := p.GenerateUPIPaymentURL()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateUPIPaymentURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateUPIPaymentURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
