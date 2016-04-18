package service

import(
    "fmt"   
    "net/http"
)

type IndexHandler struct {
    
}


func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    index(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
    fmt.Println("Handle: home,echo hello world")
}
