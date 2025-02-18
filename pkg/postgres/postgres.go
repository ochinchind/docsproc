// Package postgres implements postgres connection.
package postgres

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ochinchind/docsproc/config"
	"github.com/ochinchind/docsproc/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

// Postgres -.
type Postgres struct {
	Conn         *gorm.DB
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

// New -.
func New(url string, opts ...Option) (*Postgres, error) {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	pg := &Postgres{
		Conn: db,
	}

	return pg, nil
}

func (p *Postgres) Connect(cfg *config.Config) error {
	conn, err := gorm.Open(postgres.Open(cfg.URL), &gorm.Config{})
	if err != nil {
		return err
	}

	p.Conn = conn
	return nil
}

func (p *Postgres) Migrate() error {
	err := p.Conn.AutoMigrate(
		&entity.User{},
	)

	if err != nil {
		return fmt.Errorf("Migrating entities to Postgres - err: %w", err)
	}
	return nil
}
