// See License.txt in main respository directory

// Helper Functions for use with Templates

package golangPkg

//import  "errors"
import  "io"
import  "io/ioutil"
import  "os"
import  "text/template"

var ttmpls *template.Template

// TmplExecFile executes any given template while caching it for
// future use.
func TmplExecFile(w io.Writer, path string, data interface{}, funcs template.FuncMap) error {
	var err error

    fileData,err := ioutil.ReadFile(path) 
	if err != nil {
		return err
	}
    err = TmplExecStr(w, path, string(fileData), data, funcs)
    return err
}

// TmplExecStr executes any given template while caching it for
// future use.
func TmplExecStr(w io.Writer, name string, tmpl string, data interface{}, funcs template.FuncMap) error {
	var err 	error
	var tmplc	*template.Template

	if w == nil {
		w = os.Stdout
	}
	if ttmpls != nil {
		tmplc = ttmpls.Lookup(name)
	}
	if tmplc == nil {
		if ttmpls == nil {
			ttmpls,err = template.New(name).Funcs(funcs).Parse(tmpl)
		} else {
			ttmpls,err = ttmpls.New(name).Funcs(funcs).Parse(tmpl)
		}
		if err != nil {
            return err
		}
	}
	err = ttmpls.ExecuteTemplate(w, name, data)
    return err
}



