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
    wxSign(s.ResponseWriter, s.Request, s.Token)
}

func (*Sign) Serve2(w http.ResponseWriter, token string) {
    fmt.Fprintln(w, token)
}

func wxSign(w http.ResponseWriter, r *http.Request, token string) {
    echoStr := r.FormValue("echostr")
    content := fmt.Sprintf("token: %s\n echostr: %s", token, echoStr)
    fmt.Fprintln(w, content)
    fmt.Println(content)
}