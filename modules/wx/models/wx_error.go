package models

type WxError struct {
    ErrorCode int `json:"errcode"`
    ErrorMsg string `json:"errmsg"`
}