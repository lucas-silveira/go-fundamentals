package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // o prefixo _ é usado para importar libs
	// cujas funções não serão utilizados no código
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err) // interrompe o programa e imprime o erro
	}

	return result
}

func main() {
	db, err := sql.Open("mysql", "root:mysql@/")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment PRIMARY KEY,
		name varchar(80)
	)`)
}
