package simple_http

import (
)

type Header map[string]string

func (h Header) Get(key string) string {
	return h[key]
}

func (h Header) Set(key, value string) {
	h[key] = value
}