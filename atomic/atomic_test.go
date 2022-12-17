package atomictest

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtomic(t *testing.T) {
	var a uint32 = 0
	b := atomic.CompareAndSwapUint32(&a, atomic.LoadUint32(&a), 1)
	require.EqualValues(t, true, b)
	require.EqualValues(t, 1, a)

	b = atomic.CompareAndSwapUint32(&a, 0, 10)
	require.EqualValues(t, false, b)
	require.EqualValues(t, 1, a)

	atomic.StoreUint32(&a, 9)
	require.EqualValues(t, 9, a)
}
