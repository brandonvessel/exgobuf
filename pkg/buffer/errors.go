package buffer

import "errors"

var (
	// ErrAdvanceTooFar is returned when the pointer is advanced beyond the end of the buffer
	ErrAdvanceTooFar = errors.New("advance beyond end of buffer")
	// ErrBufferFull is returned when the buffer is full
	ErrBufferFull = errors.New("buffer full")
	// InvalidLength is returned when the byte read or write amount is invalid
	ErrInvalidLength = errors.New("invalid length")
)
