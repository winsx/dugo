package api

import(
    "fmt"
)

const(
    wxAccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s"
)

func WxAccessTockenUrl(grantType, appId, secret string) string {
    return fmt.Sprintf(wxAccessTokenUrl, grantType, appId, secret)
}