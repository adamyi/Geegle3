package main

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mvdan/xurls"
)

type Configuration struct {
	ListenAddress string
	SmtpAddress   string
	DbType        string
	DbAddress     string
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

type BotMsg struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var _configuration = Configuration{}
var _db *sql.DB
var emailRe = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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
	_, err = http.Post("https://scoreboard.corp.geegle.org/api/submit", "application/json", bytes.NewBuffer(reqBody))

	return err
}

func triggerXSS(player, service, link string) (BotMsg, error) {
	var result BotMsg
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://xssbot.corp.geegle.org/?url="+url.QueryEscape(link), nil)
	req.Header.Set("X-Geegle-SubAcc", service)
	req.Header.Set("X-UberProxy-Player", player)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

func findURLsInString(str string) []string {
	httpre, _ := xurls.StrictMatchingScheme("http")
	httpsre, _ := xurls.StrictMatchingScheme("https")
	urls := append(httpre.FindAllString(str, -1), httpsre.FindAllString(str, -1)...)
	used := make(map[string]struct{})
	var empty struct{}
	for _, v := range urls {
		used[v] = empty
	}
	res := make([]string, 0)
	for v, _ := range used {
		res = append(res, v)
	}
	return res
}

func batchTriggerXSS(player, service, addr, title, body string) {
	retmsg := "Hi " + player + "@,<br><br>"
	urls := findURLsInString(body)
	if len(urls) == 0 {
		retmsg += "Sorry, I couldn't find any links in your email"
	} else if len(urls) > 1 {
		retmsg += "There are too many links in your email (" + strings.Join(urls, ", ") + ")! I don't know which one to visit. Please provide one only."
	} else {
		bmsg, err := triggerXSS(player, service, urls[0])
		if err != nil {
			retmsg += "I found " + urls[0] + " in your email but something went wrong and I couldn't visit it (" + err.Error() + ")"
		} else if bmsg.Success {
			retmsg += "I will take a look at " + urls[0] + " (" + bmsg.Message + ")"
		} else {
			retmsg += "I found " + urls[0] + " in your email but something went wrong and I couldn't visit it (" + bmsg.Message + ")"
		}
	}
	_, err := _db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", addr, player+"@geegle.org", "Re: "+title, []byte(retmsg), time.Now().UnixNano()/1000000)
	if err != nil {
		fmt.Println(err)
		return
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

	// TODO Maybe dont init user every single time
	initUser(user)
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
	submitData := struct {
		Username string `json:"username"`
	}{
		user,
	}

	reqBody, err := json.Marshal(submitData)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = http.Post("https://scoreboard.corp.geegle.org/api/init_user", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		fmt.Println(err)
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

	if !emailRe.MatchString(e.Receiver) {
		rsp.WriteHeader(http.StatusBadRequest)
		rsp.Write([]byte("Invalid receiver address"))
		return
	}
	if len(e.Subject) > 78 {
		rsp.WriteHeader(http.StatusBadRequest)
		rsp.Write([]byte("Email subject too long"))
		return
	}
	if strings.Contains(e.Subject, "\n") || strings.Contains(e.Subject, "\r") {
		rsp.WriteHeader(http.StatusBadRequest)
		rsp.Write([]byte("No new line in subject"))
		return
	}

	_, err = _db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", user, e.Receiver, e.Subject, e.Body, time.Now().UnixNano()/1000000)
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if e.Receiver == "flag@geegle.org" {
		fmt.Println(addFlag(user, string(e.Body), true))
	}
	if e.Receiver == "adamyi@geegle.org" {
		batchTriggerXSS(strings.Split(user, "@")[0], "seclearn", e.Receiver, e.Subject, string(e.Body))
	}
	if e.Receiver == "pasteweb-feedback@geegle.org" {
		batchTriggerXSS(strings.Split(user, "@")[0], "pasteweb", e.Receiver, e.Subject, string(e.Body))
	}

	if !strings.HasSuffix(e.Receiver, "@geegle.org") {
		err = sendOutboundMail(e.Receiver, user, e.Subject, e.Body)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func sendOutboundMail(to, from, subject string, body []byte) error {
	c, err := smtp.Dial(_configuration.SmtpAddress)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Mail(from); err != nil {
		return err
	}
	if err = c.Rcpt(to); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	msg := "To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" + base64.StdEncoding.EncodeToString(body)

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
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
		fmt.Println(err)
		return
	}
	// TODO: whitelist services to call this function

	decoder := json.NewDecoder(req.Body)
	var e Email
	err = decoder.Decode(&e)
	if err != nil {
		fmt.Println(err)
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = _db.Exec("insert into email (sender, receiver, subject, body, time) values(?, ?, ?, ?, ?)", e.Sender, e.Receiver, e.Subject, e.Body, e.Time)
	if err != nil {
		fmt.Println(err)
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
