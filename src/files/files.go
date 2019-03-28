// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Program to play with templates and see what is
// available for us to use in this context.

package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)



// IsPathRegularFile cleans up the supplied file path
// and then checks the cleaned file path to see
// if it is an existing standard file. Return the
// cleaned up path and a potential error if it exists.
func IsPathRegularFile(fp string) (string,error) {
	var	err 	error
	var path	string

	fp = filepath.Clean(fp)
	path,err = filepath.Abs(fp)
	if err != nil {
		return path,errors.New(fmt.Sprint("Error getting absolute path for:", fp, err))
	}
	fi, err := os.Lstat(path)
	if err != nil {
		return path,errors.New("path not found")
	}
	if fi.Mode().IsRegular() {
		return path,nil
	}
	return path,errors.New("path not regular file")
}



