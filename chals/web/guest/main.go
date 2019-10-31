package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/clsung/grcode"
	"github.com/gorilla/mux"
	"github.com/skip2/go-qrcode"
)

func parseFile(r *http.Request) ([]byte, error) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		return nil, errors.New("File format not understood. Must be jpg/png file")
	}
	defer file.Close()

	if handler.Header.Get("Content-Type") != "image/jpeg" && handler.Header.Get("Content-Type") != "image/png" {
		return nil, errors.New(fmt.Sprintf("File must be JPG/PNG, it is a %s", handler.Header.Get("Content-Type")))
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("Error reading file")
	}

	return fileBytes, nil
}

type ScanData struct {
	Flash    string
	Response string
}

type VisitorData struct {
	Name string
	Date string
	Host string
}

func createVisitor(name string, date time.Time, host string) VisitorData {
	return VisitorData{name, date.Format("2006-01-02 3:04 pm"), host}
}

var defaultVisitors = []VisitorData{
	createVisitor("Jake Haddy", time.Now().Add(40*time.Hour), "guest"),
	createVisitor("Chris Filler", time.Now().Add(3*time.Hour).Add(18*time.Minute), "guest"),
	createVisitor("hacker 121", time.Now().Add(14*time.Hour).Add(52*time.Minute), "guest"),
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(os.Args[1]))))

	r.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/test/qrgen"))
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var inter struct {
			Flash string
		}

		if flash_cookie, err := r.Cookie("flash"); err == nil {
			inter.Flash = flash_cookie.Value
			flash_cookie.Expires = time.Now()
			flash_cookie.Raw = ""
			http.SetCookie(w, flash_cookie)
		}

		RenderTemplate(w, "index.html", inter)

	})

	r.Methods("GET").Subrouter().HandleFunc("/scan", func(w http.ResponseWriter, r *http.Request) {
		var data ScanData
		RenderTemplate(w, "scan.html", data)
	})

	r.Methods("POST").Subrouter().HandleFunc("/scan", func(w http.ResponseWriter, r *http.Request) {
		file, err := parseFile(r)
		var data ScanData
		if err != nil {
			data.Flash = err.Error()
			RenderTemplate(w, "scan.html", data)
			return
		}

		img, img_type, err := image.Decode(bytes.NewReader(file))
		if err != nil || (img_type != "png" && img_type != "jpeg") {
			fmt.Println(err)
			if err == nil {
				data.Flash = "File must be a jpg or png"
			} else {
				data.Flash = err.Error()
			}
			RenderTemplate(w, "scan.html", data)
			return
		}

		results, err := grcode.GetDataFromImage(img)
		if err != nil || len(results) == 0 {
			data.Flash = "Photo is not a qrcode"
			RenderTemplate(w, "scan.html", data)
			return
		}

		cmd, err := exec.Command(os.Args[3], results[0]).Output()
		if err != nil {
			fmt.Println(err)
			data.Flash = err.Error()
			RenderTemplate(w, "scan.html", data)
			return
		}

		fmt.Println(string(cmd), err)
		if "access granted" == string(cmd) {
			cookie := http.Cookie{Name: "flash", Value: "Nice. GEEGLE{A93D9QD39D}", Expires: time.Now().Add(30 * time.Second)}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", 302)
		} else {
			data.Flash = string(cmd)
			RenderTemplate(w, "scan.html", data)
		}
	})

	r.Methods("GET").Subrouter().HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var inter interface{}
		RenderTemplate(w, "login.html", inter)
	})

	r.Methods("POST").Subrouter().HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{Name: "flash", Value: "We have informed your host, they won't be coming down soon", Expires: time.Now().Add(30 * time.Second)}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	})

	r.Methods("GET").Subrouter().HandleFunc("/pickup/{host}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var data struct {
			Flash    string
			Visitors []VisitorData
		}

		if vars["host"] == "guest" {
			data.Visitors = defaultVisitors
		} else {
			data.Flash = fmt.Sprintf("User '%s' not found", vars["host"])
		}

		RenderTemplate(w, "pickup.html", data)
	})

	r.HandleFunc("/test/qrgen", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Missing subpath... try /test/qrgen/abcde"))
	})

	r.HandleFunc("/test/qrgen/{code}", func(w http.ResponseWriter, r *http.Request) {
		var png []byte
		png, err := qrcode.Encode(mux.Vars(r)["code"], qrcode.Medium, 256)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		RenderTemplate(w, "qr.html", base64.StdEncoding.EncodeToString(png))
	})

	http.ListenAndServe(":80", r)
}
