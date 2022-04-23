package define

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
 * base face
 */

type Base struct {
}

//construct
func NewBase() *Base {
	this := &Base{}
	return this
}

//read origin file
func (f *Base) ReadFile(filePath string) ([]byte, error)  {
	curDir, _ := os.Getwd()
	fileRealPath := fmt.Sprintf("%v/%v", curDir, filePath)
	fileData, err := ioutil.ReadFile(fileRealPath)
	return fileData, err
}
