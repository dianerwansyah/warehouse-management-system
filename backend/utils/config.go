package utils

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

// Global database connection variable
var DB *sql.DB

// Config struct to hold configuration values read from the YAML file
type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

// read config from YAML file
func LoadConfig() (Config, error) {
	config := Config{}
	// Read content of YAML configuration file
	data, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal YAML data into	 Config struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	return config, nil
}

// connaction database
func Connection(config Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	// Open a connection to database
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed connection database: %w", err)
	}

	// Check if database is reachable
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed ping database: %w", err)
	}

	return nil
}

// EnsureDBConnection checks if database connection is established, and attempts to reconnect if it's nil
func EnsureDBConnection(config Config) error {
	if DB == nil {
		log.Println("Database connection is nil, attempting to reconnect...")
		return Connection(config)
	}
	return nil
}
