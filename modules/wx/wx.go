package wx

import(
    "github.com/alimy/dugo/modules/wx/config"
    "github.com/alimy/dugo/modules/wx/routers"
    _ "github.com/alimy/dugo/modules/wx/modules/client"
)

var Wx *WeChat

type WeChat struct {
    Handler *routers.WxHandler
}

func init() {
    if Wx == nil {
        Wx = newWx()
    }
    config.Config.Update("wx609f9cb8513d8552", 
                "8e801592ee58eaaeb6f11668d81558f8",  "client_credential")
}

func newWx() *WeChat {
    return &WeChat{
        Handler: routers.NewWxHandler(),
    }
}