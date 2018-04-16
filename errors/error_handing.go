package errors

import (
	"fmt"
	"os"
)

// HandleIfError takes care of an error if not nil.
// The intent here is to centralize error handling and DRY up the code a bit.
func HandleIfError(errorToHandle error) {
	if errorToHandle == nil {
		return
	}
	fmt.Println(errorToHandle.Error())
	os.Exit(1)
}
