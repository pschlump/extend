package template

import (
	"io"
	"os"
)

// AvailableTemplates returns a array string listing the defined templates by name,
// If there are none, it returns an empty array.  (PJS-Wed Jun 15 12:47:01 MDT 2016)
// PJS - Ported to go v1.12 - Wed Mar 20 07:04:09 MDT 2019
func (t *Template) AvailableTemplates() (rv []string) {
	if t.common == nil {
		return
	}
	for name, tmpl := range t.tmpl {
		if tmpl.Tree == nil || tmpl.Root == nil {
			continue
		}
		rv = append(rv, name)
	}
	return
}

// PJS
var no_value = "<no value>"
var reportMissing = true
var logOutput io.WriteCloser = os.Stderr

// SetNoValue determines how to handle missing values.
// PJS
func SetNoValue(s string, r bool) {
	no_value = s
	reportMissing = r
}

// SetLogOutput - where to output for errors.
// PJS
func SetLogOutput(f io.WriteCloser) {
	logOutput = f
}
