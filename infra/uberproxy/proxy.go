package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
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

	username := getUsername(req)

	ctx, levelShift, err := getNetworkContext(req, username)
	if err != nil {
		returnError(UPError{Code: http.StatusBadRequest, Title: "Could not resolve the IP address for host " + req.Host, Description: "Your client has issued a malformed or illegal request."}, rsp)
		return
	}

	full_url := req.Host + req.RequestURI

	// TODO: allow anonymous access to some services
	if username == "anonymous@services.geegle.org" && req.Method != "INFO" {
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
	if req.Host == "cli-relay.corp.geegle.org" {
		handleCLIRelay(rsp, req)
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
		//TODO: remove uberproxy auth cookie from forwarded request
		preq.Header.Set(name, value[0])
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
