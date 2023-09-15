package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"time"
)

var (
	RunMode      string
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	ImgPath      string
)

func init() {

	// load the ini file
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		fmt.Printf("Fail to parse 'app.ini': %v\n", err)
	}

	// get all the params
	// get run mode
	RunMode = cfg.Section("").Key("RUN_MODE").MustString("release")
	// get http port
	HttpPort = cfg.Section("").Key("PORT").MustInt(8086)
	// get time out (second)
	ReadTimeOutInt := cfg.Section("").Key("READ_TIME_OUT").MustInt(10)
	WriteTimeOutInt := cfg.Section("").Key("WRITE_TIME_OUT").MustInt(10)
	// convert to time.Duration
	ReadTimeOut = time.Duration(ReadTimeOutInt) * time.Second
	WriteTimeOut = time.Duration(WriteTimeOutInt) * time.Second
	// get img path
	ImgPath = cfg.Section("").Key("IMG_PATH").MustString("img")
}
