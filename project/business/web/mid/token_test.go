package mid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseToken(t *testing.T) {
	token := NewToken(123, "123123")
	c, err := parseToken(token, "123123")
	require.NotNil(t, c)
	require.NoError(t, err)
}
