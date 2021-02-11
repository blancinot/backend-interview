package sql

import (
	"context"
	"strings"

	"github.com/gustvision/backend-interview/pkg/account"
)

func (s *Store) InsertTransaction(ctx context.Context, t account.Transaction) error {
	b := strings.Builder{}

	b.WriteString(`INSERT INTO transaction ( `)
	b.WriteString(`id, amount, account_id, created_at `)
	b.WriteString(`) VALUES ( `)
	b.WriteString(`$1, $2, $3, $4 `)
	b.WriteString(`);`)

	if _, err := s.ExecContext(ctx, b.String(), []interface{}{
		t.ID,
		t.Amount,
		t.AccountID,
		t.CreatedAt,
	}...); err != nil {
		return err
	}

	return nil
}

func (s *Store) FetchManyTransaction(
	ctx context.Context,
	f account.FilterTransaction,
	callback func(account.Transaction) error,
) error {
	b := strings.Builder{}
	b.WriteString(`SELECT id, amount, account_id, created_at `)
	b.WriteString(`FROM transaction `)
	b.WriteString(`WHERE account_id = $1 ;`)

	rows, err := s.QueryContext(ctx, b.String(), []interface{}{
		f.AccountID,
	}...)
	if err != nil {
		return err
	}

	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var t account.Transaction

		if err := rows.Scan(
			&t.ID,
			&t.Amount,
			&t.AccountID,
			&t.CreatedAt,
		); err != nil {
			return err
		}

		if err := callback(t); err != nil {
			return err
		}
	}

	return rows.Err()
}

func (s *Store) FetchMaxTransaction(
	ctx context.Context,
	f account.FilterTransaction,
	from, to int64,
) (account.Transaction, error) {

	b := strings.Builder{}
	b.WriteString(`SELECT t.id, t.amount, t.account_id, t.created_at `)
	b.WriteString(`FROM transaction t `)
	b.WriteString(`INNER JOIN account a ON a.id = t.account_id `)
	b.WriteString(`WHERE a.user_id = $1 AND t.created_at >= $2 AND t.created_at <= $3 `)
	b.WriteString(`ORDER BY t.amount DESC `)
	b.WriteString(`LIMIT 1;`)

	row := s.QueryRowContext(ctx, b.String(), []interface{}{
		f.UserID,
		from,
		to,
	}...)

	var transaction account.Transaction
	if err := row.Scan(
		&transaction.ID,
		&transaction.Amount,
		&transaction.AccountID,
		&transaction.CreatedAt,
	); err != nil {
		return account.Transaction{}, err
	}

	return transaction, nil
}
