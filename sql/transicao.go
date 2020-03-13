package main

import (
	_ "Golang/mysql"
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("mysql","root:root@tcp(172.17.0.2:3306)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//defer deteccao do db.Close ()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(5, "Bianca")
	stmt.Exec(6, "Carlos")
	_, err = stmt.Exec(4, "Tiago")

	if err != nil {
		tx.Rollback()
		//encerra a transação desfazendo todas as alterações realizadas.
		log.Fatal(err)
	}
	tx.Commit()
}
