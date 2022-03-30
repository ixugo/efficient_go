package reflect_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestTag 获取 tag
func TestTag(t *testing.T) {
	u := user{Name: "xx", Age: 21}

	e := reflect.ValueOf(&u).Elem()

	for i := 0; i < e.NumField(); i++ {
		v := e.Type().Field(i)
		tag := v.Tag
		fmt.Println(tag.Get("json"))
	}
}

// TestType 通过反射获取类型
func TestType(t *testing.T) {
	a := 2
	r := reflect.TypeOf(a)
	switch r.Kind() {
	case reflect.String:
		t.Log("str")
	case reflect.Array:
		t.Log("array")
	case reflect.Int, reflect.Int32, reflect.Int64:
		t.Log("int")
	default:
		t.Log("unknow")
	}
}

// TestDeepEqual 比较 map 和 slice
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "a", 2: "b", 3: "c"}
	b := map[int]string{1: "a", 2: "b", 3: "c"}
	c := map[int]string{1: "a", 2: "b", 3: "d"}
	require.EqualValues(t, reflect.DeepEqual(a, b), true)
	require.EqualValues(t, reflect.DeepEqual(a, c), false)

	d := []int{1, 2, 3}
	f := []int{1, 2, 3}
	require.EqualValues(t, reflect.DeepEqual(d, f), true)
}
