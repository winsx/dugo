package routers

import(
    "fmt"
    "net/http"
)

type WxHandler struct {
    
}

func (h *WxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    wxSign(w, r)
}

func wxSign(w http.ResponseWriter, r *http.Request) {
    echoStr := r.FormValue("echostr")
    fmt.Fprintf(w, "%s", echoStr)
    fmt.Printf("echo back string of %s\n", echoStr)
}