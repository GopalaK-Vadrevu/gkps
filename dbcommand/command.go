package command

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func OpenConnection() (*sql.DB, error) {
	var dsn = mysql.Config{
		User:                 "devuser",   //root
		Passwd:               "admin1234", //"J@cks0n1974",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "sakila",
		AllowNativePasswords: true, // Need to added
	}
	db, err = sql.Open("mysql", dsn.FormatDSN())
	return db, err
}
func CloseConnection() {
	db.Close()
}
func PingConnection() (string, error) {
	msg := "Ping Error. Not Connected"
	pingErr := db.Ping()
	if pingErr != nil {
		return msg, pingErr
	}
	msg = "Connected and Live !!"
	return msg, nil
}

func AddActor(firstname string, lastname string) (int64, error) {
	db, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	result, err := db.Exec("Insert into actor (first_name, last_name) values (?,?)", firstname, lastname)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func UpdateActor(lastname string, actorid int64) (int64, error) {
	db, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	result, err := db.Exec("Update actor SET last_name= ? WHERE actor_id= ? ", lastname, actorid)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func DeleteActor(actorid int64) (int64, error) {
	db, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	result, err := db.Exec("Delete from actor WHERE actor_id= ? ", actorid)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return id, nil
}
