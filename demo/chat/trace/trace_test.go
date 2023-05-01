package trace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
	_ = tracer
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}

func TestCutSuffix(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "a.txt",
		},
		{
			desc: "b.txt",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a, b, found := strings.Cut(tC.desc, ".")
			fmt.Println(found)
			fmt.Println(a)
			fmt.Println(b)

		})
	}
}

func TestTime(t *testing.T) {
	// local := time.FixedZone("CST", 8*3600)

	const str = `{ "time" :"2023-04-01 11:04:26"}`
	var Input struct {
		Time time.Time `json:"time"`
	}
	err := json.Unmarshal([]byte(str), &Input)
	if err != nil {
		panic(err)
	}
	date := Input.Time

	fmt.Println(time.Since(date))
	fmt.Println(date)
	fmt.Println(time.Now())

}
