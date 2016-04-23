package main

import(
    "fmt"
    "net/http"
    "time"
    "log"
    "github.com/alimy/dugo/services/wx"
    "github.com/alimy/dugo/routers"
    "github.com/alimy/dugo/services/wx/services/token"
)

var(
    server = &http.Server{
	    Addr:           ":8080",
	    Handler:        nil,
	    ReadTimeout:    10 * time.Second,
	    WriteTimeout:   10 * time.Second,
	    MaxHeaderBytes: 1 << 20,
    }
)

func run() {
    sm := http.NewServeMux()
    sm.Handle("/", &routers.HomeHandler{})
    sm.Handle("/wx", wx.Wx.Handler)
    sm.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request){
        token := token.Token.AccessToken()
        fmt.Fprintf(w, "%s", token)
        fmt.Println("%s", token)
    })
    sm.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request){
        http.Redirect(w, r, "http://www.baidu.com/", http.StatusTemporaryRedirect)
    })
    server.Handler = sm
    
    log.Fatal(server.ListenAndServe())
}

func main() {
    run()
}