package database

import (
	"database/sql"
	"fmt"
	"github.com/caneroj1/hush"
	_ "github.com/lib/pq" // anonymous import
	"log"
)

// app package database variable
var Database *sql.DB

// InitDB initializes the database connection
func InitDB() {
	secrets := hush.Hushfile()
	log.Println("Initiating connection to database.")

	user, ok := secrets.GetString("username")
	if !ok {
		panic("Could not get the username from .hushfile")
	}

	dbname, ok := secrets.GetString("dbname")
	if !ok {
		panic("Could not get the dbname from .hushfile")
	}

	password, ok := secrets.GetString("password")
	if !ok {
		panic("Could not get the password from .hushfile")
	}

	connString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		user,
		dbname,
		password)

	var err error
	Database, err = sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	log.Println("Pinging db...")

	err = Database.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("No errors. Connected to database successfully.")
}

func getCardsTableName() string {
	secrets := hush.Hushfile()
	table, ok := secrets.GetString("table")
	if !ok {
		log.Fatal("No table name exists")
		return ""
	}

	return table
}

// QueryRow returns a single row as a result from executing some sql
func QueryRow(sqlString string, params ...interface{}) *sql.Row {
	table := getCardsTableName()

	sql := fmt.Sprintf(sqlString, table)

	return Database.QueryRow(sql, params...)
}

// GetByQuery returns the rows associated with a query that is used to return data
func GetByQuery(query string, params ...interface{}) *sql.Rows {
	table := getCardsTableName()
	var (
		err  error
		rows *sql.Rows
	)

	sql := fmt.Sprintf(query, table)
	rows, err = Database.Query(sql, params...)

	if err != nil {
		log.Fatal(err)
	}

	return rows
}
