package token

import(
    "io/ioutil"
    "io"
    "github.com/alimy/dugo/modules/wx/models/api"
    . "github.com/alimy/dugo/modules/wx/modules/client"
    . "github.com/alimy/dugo/modules/wx/config"
)

var Token *WxToken 

type WxToken struct {
    accessToken string
}

func init() {
    if Token == nil {
        Token = &WxToken{}
    }
}

func (wx *WxToken) AccessToken() (accessToken string) {
    if wx.accessToken == "" {
        wx.UpdateAccessToken()
    }
    return wx.accessToken
}

func (wx *WxToken) UpdateAccessToken() {
    url := api.WxAccessTockenUrl(Config.GrantType, Config.AppID, Config.AppSecret)
    resp, err := Client.Get(url)
    if err != nil {
	// handle error
    }
    defer resp.Body.Close()
    
    wx.accessToken = parseTokenFrom(resp.Body)
}

func parseTokenFrom(r io.Reader) (token string) {
    body, _ := ioutil.ReadAll(r)
    
    return string(body)
}