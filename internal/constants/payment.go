package constants

// payment method
const (
	PaymentMethodCash     = "CASH"
	PaymentMethodQRIS     = "QRIS"
	PaymentMethodTransfer = "TRANSFER"
	PaymentMethodDebit    = "DEBIT"
	PaymentMethodKredit   = "KREDIT"
)

// payment status
const (
	PaymentStatusPaid     = "PAID"
	PaymentStatusUnpaid   = "UNPAID"
	PaymentStatusPartial  = "PARTIAL"
	PaymentStatusVoid     = "VOID"
	PaymentStatusRefunded = "REFUNDED"
)
