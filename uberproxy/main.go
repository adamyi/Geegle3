package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Configuration struct {
	ListenAddress string
	JwtKey        []byte
}

type Claims struct {
	Username string `json:"username"`
	Service  string `json:"service"`
	jwt.StandardClaims
}

var _configuration = Configuration{}

var (
	upgrader = websocket.Upgrader{CheckOrigin: checkOrigin}
	dialer   = websocket.DefaultDialer
)

func readConfig() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&_configuration)
	if err != nil {
		panic(err)
	}
}

func checkOrigin(r *http.Request) bool {
	o := r.Header.Get("Origin")
	h := r.Host
	if o == "" || h == "" {
		log.Print("Websocket missing origin and/or host")
		return false
	}
	ou, err := url.Parse(o)
	if err != nil {
		log.Printf("Couldn't parse url: %v", err)
		return false
	}
	if ou.Host != h && ou.Host != "cli-relay.corp.geegle.org" {
		log.Print("Origin doesn't match host")
		return false
	}
	return true
}

func initUPRsp(rsp http.ResponseWriter) {
	rsp.Header().Add("Server", "UberProxy")
}

// adapted from https://github.com/koding/websocketproxy/blob/master/websocketproxy.go
func handleWs(rsp http.ResponseWriter, req *http.Request, jwttoken string) {
	requestHeader := http.Header{}
	requestHeader.Set("Host", req.Host)
	requestHeader.Set("X-Geegle-JWT", jwttoken)
	if origin := req.Header.Get("Origin"); origin != "" {
		requestHeader.Add("Origin", origin)
	}
	for _, prot := range req.Header[http.CanonicalHeaderKey("Sec-WebSocket-Protocol")] {
		requestHeader.Add("Sec-WebSocket-Protocol", prot)
	}
	for _, cookie := range req.Header[http.CanonicalHeaderKey("Cookie")] {
		requestHeader.Add("Cookie", cookie)
	}
	backendURL := *req.URL
	backendURL.Host = req.Host
	backendURL.Scheme = "ws"
	// dial backend
	connBackend, resp, err := dialer.Dial(backendURL.String(), requestHeader)
	if err != nil {
		log.Printf("couldn't dial to remote backend url %s %s", backendURL.String(), err)
		if resp != nil {
			// If the WebSocket handshake fails, ErrBadHandshake is returned
			// along with a non-nil *http.Response so that callers can handle
			// redirects, authentication, etcetera.
			if err := copyResponse(rsp, resp); err != nil {
				log.Printf("couldn't write response after failed remote backend handshake: %s", err)
			}
		} else {
			http.Error(rsp, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
		return
	}
	defer connBackend.Close()

	// Only pass those headers to the upgrader.
	upgradeHeader := http.Header{}
	if hdr := resp.Header.Get("Sec-Websocket-Protocol"); hdr != "" {
		upgradeHeader.Set("Sec-Websocket-Protocol", hdr)
	}
	if hdr := resp.Header.Get("Set-Cookie"); hdr != "" {
		upgradeHeader.Set("Set-Cookie", hdr)
	}

	connPub, err := upgrader.Upgrade(rsp, req, upgradeHeader)
	if err != nil {
		log.Printf("couldn't upgrade %s", err)
		return
	}
	defer connPub.Close()

	errClient := make(chan error, 1)
	errBackend := make(chan error, 1)
	replicateWebsocketConn := func(dst, src *websocket.Conn, errc chan error) {
		for {
			msgType, msg, err := src.ReadMessage()
			if err != nil {
				m := websocket.FormatCloseMessage(websocket.CloseNormalClosure, fmt.Sprintf("%v", err))
				if e, ok := err.(*websocket.CloseError); ok {
					if e.Code != websocket.CloseNoStatusReceived {
						m = websocket.FormatCloseMessage(e.Code, e.Text)
					}
				}
				errc <- err
				dst.WriteMessage(websocket.CloseMessage, m)
				break
			}
			err = dst.WriteMessage(msgType, msg)
			if err != nil {
				errc <- err
				break
			}
		}
	}

	go replicateWebsocketConn(connPub, connBackend, errClient)
	go replicateWebsocketConn(connBackend, connPub, errBackend)

	var message string
	select {
	case err = <-errClient:
		message = "Error when copying from backend to client: %v"
	case err = <-errBackend:
		message = "Error when copying from client to backend: %v"

	}
	if e, ok := err.(*websocket.CloseError); !ok || e.Code == websocket.CloseAbnormalClosure {
		log.Printf(message, err)
	}

}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func copyResponse(rw http.ResponseWriter, resp *http.Response) error {
	copyHeader(rw.Header(), resp.Header)
	rw.WriteHeader(resp.StatusCode)
	defer resp.Body.Close()

	_, err := io.Copy(rw, resp.Body)
	return err
}

func handleUP(rsp http.ResponseWriter, req *http.Request) {
	initUPRsp(rsp)
	if req.Host == "login.corp.geegle.org" {
		handleLogin(rsp, req)
		return
	}
	full_url := req.Host + req.RequestURI
	ptstr := ""
	if req.Method != "OPTIONS" {
		c, err := req.Cookie("uberproxy_auth")
		if err != nil {
			http.Redirect(rsp, req, "http://login.corp.geegle.org/?return_url="+url.QueryEscape("http://"+full_url), 303)
			return
		}
		tknStr := c.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return _configuration.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				log.Println("Signature Invalid")
				http.Redirect(rsp, req, "http://login.corp.geegle.org/?return_url="+url.QueryEscape("http://"+full_url), 303)
				return
			}
			log.Println("JWT Error")
			log.Println(err.Error())
			http.Redirect(rsp, req, "http://login.corp.geegle.org/?return_url="+url.QueryEscape("http://"+full_url), 303)
			return
		}

		if !tkn.Valid {
			log.Println("JWT Invalid")
			http.Redirect(rsp, req, "http://login.corp.geegle.org/?return_url="+url.QueryEscape("http://"+full_url), 303)
			return
		}

		if claims.Service != "uberproxy@services.geegle.org" {
			log.Println(claims.Service)
			http.Redirect(rsp, req, "http://login.corp.geegle.org/?return_url="+url.QueryEscape("http://"+full_url), 303)
			return
		}

		// TODO: check if user has permission to access this site

		expirationTime := time.Now().Add(5 * time.Minute)
		pclaims := Claims{
			Username: claims.Username,
			Service:  claims.Service, //TODO: check service name
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		ptoken := jwt.NewWithClaims(jwt.SigningMethodHS256, pclaims)
		ptstr, err = ptoken.SignedString(_configuration.JwtKey)
		if err != nil {
			fmt.Println(err.Error())
			rsp.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if req.URL.Path == "/ws" {
		handleWs(rsp, req, ptstr)
		return
	}
	if req.Host == "cli-relay.corp.geegle.org" {
		handleCLIRelay(rsp, req)
		return
	}

	preq, err := http.NewRequest(req.Method, "http://"+full_url, req.Body)
	if err != nil {
		fmt.Println("test")
		fmt.Println(err.Error())
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	for name, value := range req.Header {
		//TODO: remove uberproxy auth cookie from forwarded request
		preq.Header.Set(name, value[0])
	}
	preq.Header.Set("X-Geegle-JWT", ptstr)
	client := &http.Client{}
	presp, err := client.Do(preq)
	if err != nil {
		fmt.Println("haha")
		fmt.Println(err.Error())
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	copyResponse(rsp, presp)
}
func handleCLIRelay(rsp http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadFile("templates/ws.html")
	fmt.Fprint(rsp, string(body))
}

func handleLogin(rsp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		//TODO: cache this
		body, _ := ioutil.ReadFile("templates/login.html")
		fmt.Fprint(rsp, string(body))
	} else if req.Method == "POST" {
		req.ParseForm()
		username := strings.ToLower(req.Form.Get("username"))
		if !strings.HasSuffix(username, "@geegle.org") {
			username = username + "@geegle.org"
		}
		_ = req.Form.Get("password")
		//TODO: verify password
		expirationTime := time.Now().Add(24 * 30 * time.Hour)
		pclaims := Claims{
			Username: username,
			Service:  "uberproxy@services.geegle.org",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		ptoken := jwt.NewWithClaims(jwt.SigningMethodHS256, pclaims)
		ptstr, err := ptoken.SignedString(_configuration.JwtKey)
		if err != nil {
			rsp.WriteHeader(http.StatusInternalServerError)
			return
		}
		authcookie := &http.Cookie{Name: "uberproxy_auth", Value: ptstr, HttpOnly: true, Domain: "corp.geegle.org"}
		http.SetCookie(rsp, authcookie)
		http.Redirect(rsp, req, req.URL.Query()["return_url"][0], 303)
	} else {
		rsp.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	readConfig()
	http.HandleFunc("/", handleUP)
	err := http.ListenAndServe(_configuration.ListenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
