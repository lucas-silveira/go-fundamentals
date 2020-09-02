package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeNow(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05") // essa data é usada somente para formatar a saída
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/", timeNow)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
