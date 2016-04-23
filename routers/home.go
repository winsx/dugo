package routers

import(
    "fmt"   
    "net/http"
)

type HomeHandler struct {
    
}


func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    home(w, r)
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
    fmt.Println("Handle: home,echo hello world")
}
