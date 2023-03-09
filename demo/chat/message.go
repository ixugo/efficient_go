package main

import (
	"encoding/json"
	"time"
)

type message struct {
	Name    string    `json:"name"`
	Message string    `json:"message"`
	When    time.Time `json:"when"`
}

func newSystemMsg(msg string) *message {
	return &message{
		Name:    "系统",
		Message: msg,
		When:    time.Now(),
	}
}

// MarshalJSON implements json.Marshaler
func (m message) MarshalJSON() ([]byte, error) {
	type tmp message
	msg := struct {
		tmp
		When string `json:"when"`
	}{
		tmp:  tmp(m),
		When: m.When.Format(time.TimeOnly),
	}
	return json.Marshal(msg)
}

var _ json.Marshaler = new(message)
