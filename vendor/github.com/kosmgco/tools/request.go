package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var h = http.Client{}

type Client struct {
	Url     string              `json:"url,omitempty"`
	Method  HttpMethod          `json:"method,omitempty"`
	Headers map[string][]string `json:"headers,omitempty"`
	Body    interface{}         `json:"body,omitempty"`
	resp    []byte
	Params  map[string]interface{} `json:"params,omitempty"`
	Host    string                 `json:"host,omitempty"`
	Timeout time.Duration          `json:"timeout,omitempty"`
}

func (c Client) Do() Client {

	c.comments()
	var body = &bytes.Buffer{}
	var u *url.URL
	var uri string
	var err error
	//构造参数
	var p = url.Values{}

	for key, value := range c.Params {
		p.Add(key, value.(string))
	}

	//if u, err = url.Parse(fmt.Sprintf("%s%s?%s", c.Host, c.Url, p.Encode())); err != nil {
	//	logrus.Errorf("%s", err)
	//	return c
	//}
	if u, err = url.Parse(fmt.Sprintf("%s%s", c.Host, c.Url)); err != nil {
		logrus.Errorf("%s", err)
		return c
	} else {
		if v := p.Encode(); v != "" {
			uri = u.String() + "?" + p.Encode()
		} else {
			uri = u.String()
		}
	}

	if c.Body != nil {
		b, err := json.Marshal(c.Body)
		if err != nil {
			logrus.Errorf("%s", err)
			return c
		}

		body.WriteString(string(b))
	}

	if c.Timeout == 0 {
		h.Timeout = 5
	} else {
		h.Timeout = c.Timeout
	}
	req, err := http.NewRequest(c.Method.String(), uri, body)

	for key, value := range c.Headers {
		req.Header.Set(key, value[0])
	}

	resp, err := h.Do(req)
	if err != nil {
		logrus.Errorf("%s", err)
		return c
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	c.resp = b

	return c
}

func (c Client) SetURL(url string) Client {
	c.Url = url
	return c
}

func (c Client) SetMethod(method types.HttpMethod) Client {
	c.Method = method
	return c
}

func (c Client) SetHeaders(headers map[string][]string) Client {
	c.Headers = headers
	return c
}

func (c Client) SetBody(body interface{}) Client {
	c.Body = body
	return c
}

func (c Client) SetParams(params interface{}) Client {

	var m map[string]interface{}
	v, _ := json.Marshal(params)
	_ = json.Unmarshal(v, &m)
	c.Params = m
	return c
}

func (c Client) SetHost(host string) Client {
	c.Host = host
	return c
}

func (c Client) SetTimeout(timeout time.Duration) Client {
	if timeout == 0 {
		c.Timeout = 5
	} else {
		c.Timeout = timeout
	}
	return c
}

func (c Client) Into(v interface{}) error {
	return json.Unmarshal(c.resp, v)
}

func (c Client) comments() {
	logrus.Infof("request: %s%s %s params: %+v body: %+v", c.Host, c.Url, c.Method, c.Params, c.Body)
}
