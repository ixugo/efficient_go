package logger

import "testing"

func TestPrint(t *testing.T) {
	log, _ := InitJSONLogger("./")
	log.Infow("a", "b", "c")

}
