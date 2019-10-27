package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Configuration struct {
	ListenAddress string
}

type Employee struct {
	ID       int
	Name     string
	Position string
}

var configuration = Configuration{}
var db *sql.DB

func findEmployeeByName(name string) ([]Employee, error) {
	rows, err := db.Query("SELECT * FROM employees WHERE name LIKE '%" + name + "%' ORDER BY id LIMIT 5")

	if err != nil {
		return nil, err
	}

	employees := make([]Employee, 0)

	for rows.Next() {
		var employee Employee
		fmt.Println(rows.Columns())
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Position)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func readConfig() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
}

func initSlRsp(rsp http.ResponseWriter) {
	rsp.Header().Add("Server", "employees")
}

func index(rsp http.ResponseWriter, req *http.Request) {
	initSlRsp(rsp)
	if req.Method == "OPTIONS" {
		return
	}

	name := req.FormValue("name")

	if name == "" {
		RenderTemplate(rsp, "index.html", []Employee{})
		return
	}

	employees, err := findEmployeeByName(name)
	if err != nil {
		http.Error(rsp, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	RenderTemplate(rsp, "index.html", employees)
}

func main() {
	readConfig()
	var err error
	db, err = sql.Open("sqlite3", os.Args[2])
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", index)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(configuration.ListenAddress, nil))
}
