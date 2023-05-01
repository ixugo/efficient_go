package request

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := http.Get("http://asd.gr.com/123/124/123")
	fmt.Printf("vv: %#v\n", err)

	if err, ok := err.(*url.Error); ok {
		fmt.Println(ok)
		fmt.Println(err.Err)
	}
	err = errors.Unwrap(err)
	fmt.Printf("v: %+v\n", err)

	fmt.Println(errors.Unwrap(nil))
	_ = resp

}
