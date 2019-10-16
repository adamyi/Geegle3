package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

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
	return dialer.DialContext(ctx, network, address)
}

func getRealAddr(host string, player string) (string, bool, error) {
	if host != "geegle.org" && (!strings.HasSuffix(host, ".geegle.org")) {
		return "", false, errors.New("not geegle domain")
	}

	if strings.HasSuffix(host, ".apps.geegle.org") {
		host = "apps.geegle.org"
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		return getL2Addr(player)
	}

	for _, ip := range ips {
		if !isDockerIP(ip) {
			return getL2Addr(player)
		}
	}

	return host + ":80", false, nil
}

func getL2Addr(player string) (string, bool, error) {
	if os.Getenv("UBERPROXY_MASTER") != "" {
		return player + ".prod.geegle.org:443", true, nil
	}
	return "master.prod.geegle.org:443", true, nil
}

func getNetworkContext(req *http.Request, username string) (context.Context, bool, error) {
	player := strings.Split(username, "@")[0]
	addr, levelshift, err := getRealAddr(req.Host, player)
	if err != nil {
		return context.Background(), false, err
	}
	if levelshift {
		if req.Header.Get("X-UberProxy-LevelShift") == "1" {
			return context.Background(), false, errors.New("domain not present in two-level UP infra")
		}
	}
	return context.WithValue(context.Background(), "up_real_addr", addr), levelshift, nil
}

func initNetworkOptions() {
	http.DefaultTransport.(*http.Transport).DialContext = upDialContext
	websocket.DefaultDialer.NetDialContext = upDialContext
}
