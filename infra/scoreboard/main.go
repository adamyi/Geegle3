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

	geemail "geegle.org/infra/geemail-client"
	_ "github.com/go-sql-driver/mysql"
)

var _db *sql.DB

type Player struct {
	Name   string
	Points int
}

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

var _configuration = Configuration{}

func initScoreboardRsp(w http.ResponseWriter) {
	w.Header().Add("Server", "scoreboard")
}

func addFlag(user string, body string, sendConfirmation bool) {
	var oPoints int
	err := _db.QueryRow("select points from scoreboard where user = ?", user).Scan(&oPoints)
	if err != nil {
		fmt.Println(err)
		msg := []byte("Sorry, something went wrong :(")
		geemail.SendEmailNow("noreply@geegle.org", user, "Error", msg)
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
				msg := []byte("Sorry, something went wrong :(")
				geemail.SendEmailNow("noreply@geegle.org", user, "Error", msg)
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
			msg := []byte(fmt.Sprintf("You found %s you have earned %d points. You now have %d points.", flags, points, oPoints+points))
			geemail.SendEmailNow("noreply@geegle.org", user, "Congrats", msg)
		}
		fmt.Println(oPoints + points)
		for _, challenge := range _configuration.Challenges {
			if challenge.DependsOnPoints <= (oPoints+points) && challenge.DependsOnPoints > oPoints {
				geemail.SendEmailWithDelay(challenge.Sender, user, challenge.Title, []byte(challenge.Body), challenge.Delay)
			}
		}
	} else {
		fmt.Println(flags, points)
		msg := []byte("Sorry, we did not recognise that flag :(")
		geemail.SendEmailNow("noreply@geegle.org", user, "Error", msg)
	}

}

func listenAndServe(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		initScoreboardRsp(w)

		data := make([]Player, 0, 30)
		rows, err := _db.Query("select user, points from scoreboard ORDER BY points DESC LIMIT 30")
		if err != nil {
			http.Error(w, "I don't know what happened", http.StatusInternalServerError)
			return
		}

		for rows.Next() {
			player := Player{}
			rows.Scan(&player.Name, &player.Points)
			data = append(data, player)
		}

		RenderTemplate(w, "index.html", data)
	})

	mux.HandleFunc("/api/submit", func(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc("/api/init_user", func(w http.ResponseWriter, r *http.Request) {
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
			Username string `json:"username"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Malformed Data", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var inited int
		err = _db.QueryRow("select count(*) from scoreboard where user = ?", data.Username).Scan(&inited)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if inited > 0 {
			w.WriteHeader(200)
			return
		}

		_db.Exec("insert into scoreboard (user, points) values (?,?)", data.Username, 0)
		addFlag(data.Username, "GEEGLE{WELCOME_TO_GEEGLE}", false)
		w.WriteHeader(200)
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
