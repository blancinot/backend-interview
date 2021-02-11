package account

import (
	"context"
)

type Transaction struct {
	ID        string
	Amount    float64
	AccountID string
	CreatedAt int64
}

type FilterTransaction struct {
	AccountID string
	UserID    string
}

type StoreTransaction interface {
	InsertTransaction(context.Context, Transaction) error
	FetchManyTransaction(context.Context, FilterTransaction, func(Transaction) error) error
	FetchMaxTransaction(context.Context, FilterTransaction, int64, int64) (Transaction, error)
}
