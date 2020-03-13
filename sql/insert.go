package main

import (
	_ "Golang/mysql"
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("Marcos")
	stmt.Exec("Paulo")

	res, _ := stmt.Exec("Livia")

	id, _ := res.LastInsertId()
	log.Println(id)

	linhas, _ := res.RowsAffected()
	log.Println(linhas)
}
