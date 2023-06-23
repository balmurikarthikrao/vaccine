package main

import (
	"database/sql"
	"log"
	"vaccine/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Connect to the MySQL database
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/vaccine")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the router
	router := router.NewRouter(db)

	// Start listening and serving requests
	router.Run(":8080")

}
