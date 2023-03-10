package dao

import (
	"context"
	"database/sql"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) List(ctx context.Context) ([]Chat, error) {
	return nil, nil
}

func (p *Postgres) Get(ctx context.Context) (Chat, error) {
	return Chat{}, nil
}

func (p *Postgres) Create(ctx context.Context) (Chat, error) {
	return Chat{}, nil
}

func (p *Postgres) Update(ctx context.Context) (Chat, error) {
	return Chat{}, nil
}

func (p *Postgres) Delete(ctx context.Context) (Chat, error) {
	return Chat{}, nil
}
