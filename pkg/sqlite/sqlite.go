package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type ISqliteInstance interface {
	Database() *sql.DB
	Close() error
}

type sqliteInstance struct {
	db *sql.DB
}

func NewSqliteInstance(dbName string) (ISqliteInstance, error) {
	instance := &sqliteInstance{}
	if err := instance.initDB(dbName); err != nil {
		log.Fatal("Failed to initialize database", err)
		return nil, err
	}
	return instance, nil
}

func (s *sqliteInstance) Database() *sql.DB {
	return s.db
}

func (s *sqliteInstance) Close() error {
	return s.db.Close()
}

func (s *sqliteInstance) initDB(dbName string) error {
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s.db", dbName))
	if err != nil {
		return err
	}

	query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        duration INTEGER,
        difficulty INTEGER
    );
    `
	if _, err := db.Exec(query); err != nil {
		return err
	}

	s.db = db
	return nil
}
