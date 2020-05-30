package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "fmt"
)

var DbConnection *sql.DB

type Person struct {
    Name string
    Age int
}

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

    // cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
    // _, err = DbConnection.Exec(cmd, "Nancy", 20)
    // if err != nil{
    //     log.Fatalln(err)
    // }

    // cmd = "UPDATE person SET age = ? WHERE name = ?"
    // _, err = DbConnection.Exec(cmd, 25, "Mike")
    // if err != nil{
    //     log.Fatal(err)
    // }

    // cmd = "SELECT * FROM person"
    // rows, _ := DbConnection.Query(cmd)
    // defer rows.Close()

    // var pp []Person
    // for rows.Next() {
    //     var p Person
    //     err := rows.Scan(&p.Name, &p.Age)
    //     if err != nil{
    //         log.Fatalln(err)
    //     }
    //     pp = append(pp, p)
    // }
    // err = rows.Err()
    // if err != nil{
    //     log.Fatalln(err)
    // }
    // for _,p := range pp{
    //     fmt.Println(p.Name, p.Age)
    // }
}