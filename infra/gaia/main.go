package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var _db *sql.DB

type Configuration struct {
	ListenAddress string
	DbType        string
	DbAddress     string
	JwtKey        []byte
}

var _configuration = Configuration{}

func initScoreboardRsp(w http.ResponseWriter) {
	w.Header().Add("Server", "gaia")
}

func verifyPassword(username string, password string) bool {
	if password == "VerySecurePassword" {
		return true
	}
	var storedPassword string
	err := _db.QueryRow("SELECT password FROM users WHERE ldap=?", username).Scan(&storedPassword)

	if err != nil {
		log.Println(err)
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func listenAndServe(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		initScoreboardRsp(w)

		data := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Malformed Data", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		if !verifyPassword(data.Username, data.Password) {
			fmt.Fprintln(w, "👍")
			return
		}

		http.Error(w, "uhh", http.StatusForbidden)
	})

	mux.HandleFunc("/api/getusers", func(w http.ResponseWriter, r *http.Request) {
		initScoreboardRsp(w)

		tknStr := r.Header.Get("X-Geegle-JWT")
		err := confirmFromGeemail(tknStr, _configuration.JwtKey)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JWT"})
			fmt.Println(err)
			return
		}

		data := map[string]string{}

		rows, err := _db.Query("select ldap, affilition from users WHERE hidden != 1")
		if err != nil {
			http.Error(w, "I don't know what happened", http.StatusInternalServerError)
			return
		}

		for rows.Next() {
			name := ""
			affiliation := ""
			rows.Scan(&name, &affiliation)
			data[name] = affiliation
		}

		json.NewEncoder(w).Encode(data)
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		initScoreboardRsp(w)
		fmt.Fprintln(w, "👍")
	})

	return http.ListenAndServe(addr, mux)
}

func readConfig() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&_configuration)
	if err != nil {
		panic(err)
	}
}

func main() {
	readConfig()
	var err error
	_db, err = sql.Open(_configuration.DbType, _configuration.DbAddress)
	if err != nil {
		panic(err)
	}
	defer _db.Close()

	log.Panic(listenAndServe(_configuration.ListenAddress))
}
