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
	Id          int    "db:id"
	Name        string `db:"name"`
	Designation string `db:"designation"`
}

func (e Employee) String() string {
	return fmt.Sprintf("%v, %v, %v", e.Id, e.Name, e.Designation)
}

func main() {
	fmt.Println("postgress Conn.....")

	insert()
	// getAllEmployee()

	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())

	// buf := make([]byte, 1<<16)//Allocate a buffer for the stack trace
	// runtime.Stack(buf, true)  //Get stack traces of all goroutines
	// fmt.Printf("%s", buf)
}

var wg sync.WaitGroup
var mut sync.Mutex

func insert() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=test password=password host=localhost port=5432 sslmode=disable")
	defer db.Close()
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
		wg.Add(1)
		counter++
		go insertChunks(chunk, &wg)
	}
	wg.Wait()
	fmt.Println(time.Since(start).Seconds())
	fmt.Println("Inserted Successfully %d", counter)
	// fmt.Println(runtime.GOMAXPROCS())
}

func getAllEmployee() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=test password=password host=localhost port=5432 sslmode=disable")
	defer db.Close()
	fmt.Println(db.Ping())
	fmt.Println(err)

	rows, _ := db.Queryx("SELECT * FROM employee1")

	for rows.Next() {
		employee := Employee{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Designation)
		if err != nil {
			fmt.Println(err)
			// log.Fatalln(err)
		}
		log.Println(employee)
	}
}
