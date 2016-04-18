package main

import(
    "fmt"
    "os"
)

func main() {
    fmt.Printf("Be listening in %v:%v will access from internet address of https://dugo-gility.c9users.io", os.Getenv("IP"), os.Getenv("PORT"))
}