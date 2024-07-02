package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Employee struct {
	Id          int
	Name        string
	Designation string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getAllEmployees", getAllEmployees).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func getAllEmployees(w http.ResponseWriter, req *http.Request) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=test password=password host=localhost port=5432 sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db == nil)
	sqlStmt := "select * from employee1"
	rows, err := db.Query(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}
	var employees []Employee
	for rows.Next() {
		var e Employee
		rows.Scan(&e.Id, &e.Name, &e.Designation)
		employees = append(employees, e)
	}
	// fmt.Println(employees)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(employees)
	fmt.Println("Endpoint Hit: getting all users")
}
