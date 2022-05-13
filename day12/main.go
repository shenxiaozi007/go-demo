package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Account struct {
    Email    string  `json:"email"`
    Password string  `json:"password"`
    Money    float64 `json:"money,string"`
}

var jsonString string = `{
    "email": "phpgo@163.com",
    "password" : "123456",
    "money" : "100.5"
}`

func main() {

    account := Account{}
    err := json.Unmarshal([]byte(jsonString), &account)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", account)
}