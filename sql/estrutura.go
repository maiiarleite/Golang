package main

import (
	_ "Golang/mysql"
	"database/sql"
)

func exec(db *sql.DB, sql string) sql.Result {
		result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/")
	//mysql - pasta / root - senha
	if err != nil {
		panic(err)
	}
	defer  db.Close()

	exec(db, "Create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios(
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY(id)
	)`)
}