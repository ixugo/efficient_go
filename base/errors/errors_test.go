package errors_test

import (
	"errors"
	"fmt"
	"testing"
)

func call() error {
	return errors.New("test")
}

func action() error {
	if err := call(); err != nil {
		return fmt.Errorf("action() -> call() err(%w)", err)
	}
	return nil
}

func TestWrap(t *testing.T) {
	err := action()
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
		fmt.Printf("%+v\n", err)
		fmt.Printf("%#v\n", err)
	}
}
