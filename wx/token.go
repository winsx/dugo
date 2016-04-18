package wx

import(
    "io/ioutil"
    "io"
)

type WxTocken interface {
    AccessToken()
    UpdateAccessToken()
}

func (wx *WeChat) AccessToken() (accessToken string) {
    if wx.accessToken == "" {
        wx.UpdateAccessToken()
    }
    return wx.accessToken
}

func (wx *WeChat) UpdateAccessToken() {
    url := WxAccessTockenUrl(wx.GrantType, wx.AppID, wx.AppSecret)
     resp, err := wx.client.Get(url)
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