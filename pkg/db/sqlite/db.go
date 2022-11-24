package sqlite

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path"
)

type SqlLite struct {
	conn *sql.DB
}

func (l *SqlLite) InitDataBase() error {
	dbPath, err := getDBPath()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if _, err := os.Stat(dbPath); err == nil {
		// path/to/whatever exists
	} else if errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(dbPath)
		if err != nil {
			log.Fatalln(err)
			return err
		}
	} else {
		log.Fatalln("error")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalln(err)
	}

	l.conn = db
	return nil
}

func (l *SqlLite) ShutDown() error {
	if err := l.conn.Close(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func getDBPath() (string, error) {
	dbBaseDir := "simpleTimeTracker"
	dbName := "db.sqlite"

	osDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	mkdirPath := path.Join(osDir, dbBaseDir)
	err = os.MkdirAll(mkdirPath, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	absolutePath := path.Join(mkdirPath, dbName)
	return absolutePath, nil
}
