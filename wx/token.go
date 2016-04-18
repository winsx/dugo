package wx

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
     resp, err := wx.client.Get(WxAccessToken(grantType, appId, secret))
    if err != nil {
	// handle error
    }
    defer resp.Body.Close()
    
    wx.accessToken = parseTokenFrom(resp)
}

func parseTokenFrom(r io.Reader) (token string) {
    body, err := ioutil.ReadAll(resp.Body)
    
    return string(body)
}