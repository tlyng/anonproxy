package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", "0.0.0.0:4080", "proxy listen address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			if r.URL.Scheme == "https" {
				r.URL.Scheme = "http"
			}
			return r, nil
		})

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			r.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.80 Safari/537.36")
			return r, nil
		})

	proxy.Verbose = *verbose
	log.Printf("Starting HTTP proxy on %s.\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
