package wx

import(
    "net/http"    
)

func newClient() (client *http.Client) {
    tr := &http.Transport{
	DisableCompression: true,
    }
    return &http.Client{Transport: tr}
}