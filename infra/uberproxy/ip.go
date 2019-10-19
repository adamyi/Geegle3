package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

var privateIPBlocks []*net.IPNet

func initPrivateIP() {
	for _, cidr := range []string{
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"100.64.0.0/10",  // RFC6598
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(fmt.Errorf("parse error on %q: %v", cidr, err))
		}
		privateIPBlocks = append(privateIPBlocks, block)
	}
}

func isDockerIP(ip net.IP) bool {
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

// hacky string breakdown for ptr record to look for docker network name
// TODO: there's probably a more elegant way way to do this but ceebs
func getServiceNameFromIP(ip string) (string, error) {
	// fmt.Println(ip)
	pip := net.ParseIP(ip)
	if pip == nil || !isDockerIP(pip) {
		return "", errors.New("not geegle service")
	}
	rdns, err := net.LookupAddr(ip)
	if err != nil {
		return "", err
	}
	if len(rdns) == 0 {
		return "", errors.New("no ptr record")
	}
	parts := strings.Split(rdns[0], ".")
	dockernet := parts[len(parts)-2]
	bcp := strings.Split(dockernet, "beyondcorp_")
	if len(bcp) != 2 {
		return "", errors.New("not beyondcorp service")
	}
	return bcp[1] + "@services.geegle.org", nil
}
