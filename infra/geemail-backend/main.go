package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Configuration struct {
	ListenAddress string
	JwtKey        []byte
}

type UserInfo struct {
	Username string  `json:"username"`
	Inbox    []Email `json:"inbox"`
	Sent     []Email `json:"sent"`
}
type Email struct {
	ID       int    `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Subject  string `json:"subject"`
	Body     []byte `json:"body"`
	Time     int    `json:"time"`
}

var _configuration = Configuration{}
var _db *sql.DB

func readConfig() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&_configuration)
	if err != nil {
		panic(err)
	}
}

func initGmRsp(rsp http.ResponseWriter) {
	rsp.Header().Add("Server", "geemail")
	rsp.Header().Add("Access-Control-Allow-Origin", "https://mail.corp.geegle.org")
	rsp.Header().Add("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	rsp.Header().Add("Access-Control-Allow-Credentials", "true")
	rsp.Header().Add("Access-Control-Allow-Headers", "Content-Type")
}

// for user to get its own info and emails
func userInfo(rsp http.ResponseWriter, req *http.Request) {
	initGmRsp(rsp)
	if req.Method == "OPTIONS" {
		return
	}

	tknStr := req.Header.Get("X-Geegle-JWT")
	user, err := getJwtUserName(tknStr, _configuration.JwtKey)
	if err != nil {
		fmt.Printf(err.Error())
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}
	// user := "adamyi@geegle.org"
	info := &UserInfo{
		Username: user,
		Inbox:    []Email{},
		Sent:     []Email{},
	}
	rows, err := _db.Query("select id, sender, receiver, subject, body, time from email where sender=? or receiver=?", user, user)
	if err != nil {
		fmt.Println(err.Error())
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var email = &Email{}
		err = rows.Scan(&email.ID, &email.Sender, &email.Receiver, &email.Subject, &email.Body, &email.Time)
		if err != nil {
			fmt.Println(err.Error())
			rsp.WriteHeader(http.StatusInternalServerError)
			return
		}
		if email.Sender == user {
			info.Sent = append(info.Sent, *email)
		}
		if email.Receiver == user {
			info.Inbox = append(info.Inbox, *email)
		}
	}
	rsp.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(rsp).Encode(info)
}

// for user to send email
func sendMail(rsp http.ResponseWriter, req *http.Request) {
	initGmRsp(rsp)
	if req.Method == "OPTIONS" {
		return
	}

	tknStr := req.Header.Get("X-Geegle-JWT")
	user, err := getJwtUserName(tknStr, _configuration.JwtKey)
	if err != nil {
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}
	decoder := json.NewDecoder(req.Body)
	var e Email
	err = decoder.Decode(&e)
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = _db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", user, e.Receiver, e.Subject, e.Body, time.Now().UnixNano())
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// to be called by trusted apps, e.g. smtpd
func addMail(rsp http.ResponseWriter, req *http.Request) {
	initGmRsp(rsp)
	if req.Method == "OPTIONS" {
		return
	}

	tknStr := req.Header.Get("X-Geegle-JWT")
	_, err := getJwtUserName(tknStr, _configuration.JwtKey)
	if err != nil {
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}
	// TODO: whitelist services to call this function

	decoder := json.NewDecoder(req.Body)
	var e Email
	err = decoder.Decode(&e)
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = _db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", e.Sender, e.Receiver, e.Subject, e.Body, time.Now().UnixNano())
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	readConfig()
	var err error
	_db, err = sql.Open("sqlite3", "geemail.db")
	if err != nil {
		log.Fatal(err)
	}
	defer _db.Close()
	http.HandleFunc("/api/userinfo", userInfo)
	http.HandleFunc("/api/sendmail", sendMail)
	http.HandleFunc("/api/addmail", addMail)
	err = http.ListenAndServe(_configuration.ListenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
