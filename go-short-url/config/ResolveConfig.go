package config

import (
	"fmt"
	"github.com/go-ini/ini"
)

type DbConfig struct {
	Driver   string `json:"db_driver"`
	UserName string `json:"db_username"`
	Password string `json:"db_password"`
	Port     string `json:"db_port"`
	DbName   string `json:"db_name"`
	Url      string `json:"db_url"`
}

func ParseDb(path string) (*DbConfig, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		panic(err)
	}
	dbConfig := DbConfig{
		Driver:   cfg.Section("db").Key("db_driver").String(),
		UserName: cfg.Section("db").Key("db_username").String(),
		Password: cfg.Section("db").Key("db_password").String(),
		Port:     cfg.Section("db").Key("db_port").String(),
		DbName:   cfg.Section("db").Key("db_name").String(),
		Url:      cfg.Section("db").Key("db_url").String(),
	}
	return &dbConfig, err
}
