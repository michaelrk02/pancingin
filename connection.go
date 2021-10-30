package main

import (
    "database/sql"
)

type DBConfig struct {
    DBHost string `json:"dbhost"`
    DBUser string `json:"dbuser"`
    DBPass string `json:"dbpass"`
    DBName string `json:"dbname"`
}

var db *sql.DB

