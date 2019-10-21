package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"
	"time"

	"io"
	"io/ioutil"

	"github.com/emersion/go-smtp"
)

var emailRe = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Backend struct{}

func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	return &Session{}, nil
}

func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return &Session{}, nil
}

type Session struct {
	From string
	To   []string
}

func (s *Session) Mail(from string) error {
	if !emailRe.MatchString(from) {
		return errors.New("Invalid email address")
	}
	if strings.Contains(from, "geegle.org") {
		return errors.New("If you are a geegler, please use https://mail.corp.geegle.org instead")
	}
	log.Println("Mail from:", from)
	s.From = from
	return nil
}

func (s *Session) Rcpt(to string) error {
	if !emailRe.MatchString(to) {
		return errors.New("Invalid email address")
	}
	if !strings.HasSuffix(to, "@geegle.org") {
		log.Println("Rejected Rcpt to:", to)
		return errors.New("Only geegle.org email addresses are allowed.")
	}
	log.Println("Rcpt to:", to)
	s.To = append(s.To, to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		go sendMail(s.From, s.To, b)
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

type Mail struct {
	senderId string
	toIds    []string
}

func sendMail(from string, to []string, data []byte) error {
	for _, r := range to {
		addEmail(from, r, data)
	}
}

func addEmail(sender string, receiver string, subject string, body string, time int64) {
	fmt.Println("Sending an emiaal")
	email := Email{0, sender, receiver, subject, []byte(body), time}
	reqBody, err := json.Marshal(email)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post("https://geemail-backend.corp.geegle.org/api/addmail", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("mail resp %+v", resp)

	rs, _ := httputil.DumpResponse(resp, true)

	fmt.Println(string(rs))

}

func main() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":1025"
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
