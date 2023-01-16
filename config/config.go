package config

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	Server = &server{}
	Redis  = &redis{}
)

type server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type redis struct {
	Address      string
	Password     string
	DataBaseType int
}

func Setup() {
	config, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("setting error: %v", err)
	}

	settings := map[string]interface{}{
		"server": Server,
		"redis":  Redis,
	}
	for key, settingStruct := range settings {
		if err := config.Section(key).MapTo(settingStruct); err != nil {
			log.Fatalf("setting error: %v", err)
		}
	}
}
