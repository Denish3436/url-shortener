package config

import (
    "os"
)

type Config struct {
    Port        string
    DBConnection string
    CacheTTL    int
}

func GetConfig() *Config {
    return &Config{
        Port:        os.Getenv("PORT"), // Read PORT from environment variables
        DBConnection: os.Getenv("DB_CONNECTION"), // Database connection string
        CacheTTL:    60,
    }
}

