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

	stmt, _ := db.Prepare("update usuarios set name = ? where id = ?")

	stmt.Exec("Alice", 1)
	stmt.Exec("Jack", 2)
}
