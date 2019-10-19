package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

var _configuration = Configuration{}

type submitData struct {
	Username         string `json:"username"`
	Body             string `json:"flag"`
	SendConfirmation bool   `json:"confirm"`
}

func initScoreboardRsp(w http.ResponseWriter) {
	w.Header().Add("Server", "")
}

// TODO: Use geemail service
func addFlag(user string, body string, sendConfirmation bool) {
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

func listenAndServe(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/view/", func(w http.ResponseWriter, r *http.Request) {
		initScoreboardRsp(w)
	})

	mux.HandleFunc("/submit/", func(w http.ResponseWriter, r *http.Request) {
		initScoreboardRsp(w)

		tknStr := r.Header.Get("X-Geegle-JWT")
		err := confirmFromGeemail(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JWT"})
			return
		}

		var data submitData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Malformed Data", http.StatusBadRequest)
			return
		}

		addFlag(submitData.Username, submitData.Body, submitData.SendConfirmation)
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
	addr := os.Args[2]
	_db, err := sql.Open("sqlite3", os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	defer _db.Close()

	log.Panic(listenAndServe(addr))
}
