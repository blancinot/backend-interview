package dto

import "github.com/gustvision/backend-interview/pkg/account"

type CreateTransactionReq struct {
	AccountID string
	Amount    float64
}

type CreateTransactionResp struct {
	account.Transaction

	AccountTotal float64
}

type GetMaxTransactionReq struct {
	UserID string
	From   int64
	To     int64
}
