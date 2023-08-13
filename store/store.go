package store

import (
	"database/sql"
	"time"

	"github.com/aksbuzz/mood-journal/internal/cache"
)

type Store struct {
	db        *sql.DB
	userCache *cache.Cache
}

func New(db *sql.DB) *Store {
	c := cache.New(time.Hour*24, "user")
	return &Store{
		db:        db,
		userCache: c,
	}
}

func (s *Store) GetDB() *sql.DB {
	return s.db
}
