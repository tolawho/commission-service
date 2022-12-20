package pnt_transaction

type PntTransactionStatus string

const (
	// Các trạng thái chính của transaction
	STATUS_SUCCESSFUL PntTransactionStatus = "successful"
	STATUS_FAILED     PntTransactionStatus = "failed"
	STATUS_PROCESSING PntTransactionStatus = "processing"
	STATUS_CANCELED   PntTransactionStatus = "cancelled"
	// Trạng thái phụ của luồng rút tiền
	STATUS_OPS_PENDING        PntTransactionStatus = "ops_pending"
	STATUS_ACCOUNTANT_PENDING PntTransactionStatus = "accountant_pending"
	STATUS_TEMPORARY          PntTransactionStatus = "temporary" // hoa hồng tạm tính
)
