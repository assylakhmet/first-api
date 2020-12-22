package configuration

import "database/sql"

func SqlConfig() (*sql.DB, error) {

	connStr := "user=postgres password=mypass dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	return db, err
}
