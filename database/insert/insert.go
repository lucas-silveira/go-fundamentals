package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	statement, _ := db.Prepare("insert into usuarios(name) values(?)")
	statement.Exec("Mary")
	statement.Exec("John")

	res, _ := statement.Exec("Peter")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	lines, _ := res.RowsAffected()
	fmt.Println(lines)
}
