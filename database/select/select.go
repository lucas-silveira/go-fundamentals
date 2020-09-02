package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type userStruct struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, _ := db.Query("select id, name from usuarios where id > ?", 2)
	defer rows.Close()

	for rows.Next() {
		var user userStruct
		rows.Scan(&user.id, &user.name) // mapeamento implícito de colunas

		fmt.Println(user)
	}
}
