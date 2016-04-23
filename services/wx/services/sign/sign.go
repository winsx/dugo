package sign

import(
    "fmt"
    "net/http"
    "github.com/alimy/dugo/services/wx/content"
)
type Sign struct {
    // now empty
}

func (*Sign) Serve(s *content.WxService) {
    wxSign(s.Token, s.ResponseWriter, s.Request)
}

func wxSign(token string, w http.ResponseWriter, r *http.Request) {
    echoStr := r.FormValue("echostr")
    content := fmt.Sprintf("token: %s\n echostr: %s", token, echoStr)
    fmt.Fprintf(w, "token:")
    fmt.Println(content)
}