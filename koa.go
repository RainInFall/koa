package koa
import (
    "net/http"
)

type Middleware func(*http.Request, *http.ResponseWriter, func())

type Application struct{
    proxy bool
    middleware []Middleware
    subdomainOffset int32
    env string
    context *Context
    request *Request
    response *Response
}

func (app *Application) GetStatus() uint {
  return app.response.Status
}
