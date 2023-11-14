package main

import (
	"fmt"
	"log"

	"gkps.dbcommand/command"
	"gkps.dbquery/query"
)

//var db *sql.DB

func main() {
	/* // Data source name properties
	dsn := mysql.Config{
		User:   "root",
		Passwd: "J@cks0n1974",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "sakila",
	} */

	var err error
	//get database handle
	//db, err = sql.Open("mysql", dsn.FormatDSN())
	_, err = command.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer command.CloseConnection() //it will be closed after the Main is completed
	//code to ping the database and check the connection
	msg, pingErr := command.PingConnection()
	if pingErr != nil {
		log.Fatal(pingErr)
		fmt.Println(msg)
	}
	fmt.Println(msg)

	//rows, err := command.UpdateActor("Ghana", 203)
	actors, err := query.GetActor(20332321)
	if err != nil {
		log.Fatal(err)
	}
	if actors != nil {
		fmt.Printf("Total actor rows effected: %v\n", actors)
	}
}
