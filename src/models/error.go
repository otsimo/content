package models

import (
	"errors"
	"time"
)

var (
	ErrorNotFound       = errors.New("Not found")
	ErrorNotImplemented = errors.New("Not implemented yet")
	ErrorInvalidID      = errors.New("Invalid ID")
	ErrorIDExists       = errors.New("ID already exists")
)

func MillisecondsNow() int64 {
	s := time.Now()
	return s.Unix()*1000 + int64(s.Nanosecond()/1e6)
}
