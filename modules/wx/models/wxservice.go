package models

import "net/http"

type WxService struct {
    Token string
    ResponseWriter http.ResponseWriter
    Request *http.Request
}