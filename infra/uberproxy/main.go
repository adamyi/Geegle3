package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net"
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

type UPError struct {
	Title       string
	Description string
	Code        int
}

func returnError(err UPError, rsp http.ResponseWriter) {
	rsp.WriteHeader(err.Code)
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	tmpl.Execute(rsp, err)
}

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
			if k == "Server" {
				dst.Set(k, v)
			} else {
				dst.Add(k, v)
			}
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
	if req.Host != "geegle.org" && (!strings.HasSuffix(req.Host, ".geegle.org")) {
		returnError(UPError{Code: http.StatusBadRequest, Title: "Could not resolve the IP address for host " + req.Host, Description: "Your client has issued a malformed or illegal request."}, rsp)
		return
	}
	if req.Host == "login.corp.geegle.org" {
		handleLogin(rsp, req)
		return
	}
	full_url := req.Host + req.RequestURI
	if req.Host != "cli-relay.corp.geegle.org" && (!strings.HasSuffix(req.Host, ".apps.geegle.org")) {
		ips, err := net.LookupIP(req.Host)
		if err != nil {
			returnError(UPError{Code: http.StatusBadRequest, Title: "Could not resolve the IP address for host " + req.Host, Description: "Your client has issued a malformed or illegal request."}, rsp)
			return
		}
		for _, ip := range ips {
			if !isDockerIP(ip) {
				returnError(UPError{Code: http.StatusBadRequest, Title: "Could not resolve the IP address for host " + req.Host, Description: "Your client has issued a malformed or illegal request."}, rsp)
				return
			}
		}
	}
	ptstr := ""
	if req.Method != "OPTIONS" {
		c, err := req.Cookie("uberproxy_auth")
		if err != nil {
			http.Redirect(rsp, req, "https://login.corp.geegle.org/?return_url="+url.QueryEscape("https://"+full_url), http.StatusTemporaryRedirect)
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
				http.Redirect(rsp, req, "https://login.corp.geegle.org/?return_url="+url.QueryEscape("https://"+full_url), http.StatusTemporaryRedirect)
				return
			}
			log.Println("JWT Error")
			log.Println(err.Error())
			http.Redirect(rsp, req, "https://login.corp.geegle.org/?return_url="+url.QueryEscape("https://"+full_url), http.StatusTemporaryRedirect)
			return
		}

		if !tkn.Valid {
			log.Println("JWT Invalid")
			http.Redirect(rsp, req, "https://login.corp.geegle.org/?return_url="+url.QueryEscape("https://"+full_url), http.StatusTemporaryRedirect)
			return
		}

		if claims.Service != "uberproxy@services.geegle.org" {
			log.Println(claims.Service)
			http.Redirect(rsp, req, "https://login.corp.geegle.org/?return_url="+url.QueryEscape("https://"+full_url), http.StatusTemporaryRedirect)
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
			returnError(UPError{Code: http.StatusInternalServerError, Title: "Internal Server Error", Description: "Something went wrong while generating JWT"}, rsp)
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
		returnError(UPError{Code: http.StatusBadGateway, Title: "Bad Gateway", Description: "Something went wrong connecting to internal service"}, rsp)
		return
	}
	for name, value := range req.Header {
		//TODO: remove uberproxy auth cookie from forwarded request
		preq.Header.Set(name, value[0])
	}
	preq.Header.Set("X-Geegle-JWT", ptstr)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	presp, err := client.Do(preq)
	if err != nil {
		returnError(UPError{Code: http.StatusBadGateway, Title: "Bad Gateway", Description: "Something went wrong connecting to internal service"}, rsp)
		return
	}
	copyResponse(rsp, presp)
}

func initHTTPOptions() {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: false,
	}

	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		if strings.HasSuffix(addr, ".apps.geegle.org:80") {
			addr = "apps.geegle.org:80"
		}
		return dialer.DialContext(ctx, network, addr)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initPrivateIP()
	initHTTPOptions()
	readConfig()
	go http.ListenAndServe(":80", http.HandlerFunc(redirectSSL))
	server := buildSSLServer()
	server.ListenAndServeTLS("", "")
}
