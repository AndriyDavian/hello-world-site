package main

import (
	"github.com/surdeus/gosrv/src/httpx"
	"log"
	"strings"
)

func UppercaseHandle(c *httpx.Context) {
	c.W.Write([]byte(strings.ToUpper(c.P)))
}

func main() {
	mux := httpx.NewMux()
	mux.Define(
		httpx.Def("/get-test/").Re("").
			Method("GET", httpx.GetTest),
		httpx.Def("/up/").Re("").
			Method("GET", UppercaseHandle),
	)

	srv := httpx.Server{
		Addr: ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}

