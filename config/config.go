package config

import (
	"os"

	"gopkg.in/yaml.v3"
)


// DATABASE CONFIG
type DBConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
}


// JWT CONFIG
type JWTConfig struct {
	AccessSecretKey  string `yaml:"access_secret_key"`
	RefreshSecretKey string `yaml:"refresh_secret_key"`
	AccessTTLMinutes int    `yaml:"access_ttl_minutes"`
	RefreshTTLHours  int    `yaml:"refresh_ttl_hours"`
}


// OTP CONFIG
type OTPConfig struct {
	ExpiryMinutes int `yaml:"expiry_minutes"`
}


// SMTP CONFIG
type SMTPConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}


// REDIS CONFIG
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}


// MAIN APP CONFIG
type Config struct {
	DB    DBConfig    `yaml:"db"`
	JWT   JWTConfig   `yaml:"jwt"`
	OTP   OTPConfig   `yaml:"otp"`
	SMTP  SMTPConfig  `yaml:"smtp"`
	Redis RedisConfig `yaml:"redis"`
}

// Global config instance
var AppConfig *Config

// CONFIG LOADER
// LoadConfig reads YAML file and converts it into Config struct
func LoadConfig(path string) (*Config, error) {

	cfg := &Config{}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}