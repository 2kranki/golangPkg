// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Program to play with templates and see what is
// available for us to use in this context.

package golangPkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/2kranki/jsonpreprocess"
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

// ReadJson preprocesses out comments and then unmarshals the data
// generically.
func ReadJson(jsonPath string) (interface{}, error) {
	var err			error
	var jsonOut		interface{}

	// Open the input template file
	input, err := os.Open(jsonPath)
	if err != nil {
		return jsonOut,err
	}
	textBuf := strings.Builder{}
	err = jsonpreprocess.WriteMinifiedTo(&textBuf, input)
	if err != nil {
		return jsonOut,err
	}

	// Read and process the template file
	err = json.Unmarshal([]byte(textBuf.String()), &jsonOut)
	if err != nil {
		return jsonOut,err
	}

	return jsonOut,err
}

// ReadJson preprocesses out comments and then unmarshals the data
// into a data structure previously defined.
func ReadJsonToData(jsonPath string, jsonOut interface{}) error {
	var err			error

	// Open the input template file
	input, err := os.Open(jsonPath)
	if err != nil {
		return err
	}
	textBuf := strings.Builder{}
	err = jsonpreprocess.WriteMinifiedTo(&textBuf, input)
	if err != nil {
		return err
	}

	// Read and process the template file
	err = json.Unmarshal([]byte(textBuf.String()), jsonOut)
	if err != nil {
		return err
	}

	return err
}

