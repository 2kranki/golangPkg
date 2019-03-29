// See License.txt in main respository directory

// Helper Functions for use with Templates

package golangPkg

import  (
    "strings"
    "testing"
)

type TestStruct struct {
    Name        string
    Defns       map[string]string
}

type TestData struct {
    name        string
    tmpl        string
    data        interface{}
    shouldWork  bool
    result      string
}

var data        TestStruct  = TestStruct{
    "x", 
    map[string]string{ "x":"X", "y":"Y", },
}

var strTests    []TestData = []TestData {
    {"test01", "text", nil, true, "text"},
    {"test02", "{{0}}", nil, true, "0"},
    {"test03", "{{.Name}}", data, true, "x"},
    {"test04", "{{index .Defns \"x\"}}", data, true, "X"},
}

func TestTmplExecFile(t *testing.T) {
    fp := "./test.tmpl.txt"
    w := &strings.Builder{}
    err := TmplExecFile(w, fp, "not really used", nil)
    if err != nil {
        t.Errorf("File0 failed, got: '%s' needed: 'just text\\n' :: %s\n", w.String(), err)
        return
    }
    if w.String() != "just text\n" {
        t.Errorf("File0 succeded but got: '%s' needed: 'just text\\n'\n", w.String())
    }
}

func TestTmplExecStr(t *testing.T) {
    for _,test := range strTests {
        w := &strings.Builder{}
        err := TmplExecStr(w, test.name, test.tmpl, test.data, nil)
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
