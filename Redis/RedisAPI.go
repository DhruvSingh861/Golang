package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type Employee struct {
	Name    string
	Address Address
}
type Address struct {
	Street string
	City   string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getAllEmployees", getAllEmployees).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func getAllEmployees(w http.ResponseWriter, req *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	var employees []Employee
	for i := 0; i < 2; i++ {
		name, _ := client.HGet(fmt.Sprint("employee", i), "name").Result()
		street, _ := client.HGet(fmt.Sprint("employee", i), "address.street").Result()
		city, _ := client.HGet(fmt.Sprint("employee", i), "address.city").Result()
		address := Address{
			Street: street,
			City:   city,
		}
		employee := Employee{
			Name:    name,
			Address: address,
		}
		// val, _ := client.HGet("tempMap", "name").Result()
		employees = append(employees, employee)
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(employees)
	fmt.Println("Endpoint Hit")
}
