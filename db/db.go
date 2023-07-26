package db

import (
	"database/sql"
	"os"
)

type DB struct {
	DBInstance *sql.DB
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Open() (err error) {
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
	if err := db.initialize(); err != nil {
		return err
	}
	return nil
}

func (db *DB) Close() (err error) {
	return db.DBInstance.Close()
}

func (db *DB) initialize() (err error) {
	_, err = db.DBInstance.Exec(`
		CREATE TABLE IF NOT EXISTS moods (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			mood VARCHAR(64) NOT NULL,
    	description TEXT NOT NULL,
    	date DATE NOT NULL
		)
	`)
	if err != nil {
		return err
	}
	return nil
}
