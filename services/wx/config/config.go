package config

var Config *WxConfig

type WxConfig struct {
    AppID string
    AppSecret string
    GrantType string
}

func init() {
    Config = &WxConfig{}
}

func (c *WxConfig) Update(appID, appSecret, grantType string) {
    c.AppID = appID
    c.AppSecret = appSecret
    c.GrantType = grantType
}
