// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package dbgen

import (
	"context"
	"database/sql"
)

type Querier interface {
	GetBook(ctx context.Context, bookids []int32) ([]Book, error)
	GetBookByTitle(ctx context.Context, title sql.NullString) ([]Book, error)
	GetUser(ctx context.Context) ([]User, error)
}

var _ Querier = (*Queries)(nil)
