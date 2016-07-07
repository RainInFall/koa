package koa

import (
    "net/http"
)

type Response struct{
    Status uint
    Request *Request
    res *http.ResponseWriter
}