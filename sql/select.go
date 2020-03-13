package main

import (
	_ "Golang/mysql"
	"database/sql"
	"log"
)

type usuario struct {
	id int
	nome string
}

func main() {
	db, err := sql.Open("mysql","root:root@tcp(172.17.0.2:3306)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer  db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id")
	defer rows.Close()

	stmt2, _ := db.Prepare("delete from usuarios where id = ? ")
	stmt2.Exec(2)
	stmt2.Exec(4)

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		log.Println(u)
	}
}