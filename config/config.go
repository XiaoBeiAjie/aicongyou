package config

import (
    "fmt"
    "os"
    "gopkg.in/yaml.v3"
)

type Config struct {
    Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Name      string `yaml:"name"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
}

var Global *Config

func Load(configFile string) error {
    file, err := os.ReadFile(configFile)
    if err != nil {
        return fmt.Errorf("failed to read config file: %w", err)
    }

    Global = &Config{}
    err = yaml.Unmarshal(file, Global)
    if err != nil {
        return fmt.Errorf("failed to parse config file: %w", err)
    }

    return nil
}

func (c *DatabaseConfig) GetDSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
        c.User,
        c.Password,
        c.Host,
        c.Port,
        c.Name,
        c.Charset,
        c.ParseTime,
        c.Loc,
    )
}

