package sign

import(
    "fmt"
    "net/http"
    "github.com/alimy/dugo/modules/wx/models"
)
type Sign struct {
    // now empty
}

func (*Sign) Serve(s *models.WxService) {
    wxSign(s.Token, s.ResponseWriter, s.Request)
}

func (*Sign) Serve2(token string, w http.ResponseWriter) {
    fmt.Fprintln(w, token)
}

func wxSign(token string, w http.ResponseWriter, r *http.Request) {
    echoStr := r.FormValue("echostr")
    content := fmt.Sprintf("token: %s\n echostr: %s", token, echoStr)
    fmt.Fprintln(w, content)
    fmt.Println(content)
}