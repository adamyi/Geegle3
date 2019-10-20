package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Challenge struct {
	Sender          string
	Title           string
	Body            string
	DependsOnPoints int
	Delay           int64
}

type Flag struct {
	Flag   string
	Points int
}

type Configuration struct {
	ListenAddress string
	DbType        string
	DbAddress     string
	JwtKey        []byte
	Challenges    []Challenge
	Flags         []Flag
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
		fmt.Println(err.Error())
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}
	var inited int
	err = _db.QueryRow("select count(*) from scoreboard where user = ?", user).Scan(&inited)
	if err != nil {
		fmt.Println(err.Error())
		rsp.WriteHeader(http.StatusInternalServerError)
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
	addFlag(user, "GEEGLE{WELCOME_TO_GEEGLE}", false)
}

func addFlag(user string, body string, sendConfirmation bool) {
	// TODO: move this to a separate flag service
	var oPoints int
	err := _db.QueryRow("select points from scoreboard where user = ?", user).Scan(&oPoints)
	if err != nil {
		msg := "Sorry, something went wrong :("
		_db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", "noreply@geegle.org", user, "Error", msg, time.Now().UnixNano()/1000000)
		return
	}
	flags := ""
	points := 0
	for _, flag := range _configuration.Flags {
		if strings.Contains(body, flag.Flag) {
			var count int
			err = _db.QueryRow("select count(*) from submission where flag = ? and user = ?", flag.Flag, user).Scan(&count)
			if err != nil {
				msg := "Sorry, something went wrong :("
				_db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", "noreply@geegle.org", user, "Error", msg, time.Now().UnixNano()/1000000)
				return
			}
			if count == 0 {
				points += flag.Points
				flags += flag.Flag + ", "
				_db.Exec("insert into submission (flag, user, time) values(?, ?, ?)", flag.Flag, user, time.Now().UnixNano()/1000000)
			}
		}
	}

	if points > 0 {
		_db.Exec("update scoreboard set points = ? where user = ?", oPoints+points, user)
		if sendConfirmation {
			msg := fmt.Sprintf("You found %s you have earned %d points. You now have %d points.", flags, points, oPoints+points)
			_db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", "noreply@geegle.org", user, "Congrats", msg, time.Now().UnixNano()/1000000)
		}
		for _, challenge := range _configuration.Challenges {
			if challenge.DependsOnPoints <= (oPoints+points) && challenge.DependsOnPoints > oPoints {
				_db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", challenge.Sender, user, challenge.Title, challenge.Body, time.Now().UnixNano()/1000000+challenge.Delay)
			}
		}
	} else {
		msg := "Sorry, we did not recognise that flag :("
		_db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", "noreply@geegle.org", user, "Error", msg, time.Now().UnixNano()/1000000)
	}

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
		addFlag(user, string(e.Body), true)
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
	_db, err = sql.Open(_configuration.DbType, _configuration.DbAddress)
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
