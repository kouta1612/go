package main

import (
	"errors"
	"fmt"
)

type WrappedError struct {
	inner error
}

func main() {
	inner := errors.New("this is inner error")
	wrapper := WrappedError{inner}
	fmt.Println(errors.Unwrap(&wrapper))
}

func (e *WrappedError) Error() string {
	return fmt.Sprintf("this is wrapper error: %v", e.inner.Error())
}

func (e *WrappedError) Unwrap() error {
	return e.inner
}
