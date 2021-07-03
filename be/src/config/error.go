package config

import (
	"net/http"
)

type HttpErr struct {
	Code int
}

func (s *HttpErr) Error() string {
	return http.StatusText(s.Code)
}
