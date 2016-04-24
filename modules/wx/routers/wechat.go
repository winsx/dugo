package routers

import(
    "time"
    "net/http"
    "github.com/alimy/dugo/modules/wx/modules/token"
    "github.com/alimy/dugo/modules/wx/modules/sign"
    "github.com/alimy/dugo/modules/wx/models"
)

type WxHandler struct {
    Gone chan models.Done
    service chan *models.WxService
    test chan http.ResponseWriter
}

func NewWxHandler() (handler *WxHandler) {
     handler = &WxHandler {
        Gone: make(chan models.Done),
        service: make(chan *models.WxService),
        test: make(chan http.ResponseWriter),
    }
    
    go handler.run()
    
    return handler
}

func (h *WxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s := &models.WxService{ResponseWriter: w, Request: r}
    //go serve(s)
    h.service <- s
    h.test <- w
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
            case w := <- h.test :
                go serve3(token.Token.AccessToken(), w)
        }
    }
}

func serve(s *models.WxService) {
    wxSign := &sign.Sign{}
    wxSign.Serve(s)
}

func serve3(token string, w http.ResponseWriter) {
    wxSign := &sign.Sign{}
    wxSign.Serve2(token, w)
}
