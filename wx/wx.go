package wx

import(
    "fmt"
    "net/http"
    "github.com/alimy/dugo/wx/config"
    "github.com/alimy/dugo/wx/serve"
)

const(
    wxAccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s"
)

var Wx *WeChat

type WeChat struct {
    config.WxConfig
    serve.WxHandler
    accessToken string
    client *http.Client
}

func init() {
    Wx = NewWx()
}

func NewWx() (wx *WeChat) {
    wx = &WeChat{WxConfig:
        config.WxConfig{
            AppID: "wx609f9cb8513d8552",
            AppSecret: "8e801592ee58eaaeb6f11668d81558f8",
            GrantType: "client_credential",
        },
    }

    wx.client = newClient()
    wx.UpdateAccessToken()
    
    return wx
}

func WxAccessTockenUrl(grantType, appId, secret string) string {
    return fmt.Sprintf(wxAccessTokenUrl, grantType, appId, secret)
}