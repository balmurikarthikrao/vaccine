package main

import (
	"database/sql"
	"log"
	"vaccine/router"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "password"
	dbname   = "vaccine"
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

	// // Define the API endpoints
	// router.HandleFunc("/beneficiaries", createBeneficiary).Methods("POST")
	// router.HandleFunc("/appointments", createAppointment).Methods("POST")
	// router.HandleFunc("/appointments/{id}", getAppointment).Methods("GET")
	// router.HandleFunc("/appointments/{id}", cancelAppointment).Methods("DELETE")

	// // Start the server
	// log.Println("Server started on port 8000")
	// log.Fatal(http.ListenAndServe(":8000", router))
}
