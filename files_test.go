// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test files package

package golangPkg

import "fmt"
import "testing"

func TestIsPathRegularFile(t *testing.T) {
    var path        string
    var err         error

    path,err = IsPathRegularFile("./files.go")
    if err != nil {
		t.Errorf("IsPathRegularFile(./files.go) failed: %s\n", err)
    }
    fmt.Println("./files.go path:",path)

    path,err = IsPathRegularFile("./xyzzy.go")
    if err == nil {
		t.Errorf("IsPathRegularFile(./xyzzy.go) should have failed!\n")
    }
    fmt.Println("./xyzzy.go path:",path)

}

