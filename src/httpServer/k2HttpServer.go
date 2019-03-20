// Helper Functions for use in HTTP Servers
package k2HttpServer

import "errors"
import "fmt"
import "html/template"
import "log"
import "net/http"
import "os"

var Templates *template.Template

func isPathRegularFile(fn string) bool {
	fi, err := os.Lstat(fn)
	if err != nil {
		log.Fatalln("Error while stating file:", err)
	}
	if fi.Mode().IsRegular() {
		return true
	}
	return false
}

// WriteTemplate executes any given template while caching it for
// future use.
func WriteTemplate(w http.ResponseWriter, name string, data interface{}) {
	var err = errors.New("")
	path := name + ".gohtml"
	tmpl := Templates.Lookup(path)
	if tmpl == nil {
		if !isPathRegularFile(path) {
			http.Error(w, "Error can not stat: "+path, http.StatusInternalServerError)
			return
		}
		if Templates == nil {
			tmpl, err = template.ParseFiles(path)
		} else {
			tmpl, err = Templates.ParseFiles(path)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	err = tmpl.ExecuteTemplate(w, path, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


// WriteHttpResponseHtml should properly write out HTML bytes to the
// ResponseWriter.  I have looked at the source for http.ResponseWriter,
// but I do not fully understand it yet.  It is probable that this
// function is not needed.
func WriteHttpResponseHtml(w http.ResponseWriter, status int, data []byte) {
	var dataLen int
	if data != nil {
		dataLen = len(data)
	}
	//w.WriteHeader(status)           // Response Writer does this automatically
	c := fmt.Sprint("Content-Length: %d\r\n", dataLen)
	w.Write([]byte(c))
	w.Write([]byte("Content-Type: text/html\r\n\r\n"))
	if dataLen > 0 {
		w.Write(data)
	}
}
