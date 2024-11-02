package main

import "text/template"

var generatedHeaderTmpl = template.Must(template.New("generated").Parse(`
package {{.}}

import (
	"go/ast"

	"github.com/knoxgao67/gosec/v2"
	"github.com/knoxgao67/gosec/v2/issue"
)
`))
