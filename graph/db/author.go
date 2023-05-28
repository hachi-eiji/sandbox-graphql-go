package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"sandbox-grahql-go/sqls"
)

func ListAuthors() ([]sqls.Author, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, DatabaseConn())

	if err != nil {
		return nil, err
	}

	defer func(conn *pgx.Conn, ctx context.Context) {
		_ = conn.Close(ctx)
	}(conn, ctx)

	db := sqls.New(conn)

	return db.ListAuthors(ctx)
}
