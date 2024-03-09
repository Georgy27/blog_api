package pg

import (
	"context"
	"github.com/pkg/errors"

	"github.com/Georgy27/blogger_api/pkg/db"
	"github.com/jackc/pgx/v4/pgxpool"
)

type pgClient struct {
	db db.DB
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		db: &pg{pool: dbc},
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.db
}

func (c *pgClient) Close() error {
	if c.db != nil {
		c.db.Close()
	}

	return nil
}
