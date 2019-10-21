package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"io"

	geemail "geegle.org/infra/geemail-client"
	"github.com/emersion/go-smtp"
	"github.com/jhillyerd/enmime"
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
	To   map[string]struct{}
}

func (s *Session) Mail(from string) error {
	if !emailRe.MatchString(from) {
		return errors.New("Invalid email address")
	}
	if strings.Contains(from, "geegle.org") {
		return errors.New("If you are a geegler, please use https://mail.corp.geegle.org instead")
	}
	s.Reset()
	s.From = from
	return nil
}

func (s *Session) Rcpt(to string) error {
	if !emailRe.MatchString(to) {
		return errors.New("Invalid email address")
	}
	if !strings.HasSuffix(to, "@geegle.org") {
		return errors.New("Only geegle.org email addresses are allowed.")
	}
	if _, e := s.To[to]; e {
		return errors.New("Duplicate email address")
	}
	var m struct{}
	s.To[to] = m
	return nil
}

func (s *Session) Data(r io.Reader) error {
	env, err := enmime.ReadEnvelope(r)
	if err != nil {
		return err
	}
	body := env.HTML
	if body == "" {
		body = env.Text
	}
	// NOTE: all other headers are discarded
	// NOTE: attachments are discarded
	// NOTE: spf, dkim, dmarc are not checked
	go sendMail(s.From, s.To, env.GetHeader("Subject"), []byte(body))
	return nil
}

func (s *Session) Reset() {
	s.To = make(map[string]struct{})
	s.From = ""
}

func (s *Session) Logout() error {
	return nil
}

func sendMail(from string, to map[string]struct{}, subject string, data []byte) {
	fmt.Println("sending email")
	fmt.Println(from)
	fmt.Println(to)
	fmt.Println(subject)
	fmt.Println(string(data))
	for r := range to {
		geemail.SendEmailNow(from, r, subject, data)
	}
}

func main() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":25"
	s.Domain = "geegle.org"
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
