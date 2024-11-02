package testutils

import "github.com/knoxgao67/gosec/v2"

// SampleCodeG106 - ssh InsecureIgnoreHostKey
var SampleCodeG106 = []CodeSample{
	{[]string{`
package main

import (
	"golang.org/x/crypto/ssh"
)

func main() {
		_ =  ssh.InsecureIgnoreHostKey()
}
`}, 1, gosec.NewConfig()},
}
