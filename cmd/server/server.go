package main

import (
	"github.com/hack-fan/config"
	"github.com/hyacinthus/corsproxy"
)

func main() {
	// load config
	var cfg = new(corsproxy.Config)
	config.MustLoad(cfg)
	// proxy
	var proxy = corsproxy.NewProxy(cfg, nil)
	proxy.Start()
}
