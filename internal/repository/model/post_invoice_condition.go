package model

type PostInvoiceCondition struct {
	SuppliersId    uint64
	IssueDate      string
	PaymentAmount  float64
	Fee            float64
	FeeRate        float64
	Tax            float64
	TaxRate        float64
	TotalAmount    float64
	PaymentDueDate string
	Status         string
}
