package system

import "io/ioutil"

// IOUtiler is interface of io util classes
type IOUtiler interface {
	ReadFile(fileName string) ([]byte, error)
}

// IOUtil is a wrapper of io/ioutil
type IOUtil struct{}

// ReadFile opens file and reads all contents in that file.
func (u IOUtil) ReadFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}
