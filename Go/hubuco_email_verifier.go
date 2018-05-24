package main

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

const
    // !!! PUT YOUR API KEY HERE !!!
    api_key = "YOUR_API_KEY"


type RespType struct {
    ResultCode int `json:"resultcode"`
    Error string `json:"error"`
}


func main() {
    if len(os.Args) != 2 {
    fmt.Printf("usage: %s <email>\n", os.Args[0])
    os.Exit(1)
    }

    email := os.Args[1]

    response, err := http.Get("https://api.hubuco.com/api/v3/?api="+api_key+"&"+"email="+email)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }

    res := RespType{}
    if err := json.Unmarshal([]byte(contents), &res); err != nil {
        panic(err)
    }

    switch res.ResultCode {
    case 1:
        fmt.Println("Ok")
    case 2:
        fmt.Println("Catch All")
    case 3:
        fmt.Println("Unknown")
    case 4:
        fmt.Println("Error: " + res.Error)
    case 5:
        fmt.Println("Disposable")
    case 6:
        fmt.Println("Invlaid")
    }

}

