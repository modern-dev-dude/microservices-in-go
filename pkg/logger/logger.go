package logger

import (
	"errors"
	"fmt"
	"net"
	"net/http"
)

type LogBody struct {
	id     string
	ip     string
	method string
}

func WriteLogToConsole(r *http.Request, reqId string) {
	ip, err := getIp(r)
	// silently fail if the ip is not found
	if err != nil {
		fmt.Printf("Error getting IP: %v", err)
	}

	if err != nil {
		fmt.Printf("Error getting body: %v", err)
	}

	requestInfo := LogBody{
		id:     reqId,
		ip:     ip,
		method: r.Method,
	}

	fmt.Printf("Request body: %v", requestInfo)
}

func getIp(r *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "error getting ip", err
	}

	netIp := net.ParseIP(ip)
	if netIp != nil {
		return ip, nil
	}

	return "error getting ip", errors.New("IP not found")
}
