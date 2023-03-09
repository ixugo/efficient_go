package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	b, err := json.Marshal(newSystemMsg("Hello"))
	require.NoError(t, err)
	fmt.Println(string(b))
}
