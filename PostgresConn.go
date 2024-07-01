package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Employee struct {
	// Each field has a struct tag in the form of db:"column_name".
	// These tags are used by the sqlx library (and other similar libraries) to
	// map the struct fields to the corresponding columns in a database table.
	Id          int    "db:id"
	Name        string `db:"name"`
	Designation string `db:"designation"`
}

func main() {
	fmt.Println("postgress Conn.....")
	insert()
	// getAllEmployee()
}

var wg sync.WaitGroup
var mut sync.Mutex

func insert() {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=test password=password host=localhost port=5432 sslmode=disable")
	fmt.Println(db.Ping())
	fmt.Println(err)
	employees := []Employee{}
	for i := 0; i < 10000000; i++ {
		employees = append(employees, Employee{Id: i, Name: fmt.Sprint("DHRUV", i), Designation: "SE"})
	}
	start := time.Now()
	counter := 0
	chunkSize := 10000
	for i := 0; i < len(employees); i += chunkSize {
		end := i + chunkSize

		if end > len(employees) {
			end = len(employees)
		}

		chunk := employees[i:end]
		wg.Add(1)
		counter++

		insertChunks := func(chunk []Employee, wg *sync.WaitGroup) {
			values := make([]map[string]interface{}, len(chunk))
			for j, emp := range chunk {
				values[j] = map[string]interface{}{
					"id":          emp.Id,
					"name":        emp.Name,
					"designation": emp.Designation,
				}
			}
			// mut.Lock()
			// fmt.Println(values[0])
			sqlStatement := `INSERT INTO employee1 (id,name,designation) VALUES (:id, :name, :designation)`

			_, err1 := db.NamedExec(sqlStatement, values)
			if err1 != nil {
				fmt.Println(err1)
			}
			// mut.Unlock()
			defer wg.Done()
		}
		go insertChunks(chunk, &wg)
	}
	wg.Wait()

	fmt.Println(time.Since(start).Seconds())
	time.Sleep(10 * time.Millisecond)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Inserted Successfully %d", counter)
	fmt.Println(runtime.NumCPU())
	// fmt.Println(runtime.GOMAXPROCS())
}

func getAllEmployee() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=test password=password host=localhost port=5432 sslmode=disable")
	fmt.Println(db.Ping())
	fmt.Println(err)

	employee := Employee{}
	rows, _ := db.Queryx("SELECT * FROM employee")

	for rows.Next() {
		err := rows.StructScan(&employee)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%v\n", employee)
	}
}
