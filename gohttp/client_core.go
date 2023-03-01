package gohttp

import (
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{})
