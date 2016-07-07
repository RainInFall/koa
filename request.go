package koa

import (
    "net/http"
    URL "net/url"
    "strings"
    "errors"
    "sort"
)

type Request struct {
    Response *Response
    Req *http.Request
    App *Application
}

func (req *Request) GetHeader() http.Header {
    return req.Req.Header
}

func (req *Request) GetHeaders() http.Header {
    return req.Req.Header
}

func (req *Request) GetUrl() string{
    return req.Req.URL.RequestURI()
}

func (req *Request) SetUrl(url string) error{
    var err error
    if newUrl, err := URL.Parse(url); nil == err {
        req.Req.URL = req.Req.URL.ResolveReference(newUrl)
    }
    return err
}

func (req *Request) GetOrigin() string {
  return req.Req.URL.Scheme+"://"+req.Req.URL.Host
}

func (req *Request) GetHref() string {
  return req.Req.URL.String()
}

func (req *Request) GetMethod() string {
  return req.Req.Method
}

func (req *Request) SetMethod(method string) {
  req.Req.Method = method
}

func (req *Request) GetPath() string {
  return req.Req.URL.Path
}

func (req *Request) GetQuery() URL.Values {
  return req.Req.URL.Query()
}

func (req *Request) SetQUery(query URL.Values) {
    req.Req.URL.RawQuery = query.Encode()
}

func (req *Request) GetQueryString() string {
  return req.Req.URL.RequestURI()
}

func (req *Request) SetQueryString(query string) {
  req.Req.URL.RawQuery = query
}

func (req *Request) GetSearch() string {
  if len(req.Req.URL.RawQuery) == 0 {
    return ""
  }
  return "?"+req.Req.URL.RawQuery
}

func (req *Request) SetSearch(search string) {
  req.Req.URL.RawQuery = search
}

func (req *Request) GetHost() string {
  var proxy = req.App.proxy
  var host string
  var err error
  if proxy {
    host, err = req.Get("X-Forwarded-Host")
  }
  if nil != err {
    host = req.Req.Host
  }
  host = strings.Split(host, ",")[0]
  return strings.TrimSpace(host)
}

func (req *Request) GetHostname() string {
  return strings.Split(req.GetHost(), ":")[0]
}

func (req *Request) GetFresh() bool {
  if method := req.GetMethod();
    strings.Compare("GET", method) != 0 &&
    strings.Compare("HEAD", method) != 0 {
    return false
  }
  if status := req.App.GetStatus();
    (status >= 200 && status < 300) || 304 == status {
    return false
    //TODO:move fresh from npm to GO
  }
  return false;
}

func (req *Request) GetStale() bool {
  return !req.GetFresh()
}

func (req *Request) IsIdempotent() bool {
  methods := []string{"DELETE", "GET", "HEAD", "OPTIONS", "PUT", "TRACE"}
  method := req.GetMethod()
  index := sort.SearchStrings( methods, method)
  return 0 == strings.Compare(method, methods[index])
}

func (req *Request) GetCharset() string {
  //TODO:move contentType fomr npm to GO
  return ""
}

func (req *Request) Get(key string) (string, error) {
  if pair := req.Req.Header[key]; len(pair) == 0 {
    return "", errors.New("nil")
  } else {
    return pair[0], nil
  }
}
