package main

import(
    "fmt"
    "net/http"
    "time"
    "log"
    "github.com/alimy/dugo/wx"
    "github.com/alimy/dugo/service"
)

var server = &http.Server{
	Addr:           ":8080",
	Handler:        nil,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

func run() {
    sm := http.NewServeMux()
    sm.Handle("/", &service.IndexHandler{})
    sm.Handle("/wx", wx.Wx)
    sm.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request){
        token := wx.Wx.AccessToken()
        fmt.Fprintf(w, "%s", token)
        fmt.Println("%s", token)
    })
    server.Handler = sm
    
    log.Fatal(server.ListenAndServe())
}

func main() {
    run()
}