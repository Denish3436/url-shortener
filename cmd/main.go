
package main

import (
    "log"
    "url-shortener/config"
)

func main(
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    cfg := config.GetConfig()

    db, err := sql.Open("postgres", cfg.DBConnection)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
