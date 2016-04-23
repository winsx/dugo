package routers

import(
    "time"
    "net/http"
    "fmt"
    "reflect"
    "github.com/alimy/dugo/services/wx/services/token"
    "github.com/alimy/dugo/services/wx/services/sign"
    "github.com/alimy/dugo/services/wx/content"
)

type WxHandler struct {
    Gone chan content.Done
    service chan *content.WxService
}

func NewWxHandler() (handler *WxHandler) {
     handler = &WxHandler {
        Gone: make(chan content.Done),
        service: make(chan *content.WxService),
    }
    
    go handler.run()
    
    return handler
}

func (h *WxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "1.pre-response")
    h.service <- &content.WxService{ResponseWriter: reflect.ValueOf(w).Interface().(http.ResponseWriter), Request: r}
}

func (h *WxHandler) run() {
    for {
        select {
            case <-h.Gone :
                return
            case <-time.After(7200 * time.Second) :
                token.Token.UpdateAccessToken()
            case service := <- h.service :
                service.Token = token.Token.AccessToken()
                go serve(service)
        }
    }
}

func serve(s *content.WxService) {
    wxSign := &sign.Sign{}
    fmt.Fprintln(s.ResponseWriter, "2.pre-response")
    wxSign.Serve(s)
}
