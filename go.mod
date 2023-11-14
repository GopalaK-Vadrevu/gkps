module gkps

go 1.21.1

replace gkps.dbcommand/command => ./dbcommand

require (
	gkps.dbcommand/command v0.0.0-00010101000000-000000000000
	gkps.dbquery/query v0.0.0-00010101000000-000000000000
)

require github.com/go-sql-driver/mysql v1.7.1 // indirect

replace gkps.dbquery/query => ./dbquery
