package config

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	Server   = &server{}
	Redis    = &redis{}
	Database = &database{}
)

type server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type database struct {
	Type     string
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	Charset  string
}

type redis struct {
	Address        string
	Password       string
	DataBaseType   int
	ExpirationTime time.Duration
}

func Setup() {
	config, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatalf("setting error: %v", err)
	}

	settings := map[string]interface{}{
		"server":   Server,
		"redis":    Redis,
		"database": Database,
	}
	for key, settingStruct := range settings {
		if err := config.Section(key).MapTo(settingStruct); err != nil {
			log.Fatalf("setting error: %v", err)
		}
	}
}
