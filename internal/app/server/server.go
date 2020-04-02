package server

import (
	"net/http"
)

func NewServer() error {
	NewHandle()
	err := http.ListenAndServe(":65000", nil)
	return err
}
