package testutils

import "github.com/knoxgao67/gosec/v2"

// SampleCodeG504 - Blocklisted import CGI
var SampleCodeG504 = []CodeSample{
	{[]string{`
package main

import (
	"net/http/cgi"
	"net/http"
 )

func main() {
	cgi.Serve(http.FileServer(http.Dir("/usr/share/doc")))
}
`}, 1, gosec.NewConfig()},
}
