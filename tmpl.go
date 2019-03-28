// See License.txt in main respository directory

// Helper Functions for use with Templates

package tmpl

//import  "errors"
import  "io"
import  "io/ioutil"
import  "os"
import  "text/template"

var Tmpls *template.Template

// ExecTmplFile executes any given template while caching it for
// future use.
func ExecTmplFile(w io.Writer, path string, data interface{}, funcs template.FuncMap) error {
	var err error

    fileData,err := ioutil.ReadFile(path) 
	if err != nil {
		return err
	}
    err = ExecTmplStr(w, path, string(fileData), data, funcs)
    return err
}

// ExecTmplStr executes any given template while caching it for
// future use.
func ExecTmplStr(w io.Writer, name string, tmpl string, data interface{}, funcs template.FuncMap) error {
	var err 	error
	var tmplc	*template.Template

	if w == nil {
		w = os.Stdout
	}
	if Tmpls != nil {
		tmplc = Tmpls.Lookup(name)
	}
	if tmplc == nil {
		if Tmpls == nil {
			Tmpls,err = template.New(name).Funcs(funcs).Parse(tmpl)
		} else {
			Tmpls,err = Tmpls.New(name).Funcs(funcs).Parse(tmpl)
		}
		if err != nil {
            return err
		}
	}
	err = Tmpls.ExecuteTemplate(w, name, data)
    return err
}



