package geemail

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Email struct {
	ID       int    `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Subject  string `json:"subject"`
	Body     []byte `json:"body"`
	Time     int64  `json:"time"`
}

func SendEmail(sender string, receiver string, subject string, body []byte, time int64) error {
	email := Email{0, sender, receiver, subject, body, time}
	reqBody, err := json.Marshal(email)
	if err != nil {
		return err
	}

	_, err = http.Post("https://geemail-backend.corp.geegle.org/api/addmail", "application/json", bytes.NewBuffer(reqBody))
	return err

}
