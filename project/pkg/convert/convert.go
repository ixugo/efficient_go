package convert

import "strconv"

// StrTo 字符串转换整型
type StrTo string

func (s StrTo) String() string {
	return string(s)
}

// Int ...
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

// MustInt ...
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

// UInt32 ...
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

// MustUInt32 ...
func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}
