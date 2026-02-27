package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type DBConfig struct{
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	DBName string `yaml:"dbname"`
	Port int `yaml:"port"`
	SSLMode string `yaml:"sslmode"`
}

type JWTConfig struct{
	AccessSecretKey string `yaml:"access_secret_key"`
	RefreshSecretKey string `yaml:"refresh_secret_key"`
	AccessTTLMinutes int  `yaml:"access_ttl_minutes"`
	RefreshTTLHours  int  `yaml:"refresh_ttl_hours"`
}

type Config struct{
	DB DBConfig `yaml:"db"` //DB = Struct value and stuct field (A variable is a named storage location that holds a value.)
	JWT JWTConfig `yaml:"jwt"`
}

var AppConfig *Config

func LoadConfig(path string) (*Config, error){
	cfg := &Config{} //Create a Config object and store its pointer in cfg

	file, err := os.ReadFile(path)// in here go opens the file path and read all the context(like text, img, pdf etc) then stores it in memory in here file and returns bytes
	if err != nil{
		return nil, err
	}

	err = yaml.Unmarshal(file, cfg) // Unmarshal means: Convert bytes → Go struct. in here file contains yaml data 
	if err != nil{					//-> YAML parser reads bytes -> It looks at your struct(Config) -> It matches YAML keys with tags -> It writes values directly into memory using the pointer.
		return nil, err
	}
	
	return cfg, nil
}