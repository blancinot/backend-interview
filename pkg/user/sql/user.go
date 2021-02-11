package sql

import (
	"context"
	"strings"

	"github.com/gustvision/backend-interview/pkg/user"
)

func (s *Store) Fetch(ctx context.Context, f user.Filter) (user.User, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, name `)
	b.WriteString(`FROM users `)
	b.WriteString(`WHERE id = $1 ;`)

	row := s.QueryRowContext(ctx, b.String(), []interface{}{
		f.ID,
	}...)

	var u user.User

	if err := row.Scan(
		&u.ID,
		&u.Name,
	); err != nil {
		return user.User{}, err
	}

	return u, nil
}
