package main

import (
    "database/sql"
)

// Anda dapat mengubah-ubah konfigurasi ini
// Sesuaikan saja dengan keadaan server pada komputer anda
var dbhost = "localhost:1433"
var dbuser = "SA"
var dbpass = "MsSQL-Admin"
var dbname = "DBMSProject"

var db *sql.DB

