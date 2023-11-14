package query

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type Actor struct {
	actor_id   int64
	first_name string
	last_name  string
}

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
func GetActor(actorid int64) ([]Actor, error) {

	db, err := OpenConnection()
	if err != nil {
		return nil, err
	}
	var actors []Actor
	result, err := db.Query("Select actor_id, first_name, last_name from actor where actor_id =?",
		actorid)
	if err != nil {
		return nil, err //fmt.Errorf("GetActor %v: %v", actorid, err)
	}
	defer result.Close()
	defer CloseConnection()
	for result.Next() {
		var acts Actor
		if err := result.Scan(&acts.actor_id, &acts.first_name, &acts.last_name); err != nil {
			return nil, err //fmt.Errorf("GetActor %v: %v", actorid, err)
		}
		actors = append(actors, acts)

		if err := result.Err(); err != nil {
			return nil, err //fmt.Errorf("GetActor %v: %v", actorid, err)
		}
	}
	return actors, nil
}
