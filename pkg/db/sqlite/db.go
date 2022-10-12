package sqlite

import (
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed db.sqlite
var sqlLiteFile []byte
