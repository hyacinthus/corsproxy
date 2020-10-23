package main

import (
	"fmt"
	"net/http"

	"github.com/hack-fan/config"
	"github.com/hyacinthus/corsproxy"
)

func main() {
	// load config
	var cfg = new(corsproxy.Config)
	config.MustLoad(cfg)
	// proxy
	var proxy = corsproxy.NewProxy(cfg, nil)
	http.HandleFunc("/", proxy.ServeHTTP)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}
