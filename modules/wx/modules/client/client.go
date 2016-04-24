package client

import(
    "net/http"    
)

var Client *http.Client

func init() {
    Client = newClient()    
}


func newClient() (client *http.Client) {
    tr := &http.Transport{
	DisableCompression: true,
    }
    return &http.Client{Transport: tr}
}