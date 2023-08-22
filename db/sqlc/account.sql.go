package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (
	owner,
	balance,
	currency
) VALUES (
	$1, $2, $3
) RETURNING id, owner, balance, currency, create_at`

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var account Account
	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)

	return account, err
}
