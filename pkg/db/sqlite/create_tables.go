package sqlite

import "log"

const testTable = `CREATE TABLE if not exists notes(
  id INTEGER PRIMARY KEY AUTOINCREMENT, 
  name TEXT not null,
  start datetime not null,
  end datetime not null);`

func (l *SqlLite) CreateTables() error {
	_, err := l.conn.Exec(testTable)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
