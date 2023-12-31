package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/aksbuzz/mood-journal/config"
	"github.com/gofiber/fiber/v2/log"
)

//go:embed migration
var migrationFS embed.FS

//go:embed seed
var seedFS embed.FS

var (
	SchemaFileName = "Schema.sql"
	SeedFileName   = "Seed.sql"
)

type DB struct {
	DBInstance *sql.DB
	cfg        *config.Config
}

func New(cfg *config.Config) *DB {
	return &DB{
		cfg: cfg,
	}
}

func (db *DB) Open(ctx context.Context) (err error) {
	dbFilename := db.cfg.Database.Filename
	if _, err := os.Stat(dbFilename); os.IsNotExist(err) {
		file, err := os.Create(dbFilename)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	log.Info("Opening DB")
	sqliteDB, err := sql.Open("sqlite", dbFilename)
	if err != nil {
		return err
	}
	db.DBInstance = sqliteDB
	log.Info("Applying latest migrations")
	if err := db.applyMigrations(ctx); err != nil {
		return err
	}
	if err := db.seedDB(ctx); err != nil {
		return err
	}
	return nil
}

func (db *DB) Close() (err error) {
	return db.DBInstance.Close()
}

func (db *DB) applyMigrations(ctx context.Context) (err error) {
	schemaPath := fmt.Sprintf("migration/%s", SchemaFileName)
	buf, err := migrationFS.ReadFile(schemaPath)
	if err != nil {
		return err
	}
	tx, err := db.DBInstance.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, string(buf)); err != nil {
		return err
	}
	return tx.Commit()
}

func (db *DB) seedDB(ctx context.Context) (err error) {
	seeded, err := db.isSeeded(ctx)
	if err != nil {
		return err
	}
	if seeded {
		return nil
	}

	seedPath := fmt.Sprintf("seed/%s", SeedFileName)
	buf, err := seedFS.ReadFile(seedPath)
	if err != nil {
		return err
	}

	tx, err := db.DBInstance.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, string(buf)); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "CREATE TABLE seeded_flag (id INTEGER PRIMARY KEY);"); err != nil {
		return err
	}
	return tx.Commit()
}

func (db *DB) isSeeded(ctx context.Context) (bool, error) {
	row := db.DBInstance.QueryRowContext(ctx, "SELECT count(*) FROM sqlite_master WHERE type='table' AND name='seeded_flag';")
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
