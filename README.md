# graduation

[![Floobits Status](https://floobits.com/nicholasjackson/wwg_graduation.svg)](https://floobits.com/nicholasjackson/wwg_graduation/redirect)

Graduation Assigment for the learn to code course

## Example of go code in pusher

```go
import "github.com/pusher/pusher-http-go"

client := pusher.Client{
  AppId: "APP_ID",
  Key: "APP_KEY",
  Secret: "APP_SECRET",
}

client.Trigger(channels, event, data)
```

## Tasks
* Contact remote API and retrieve the Plans - encoded JSON document
- net/http, encoding/json, crypto

* Create something which connects to pusher and can forward requests to sphero
- pusherAPI, GoBot.io

* Create an API which can send requests to pusher using JSON
- net/http encoding/json, pusherAPI

* Write a command line tool to read the plans and send them to the JSON API
- os, net/http, encoding/JSON, loops,
