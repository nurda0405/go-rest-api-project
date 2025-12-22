package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

type HPPOptions struct {
	CheckBody               bool
	CheckQuery              bool
	CheckForOnlyContentType string
	Whitelist               []string
}

func HPP(hppOptions HPPOptions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if hppOptions.CheckBody && r.Method == http.MethodPost && checkContent(r, hppOptions.CheckForOnlyContentType) {
				filterBodyParams(r, hppOptions.Whitelist)
			}
			if hppOptions.CheckQuery && r.URL.Query() != nil {
				filterQueryParams(r, hppOptions.Whitelist)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func checkContent(r *http.Request, checkContent string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), checkContent)
}

func filterBodyParams(r *http.Request, whitelist []string) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[len(v)-1])
		}
		if !isWhitelisted(k, whitelist) {
			delete(r.Form, k)
		}
	}
}

func filterQueryParams(r *http.Request, whitelist []string) {
	query := r.URL.Query()

	for k, v := range query {
		if len(v) > 1 {
			query.Set(k, v[len(v)-1])
		}
		if !isWhitelisted(k, whitelist) {
			query.Del(k)
		}
	}
	r.URL.RawQuery = query.Encode()
}

func isWhitelisted(param string, whitelist []string) bool {
	for _, v := range whitelist {
		if v == param {
			return true
		}
	}
	return false
}
