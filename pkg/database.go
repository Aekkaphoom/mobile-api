package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/microsoft/go-mssqldb"
)

func SQLConnection(host, user, pass, dbname string, port int) (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		host, user, pass, port, dbname)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
