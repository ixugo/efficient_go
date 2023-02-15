package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ixugo/efficient_go/demo/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

type Runtime int32

var _ json.Unmarshaler = new(Runtime)
var _ json.Marshaler = new(Runtime)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r *Runtime) UnmarshalJSON(b []byte) error {
	v := string(b)
	v = strings.Trim(v, `"`)
	v = strings.TrimSpace(strings.TrimRight(v, `mins`))
	i, err := strconv.Atoi(v)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*r = Runtime(i)
	return nil
}

func (r Runtime) MarshalJSON() ([]byte, error) {
	v := strconv.Quote(fmt.Sprintf("%d mins", r))
	return []byte(v), nil
}

func ValidateMovie(v *validator.Validator, m *Movie) {
	v.Check(m.Title != "", "title", "must be provided")
	v.Check(len(m.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(m.Year != 0, "year", "must be provided")
}
