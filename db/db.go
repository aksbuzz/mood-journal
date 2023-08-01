package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"os"
)

//go:embed migration
var migrationFS embed.FS

//go:embed seed
var seedFS embed.FS

type DB struct {
	DBInstance *sql.DB
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Open(ctx context.Context) (err error) {
	fileName := "moods.db"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	sqliteDB, err := sql.Open("sqlite", "moods.db")
	if err != nil {
		return err
	}
	db.DBInstance = sqliteDB
	// Apply Migrations
	if err := db.applyMigrations(ctx); err != nil {
		return err
	}
	// Seed the database
	if err := db.seedDB(ctx); err != nil {
		return err
	}
	return nil
}

func (db *DB) Close() (err error) {
	return db.DBInstance.Close()
}

func (db *DB) applyMigrations(ctx context.Context) (err error) {
	fileName := "Schema.sql"
	schemaPath := fmt.Sprintf("migration/%s", fileName)
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

	fileName := "Mood.sql"
	seedPath := fmt.Sprintf("seed/%s", fileName)
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
