package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var _db *sql.DB

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
	JwtKey        []byte
	Challenges    []Challenge
	Flags         []Flag
}

type Email struct {
	ID       int    `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Subject  string `json:"subject"`
	Body     []byte `json:"body"`
	Time     int64  `json:"time"`
}

var _configuration = Configuration{}

func initScoreboardRsp(w http.ResponseWriter) {
	w.Header().Add("Server", "")
}

func sendEmail(sender string, receiver string, subject string, body string, time int64) {
	email := Email{0, sender, receiver, subject, []byte(body), time}
	reqBody, err := json.Marshal(email)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post("https://mail.corp.geegle.org/api/addmail", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("mail resp %+v", resp)

}

func addFlag(user string, body string, sendConfirmation bool) {
	var oPoints int
	err := _db.QueryRow("select points from scoreboard where user = ?", user).Scan(&oPoints)
	if err != nil {
                fmt.Println(err)
		msg := "Sorry, something went wrong :("
		sendEmail("noreply@geegle.org", user, "Error", msg, time.Now().UnixNano()/1000000)
		return
	}
	flags := ""
	points := 0
	for _, flag := range _configuration.Flags {
		if strings.Contains(body, flag.Flag) {
			var count int
			err = _db.QueryRow("select count(*) from submission where flag = ? and user = ?", flag.Flag, user).Scan(&count)
			if err != nil {
                                fmt.Println(err)
				msg := "Sorry, something went wrong :("
				sendEmail("noreply@geegle.org", user, "Error", msg, time.Now().UnixNano()/1000000)
				return
			}
			if count == 0 {
				points += flag.Points
				flags += flag.Flag + ", "
				_db.Exec("insert into submission (flag, user, time) values(?, ?, ?)", flag.Flag, user, time.Now().UnixNano()/1000000)
			}
		}
	}

        fmt.Println(flags, points)
	if points > 0 {
                fmt.Println("im here now")
		_db.Exec("update scoreboard set points = ? where user = ?", oPoints+points, user)
		if sendConfirmation {
			msg := fmt.Sprintf("You found %s you have earned %d points. You now have %d points.", flags, points, oPoints+points)
			sendEmail("noreply@geegle.org", user, "Congrats", msg, time.Now().UnixNano()/1000000)
		}
		for _, challenge := range _configuration.Challenges {
			if challenge.DependsOnPoints <= (oPoints+points) && challenge.DependsOnPoints > oPoints {
				sendEmail(challenge.Sender, user, challenge.Title, challenge.Body, time.Now().UnixNano()/1000000+challenge.Delay)
			}
		}
	} else {
		msg := "Sorry, we did not recognise that flag :("
		sendEmail("noreply@geegle.org", user, "Error", msg, time.Now().UnixNano()/1000000)
	}

}

func listenAndServe(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/view/", func(w http.ResponseWriter, r *http.Request) {
		initScoreboardRsp(w)
	})

	mux.HandleFunc("/submit/", func(w http.ResponseWriter, r *http.Request) {
		initScoreboardRsp(w)

		tknStr := r.Header.Get("X-Geegle-JWT")
		err := confirmFromGeemail(tknStr, _configuration.JwtKey)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JWT"})
			fmt.Println(err)
			return
		}

		data := struct {
			Username         string `json:"username"`
			Body             string `json:"flag"`
			SendConfirmation bool   `json:"confirm"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Malformed Data", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		addFlag(data.Username, data.Body, data.SendConfirmation)
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
	_db, err = sql.Open("sqlite3", os.Args[2])
	if err != nil {
		panic(err)
	}
	defer _db.Close()

	log.Panic(listenAndServe(_configuration.ListenAddress))
}
