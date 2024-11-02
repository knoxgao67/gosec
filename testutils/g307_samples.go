package testutils

import "github.com/knoxgao67/gosec/v2"

// SampleCodeG307 - Poor permissions for os.Create
var SampleCodeG307 = []CodeSample{
	{[]string{`
package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
		f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()
}
`}, 0, gosec.NewConfig()},
	{[]string{`
package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
		f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()
}
`}, 1, gosec.Config{"G307": "0o600"}},
}
