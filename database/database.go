package database

import (
	"database/sql"
	"fmt"
)

func InitDB() *sql.DB {
	connectionString := "root:root@tcp(localhost:8080)/northwind"

	databaseConnection, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error()) //manejo de errerores
		fmt.Println("errorrrrrrrrrrrrrrrrr")
	}
	return databaseConnection
}
