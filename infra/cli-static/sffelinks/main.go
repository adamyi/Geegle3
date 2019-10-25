package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cgi"
	"os"
)

type Flag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type StoreRequest struct {
	FileName string `json:"filename"`
	Content  []byte `json:"content"`
	Flags    []Flag `json:"flags"`
}

type SFFEMsg struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
	Url   string `json:"url"`
}

func DoStoreFile(req *StoreRequest) (string, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("https://sffe.corp.geegle.org/api/store", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result SFFEMsg
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	if !result.Ok {
		return "", errors.New(result.Error)
	}
	return result.Url, nil
}

func GetFileLinks(username string, filepath string) []string {
	var userflag Flag
	userflag.Name = "ldap"
	userflag.Value = username
	var sr []StoreRequest
	jsonFile, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return nil
	}
	jsonContent, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
		return nil
	}
	err = json.Unmarshal(jsonContent, &sr)
	if err != nil {
		log.Println(err)
		return nil
	}
	var urls []string
	for _, r := range sr {
		storedFile, err := os.Open("/" + r.FileName)
		if err != nil {
			log.Println(err)
			continue
		}
		r.Content, err = ioutil.ReadAll(storedFile)
		if err != nil {
			log.Println(err)
			continue
		}
		r.Flags = append(r.Flags, userflag)
		url, err := DoStoreFile(&r)
		if err != nil {
			log.Println(err)
			continue
		}
		urls = append(urls, url)
	}
	return urls
}

func CGIHandler(w http.ResponseWriter, r *http.Request) {
	tknStr := r.Header.Get("X-Geegle-JWT")
	user, err := getUsername(tknStr, []byte("superSecretJWTKEY"))
	check(err, "authentication error")
	urls := GetFileLinks(user, "/clisffe.json")
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(urls)
	if err != nil {
		log.Panic(err)
	}
}

func check(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func main() {
	err := cgi.Serve(http.HandlerFunc(CGIHandler))
	check(err, "cannot serve request")
}
