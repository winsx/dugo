package config

var Config *WxConfig

type WxConfig struct {
    AppID string
    AppSecret string
    GrantType string
}

func init() {
    Config = &WxConfig{
        AppID: "wx609f9cb8513d8552",
        AppSecret: "8e801592ee58eaaeb6f11668d81558f8",
        GrantType: "client_credential",
        }
}
