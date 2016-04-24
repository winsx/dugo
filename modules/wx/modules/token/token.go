package token

import(
    "io/ioutil"
    "encoding/json"
    "github.com/alimy/dugo/modules/wx/models/api"
    "github.com/alimy/dugo/modules/wx/models"
    . "github.com/alimy/dugo/modules/wx/modules/client"
    . "github.com/alimy/dugo/modules/wx/config"
)

var Token *WxToken 

type WxToken struct {
    models.WxError
    AccessToken string `json:"access_token"`
    ExpiresIn int `json:"expires_in"`
}

func init() {
    if Token == nil {
        Token = &WxToken{}
        Token.UpdateAccessToken()
    }
}


func (wx *WxToken) UpdateAccessToken() {
    url := api.WxAccessTockenUrl(Config.GrantType, Config.AppID, Config.AppSecret)
    resp, err := Client.Get(url)
    if err != nil {
	// handle errors
    }
    defer resp.Body.Close()
    
     body, _ := ioutil.ReadAll(resp.Body)
     
     json.Unmarshal(body, wx)
}