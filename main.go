package main

import (
	"fmt"
	"ramongo/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDB()
	//logica
	defer databaseConnection.Close()

	fmt.Println(databaseConnection)

	// r := chi.NewRouter()
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// http.ListenAndServe(":3000", r)
}
