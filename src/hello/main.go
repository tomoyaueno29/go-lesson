package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var DbConnection *sql.DB

func main(){
    DbConnection, _ := sql.Open("sqlite3", "./example.sql")
    defer DbConnection.Close()
    cmd := `CREATE TABLE IF NOT EXISTS person(
                name String,
                age INT)`
    _, err := DbConnection.Exec(cmd)
    if err != nil{
        log.Fatal(err)
    }

    cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
    _, err = DbConnection.Exec(cmd, "Mike", 24)
    if err != nil{
        log.Fatalln(err)
    }
}