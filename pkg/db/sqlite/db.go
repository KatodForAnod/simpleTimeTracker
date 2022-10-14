package sqlite

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type SqlLite struct {
	conn *sql.DB
}

func (l *SqlLite) InitDataBase() error {
	path := "db.sqlite"

	if _, err := os.Stat(path); err == nil {
		// path/to/whatever exists
	} else if errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(path)
		if err != nil {
			log.Fatalln(err)
			return err
		}
	} else {
		panic("error!")
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}

	l.conn = db
	return nil
}
