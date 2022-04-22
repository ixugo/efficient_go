package interface_test

import (
	"testing"
)

type customError struct{}

func (c *customError) Error() string {
	return "t"
}

func fail() *customError {
	return nil
}
func TestNilInterface(t *testing.T) {
	var err error
	// 此时的 error 接口，有了类型。
	// 所以不为 nil
	if err = fail(); err != nil {
		t.Log(err)
	}

}
