package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
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
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&_configuration)
	if err != nil {
		panic(err)
	}
}

func addFlag(username string, body string, confirmation bool) error {
	submitData := struct {
		Username         string `json:"username"`
		Body             string `json:"flag"`
		SendConfirmation bool   `json:"confirm"`
	}{
		username, body, confirmation,
	}

	reqBody, err := json.Marshal(submitData)
	if err != nil {
		return err
	}

	resp, err := http.Post("https://scoreboard.corp.geegle.org/submit", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		return err
	}

	fmt.Printf("%+v", resp)
	return nil
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
	var inited int
	err = _db.QueryRow("select count(*) from scoreboard where user = ?", user).Scan(&inited)
	if err != nil {
		fmt.Printf(err.Error())
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}
	if inited == 0 {
		initUser(user)
	}
	// user := "adamyi@geegle.org"
	info := &UserInfo{
		Username: user,
		Inbox:    []Email{},
		Sent:     []Email{},
	}
	rows, err := _db.Query("select id, sender, receiver, subject, body, time from email where (sender=? or receiver=?) and time < ? order by time desc", user, user, time.Now().UnixNano()/1000000)
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

func initUser(user string) {
	_db.Exec("insert into scoreboard (user, points) values (?,?)", user, 0)
	fmt.Println(addFlag(user, "GEEGLE{WELCOME_TO_GEEGLE}", false))
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
	_, err = _db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", user, e.Receiver, e.Subject, e.Body, time.Now().UnixNano()/1000000)
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: make better
	if e.Receiver == "flag@geegle.org" {
		fmt.Println(addFlag(user, string(e.Body), true))
	}
	// TODO: integrate with headless chrome
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
	_, err = _db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", e.Sender, e.Receiver, e.Subject, e.Body, time.Now().UnixNano()/1000000)
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	readConfig()
	var err error
	_db, err = sql.Open("sqlite3", os.Args[2])
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
