package main

import (
	"net/http"
	"time"

	"github.com/rs/cors"
)

var client = &http.Client{Timeout: time.Second * 30}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取真实 host
		host := r.Header.Get("X-Real-Host")
		if host == "" {
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("{\"message\": \"your request must have a header X-Real-Host\"}"))
		}
		// 移除 host Header
		r.Header.Del("X-Real-Host")
		// 替换请求
		r.Host = host
		r.Header.Set("Host", host)
		r.URL.Host = host
		// 发送请求 FIXME: 这里不对，不能直接发，要重造一个
		resp, err := client.Do(r)
	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
