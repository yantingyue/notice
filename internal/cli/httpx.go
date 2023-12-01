package cli

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

const (
	timeOut           = 60 * time.Second
	GET_METHOD        = "GET"
	POST_METHOD       = "POST"
	JSON_CONTENT_TYPE = "application/json"
)

var fastHttpClient fasthttp.Client

func Get(host string) (m map[string]interface{}, err error) {
	statusCode, body, err := fasthttp.GetTimeout(nil, host, timeOut)
	if err != nil {
		return
	}
	if statusCode != fasthttp.StatusOK {
		err = errors.New(fmt.Sprintf("request failed statusCode[%d]", statusCode))
		return
	}
	if body == nil {
		err = errors.New("response body is nil")
		return
	}
	m = make(map[string]interface{})
	if err = json.Unmarshal(body, &m); err != nil {
		return
	}
	return
}

func Post(host string, header map[string]string, payload []byte) (body []byte, err error) {
	req := &fasthttp.Request{}
	req.SetRequestURI(host)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	req.SetBody(payload)
	resp := &fasthttp.Response{}
	client := &fasthttp.Client{}
	if err = client.DoTimeout(req, resp, timeOut); err != nil {
		return
	}
	body = resp.Body()
	return
}

func FastPostJson(url string, params map[string]interface{}, body []byte) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	// content-type default is application/x-www-form-urlencoded
	req.Header.SetContentType(JSON_CONTENT_TYPE)
	req.Header.SetMethod(POST_METHOD)
	//req.Header.SetHost("api.oneniceapp.com")
	req.SetRequestURI(setUrl(url, params))
	req.SetBody(body)
	return req
}

func setUrl(url string, params map[string]interface{}) string {
	values := make([]string, 0)
	for k, v := range params {
		values = append(values, fmt.Sprintf("%s=%v", k, v))
	}
	if len(values) > 0 {
		url = url + "?" + strings.Join(values, "&")
	}
	return url
}

func FastResponse(req *fasthttp.Request) ([]byte, error) {
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)
	if err := fastHttpClient.Do(req, resp); err != nil {
		return []byte{}, err
	}
	return resp.Body(), nil
}
