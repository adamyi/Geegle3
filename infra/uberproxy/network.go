package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
        "fmt"
	"github.com/gorilla/websocket"
)

var (
	dialer = &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: false,
	}
)

func upDialContext(ctx context.Context, network, address string) (net.Conn, error) {
	if v := ctx.Value("up_real_addr"); v != nil {
		address = v.(string)
	}
        fmt.Println("updialContext", address)
	return dialer.DialContext(ctx, network, address)
}

func getRealAddr(host string) (string, error) {
	if host != "geegle.org" && (!strings.HasSuffix(host, ".geegle.org")) {
		return "", errors.New("not geegle domain")
	}

	if strings.HasSuffix(host, ".apps.geegle.org") {
		host = "apps.geegle.org"
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}

	for _, ip := range ips {
		if !isDockerIP(ip) {
			return "", errors.New("not geegle internal")
		}
	}

	return host + ":80", nil
}

func getL2Addr(player string) (string, error) {
	if os.Getenv("UBERPROXY_CLUSTER") != "master" {
		player = "master"
	}
	host := player + ".prod.geegle.org"
	ips, err := net.LookupIP(host)
        fmt.Println("getL2", ips[0].String, err, host)
	if err != nil || len(ips) == 0 {
		return "", errors.New("not valid geegle")
	}

	return ips[0].String() + ":443", nil
}

func getNetworkContext(req *http.Request, username string) (context.Context, bool, error) {
	addr, err := getRealAddr(req.Host)
	if err == nil {
		return context.WithValue(context.Background(), "up_real_addr", addr), false, nil
	}

	if req.Header.Get("X-UberProxy-LevelShift") == "1" {
		return context.Background(), false, errors.New("domain not present in two-level UP infra")
	}

	players := strings.Split(strings.Split(username, "@")[0], "+")
	hp := req.Header.Get("X-UberProxy-Player")
	if hp != "" {
		players = append(players, hp)
	}
	for _, player := range players {
		addr, err = getL2Addr(player)
		if err == nil {
			return context.WithValue(context.Background(), "up_real_addr", addr), true, nil
		}
	}
	return context.Background(), false, errors.New("not found anywhere")
}

func initNetworkOptions() {
	http.DefaultTransport.(*http.Transport).DialContext = upDialContext
	websocket.DefaultDialer.NetDialContext = upDialContext
}
