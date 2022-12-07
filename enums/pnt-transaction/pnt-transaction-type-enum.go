package pnt_transaction

type PntTransactionType string

const (
	TYPE_PAYMENT    PntTransactionType = "payment"
	TYPE_WITHDRAW   PntTransactionType = "withdrawal"
	TYPE_COMMISSION PntTransactionType = "commission"
)
