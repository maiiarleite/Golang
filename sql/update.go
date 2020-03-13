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

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Francisco", 2)
	stmt.Exec("Joana", 5)
	stmt.Exec("Marcos", 1)

	stmt2, _ := db.Prepare("delete from usuarios where id = ? ")
	stmt2.Exec(3)
}
