// See License.txt in main respository directory

// Helper Functions for use with Templates.
// I found discussion that once a template in the cache is
// executed, it locks the cache and no more templates may
// be added.  This only seems to affect the html/template
// mechanism and not the text/template side.

package golangPkg

import (
	"errors"
	"fmt"
	"html/template"
)
import  "io"
import  "io/ioutil"
import  "os"

var htmpls *template.Template

// HtmlExecStr executes any given template while caching it for
// future use.
func HtmlExecStr(w io.Writer, name string, data interface{}) error {
	var err 	error
	var tmplc	*template.Template

	if w == nil {
		w = os.Stdout
	}
	if htmpls != nil {
		tmplc = htmpls.Lookup(name)
	}
	if tmplc == nil || htmpls == nil{
		return errors.New(fmt.Sprint(name,"was not previously cached"))
	}
	err = tmplc.ExecuteTemplate(w, name, data)
    return err
}

// HtmlLoadFile loads a given template to the cache for
// future use.
func HtmlLoadFile(path string, funcs template.FuncMap) error {
	var err error

    fileData,err := ioutil.ReadFile(path) 
	if err != nil {
		return err
	}
    err = HtmlLoadStr(path, string(fileData), funcs)
    return err
}

// HtmlLoadStr loads a given template into the cache.
func HtmlLoadStr(name string, tmpl string, funcs template.FuncMap) error {
	var err 	error
	var tmplc	*template.Template

	if htmpls != nil {
		tmplc = htmpls.Lookup(name)
	}
	if tmplc == nil {
		if htmpls == nil {
			htmpls,err = template.New(name).Funcs(funcs).Parse(tmpl)
		} else {
			htmpls,err = htmpls.New(name).Funcs(funcs).Parse(tmpl)
		}
	}
    return err
}

// HtmlCacheReset resets the cache.
func HtmlCacheReset() {
    htmpls = nil
}


