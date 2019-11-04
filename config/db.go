package config

import "database/sql"
import "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")
