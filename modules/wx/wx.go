package wx

import(
    "github.com/alimy/dugo/modules/wx/routers"
    _ "github.com/alimy/dugo/modules/wx/config"
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
}

func newWx() *WeChat {
    return &WeChat{
        Handler: routers.NewWxHandler(),
    }
}