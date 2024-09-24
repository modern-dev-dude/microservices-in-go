package logger

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// course
func init() {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err.Error())
	}
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func CustomError(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

// not course
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
