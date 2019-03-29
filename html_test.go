// See License.txt in main respository directory

// Helper Functions for use with Templates

package golangPkg

import  (
    "strings"
    "testing"
)

type HtmlTestStruct struct {
    Name        string
    Defns       map[string]string
}

type HtmlTestData struct {
    name        string
    tmpl        string
    data        interface{}
    shouldWork  bool
    result      string
}

var dataHtml        HtmlTestStruct  = HtmlTestStruct{
    "x", 
    map[string]string{ "x":"X", "y":"Y", },
}

var strTestsHtml    []HtmlTestData = []HtmlTestData {
    {"htest01", "<a {{.Name}}>", dataHtml, true, "<a x>"},
    {"htest02", "{{0}}", nil, true, "0"},
    {"htest03", "{{.Name}}", dataHtml, true, "x"},
    {"htest04", "{{index .Defns \"x\"}}", dataHtml, true, "X"},
}

var strTestFile = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>functions</title>
</head>
<body>
</body>
</html>
`

func TestHtmlExecFile(t *testing.T) {
    fp := "./test.html.txt"
    w := &strings.Builder{}
    HtmlCacheReset()
    err := HtmlLoadFile(fp, nil)
    if err != nil {
        t.Errorf("File0 failed, got: '%s' needed: 'just text\\n' :: %s\n", w.String(), err)
        return
    }
    err = HtmlExecStr(w, fp, "not really used")
    if err != nil {
        t.Errorf("File0 failed, got: '%s' needed: 'just text\\n' :: %s\n", w.String(), err)
        return
    }
    if w.String() != strTestFile {
        t.Errorf("File0 succeded but got: '%s' needed: '%s'\n", w.String(), strTestFile)
    }
}

func TestHtmlExecStr(t *testing.T) {

    HtmlCacheReset()

    // Load the templates into the cache
    for _,test := range strTestsHtml {
        err := HtmlLoadStr(test.name, test.tmpl, nil)
        if err != nil {
            t.Errorf("%s failed load into cache :: %s\n", test.name, err)
        }
    }

    // Execute the templates
    for _,test := range strTestsHtml {
        w := &strings.Builder{}
        err := HtmlExecStr(w, test.name, test.data)
        if err != nil {
            if test.shouldWork {
                t.Errorf("%s failed, got: '%s' needed: '%s' :: %s\n", 
                        test.name, w.String(), test.result,err)
                continue
            }
        }
        if test.shouldWork {
            if w.String() != test.result {
                t.Errorf("%s succeded but got: '%s' needed: '%s'\n", test.name, w.String(), test.result)
            }
        } else {
            t.Errorf("%s succeded but got: '%s' needed: '%s'\n", test.name, w.String(), test.result)
        }

    }
}
