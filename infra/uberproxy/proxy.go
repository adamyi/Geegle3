package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
        "fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func handleUP(rsp http.ResponseWriter, req *http.Request) {
	initUPRsp(rsp)

	if req.Host == "login.corp.geegle.org" {
		handleLogin(rsp, req)
		return
	}

	if req.Host == "geegle.org" || req.Host == "www.geegle.org" {
		req.Host = "search.corp.geegle.org"
	}

	username := getUsername(req)

	ctx, levelShift, err := getNetworkContext(req, username)
	if err != nil {
                fmt.Println("getNetowrkContext - ",  levelShift, err)
		returnError(UPError{Code: http.StatusBadRequest, Title: "Could not resolve the IP address for host " + req.Host, Description: "Your client has issued a malformed or illegal request."}, rsp)
		return
	}

	full_url := req.Host + req.RequestURI
        fmt.Println("getNetworkContext", levelShift, full_url)

	// TODO: allow anonymous access to some services
	// TODO: not hard code search
	// TODO: zero-trust ac policy
	if username == "anonymous@services.geegle.org" && req.Method != "OPTIONS" && req.Host != "search.corp.geegle.org" {
		http.Redirect(rsp, req, "https://login.corp.geegle.org/?return_url="+url.QueryEscape("https://"+full_url), http.StatusTemporaryRedirect)
		return
	}

	servicename := strings.Split(ctx.Value("up_real_addr").(string), ".")[0]

	if levelShift {
		servicename = "uberproxy"
	}

	servicename += "@services.geegle.org"

	ptstr := ""
	expirationTime := time.Now().Add(5 * time.Minute)
	pclaims := Claims{
		Username: username,
		Service:  servicename,
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

	if req.URL.Path == "/ws" {
		handleWs(ctx, rsp, req, ptstr, levelShift)
		return
	}

	bodyBytes, _ := ioutil.ReadAll(req.Body)
	// req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	scheme := "http://"
	if levelShift {
		scheme = "https://"
	}

	preq, err := http.NewRequestWithContext(ctx, req.Method, scheme+full_url, bytes.NewReader(bodyBytes))
	if err != nil {
		returnError(UPError{Code: http.StatusBadGateway, Title: "Bad Gateway", Description: "Something went wrong connecting to internal service"}, rsp)
		return
	}

	for name, value := range req.Header {
		val := value[0]
		ln := strings.ToLower(name)
		if ln == "cookie" {
			cookies := strings.Split(val, ";")
			l := len(cookies)
			for i, cookie := range cookies {
				if strings.HasPrefix(strings.TrimLeft(strings.ToLower(cookie), " "), "uberproxy_auth") {
					cookies[i] = cookies[l-1]
					l -= 1
				}
			}
			if l > 0 {
				val = strings.TrimLeft(strings.Join(cookies[:l], ";"), " ")
			} else {
				val = ""
			}
		} else if ln == "x-geegle-subacc" {
			continue
		}
		if val != "" {
			preq.Header.Set(name, val)
		}
	}

	preq.Header.Set("X-Geegle-JWT", ptstr)
	if levelShift {
		preq.Header.Set("X-UberProxy-LevelShift", "1")
	}

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
