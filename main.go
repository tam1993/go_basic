package main

// fmt is a package that provides I/O functions

import (
	"errors"
	"study/mypackage"
	// "github.com/google/uuid"
)

func main() {
	/* for */
	// mypackage.ForLoop()

	/* array */
	// mypackage.Array()

	/* struct */
	// mypackage.Struct()
	// mypackage.ForArrayStruct()

	/* function */
	// mypackage.FunctionStruct()

	/* Pointer */
	// mypackage.Pointer()
	mypackage.PointerStruct()

	/* interface */
	// mypackage.Interface()

}

// for test error
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can not divide zero")
	}
	return a / b, nil
}
