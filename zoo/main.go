package main

import (
    "fmt"

    "hogemodule/animals"

    "hogemodule/hoge"
)

func main() {
    fmt.Println(AppName())
    fmt.Println(animals.ElephantFeed())
    fmt.Println(animals.MonkeyFeed())
    fmt.Println(animals.RabbitFeed())
    fmt.Println(hoge.HogeStr())
}

