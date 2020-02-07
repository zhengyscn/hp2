package config

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	err        error
	cfg        *ini.File
	configPath string = "config/hp2.ini"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Master struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     int
	Name     string
	Enable   bool
}

type Slave struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     int
	Name     string
	Enable   bool
}

type Casbin struct {
	RbacConfigPath string
	PolicyPath     string
}

var ServerConfig = &Server{}
var MasterDatabaseConfig = &Master{}
var SlaveDatabaseConfig = &Slave{}
var CasbinDatabaseConfig = &Casbin{}

func structToMap(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo err: %v", err)
	}
}

func init() {
	cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalf("Config Setup, fail to parse %s: %v", configPath, err)
	}

	structToMap("server", ServerConfig)
	structToMap("master", MasterDatabaseConfig)
	structToMap("slave", SlaveDatabaseConfig)
	structToMap("casbin", CasbinDatabaseConfig)
}
