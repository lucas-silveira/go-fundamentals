package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	tx, _ := db.Begin() // inicia uma transação

	stmt, _ := tx.Prepare("insert into usuarios(id, name) values(?,?)")

	stmt.Exec(10, "Mary")
	stmt.Exec(11, "Carl")
	_, err = stmt.Exec(10, "Peter") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
