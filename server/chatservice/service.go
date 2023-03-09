package chatservice

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ServiceWeaver/weaver"

	_ "github.com/lib/pq"
)

type Chat struct {
	weaver.AutoMarshal

	Name string `json:"name"`
}

type Service interface {
	CreateChat(context.Context, string) (Chat, error)
	GetChat(context.Context, string) (Chat, error)
	ListChat(context.Context) ([]Chat, error)
	DeleteChat(context.Context, string) error
}

type config struct {
	DBURI string `toml:"db_uri"`
}

type chatService struct {
	weaver.Implements[Service]
	weaver.WithConfig[config]

	db *sql.DB
}

func (s *chatService) Init(ctx context.Context) error {
	cfg := s.Config()

	var err error
	s.db, err = sql.Open("postgres", cfg.DBURI)
	if err != nil {
		return fmt.Errorf("error opening database %s: %w", cfg.DBURI, err)
	}

	return nil
}

func (s *chatService) CreateChat(context.Context, string) (Chat, error) {
	panic("implement me")
}
func (s *chatService) GetChat(context.Context, string) (Chat, error) {
	panic("implement me")
}
func (s *chatService) ListChat(context.Context) ([]Chat, error) {
	panic("implement me")
}
func (s *chatService) DeleteChat(context.Context, string) error {
	panic("implement me")
}
