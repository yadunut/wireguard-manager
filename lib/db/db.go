package db

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"github.com/yadunut/wireguard-manager/lib/clients"
	"golang.org/x/xerrors"
)

var (
	log = logrus.WithField("db", "sqlite3")
)

type DB struct {
	*sql.DB
}

func InitDB(filepath string) (*DB, error) {
	log.Info("Opening DB")
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, xerrors.New("db is nil")
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	table := `CREATE TABLE IF NOT EXISTS users (
ID INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
NAME TEXT NOT NULL,
PRIVATEKEY TEXT NOT NULL,
IP TEXT NOT NULL
);`

	if _, err := db.Exec(table); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) AddClient(client clients.Client) error {
	return nil
}
