package models

import (
	"bytes"
	"encoding/gob"
)

// User model
type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Encode to byte array
func (u *User) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(u)
	return buf.Bytes(), err
}

// Decode from byte array
func (u *User) Decode(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(u)
	return err
}
