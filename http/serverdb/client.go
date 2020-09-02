package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// UserStruct ...
type UserStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserHandler ...
func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		getUserByID(w, r, id)
	case r.Method == "GET":
		getAllUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Cliente não encontrado.")
	}
}

func getUserByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var user UserStruct
	db.QueryRow("select id, name from usuarios where id = ?", id).Scan(&user.ID, &user.Name)
	// o método QueryRow retorna somente uma linha

	json, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, _ := db.Query("select id, name from usuarios limit 10")

	defer rows.Close()

	var users []UserStruct

	for rows.Next() {
		var user UserStruct
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
