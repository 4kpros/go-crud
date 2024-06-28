package helpers

import (
	"fmt"
	"os"
)

func ReadFileToString(path string) (contentStr string, err error) {
	content, errRead := os.ReadFile(path)
	if errRead != nil {
		fmt.Printf("\nERROR ReadFileToString ==> %s", err)
		err = errRead
		return
	}
	contentStr = string(content)
	return
}
