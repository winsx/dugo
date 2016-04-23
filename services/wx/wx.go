package wx

import(
    "github.com/alimy/dugo/services/wx/config"
    "github.com/alimy/dugo/services/wx/routers"
    "github.com/alimy/dugo/services/wx/services/token"
    _ "github.com/alimy/dugo/services/wx/services/client"
)

type WeChat struct {
    token.WxToken
    routers.WxHandler
}

func init() {
    config.Config.Update("wx609f9cb8513d8552", 
                "8e801592ee58eaaeb6f11668d81558f8",  "client_credential") 
}

func Classic() (wx *WeChat) {
    return &WeChat{}
}