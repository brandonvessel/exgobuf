package buffer

import "fmt"

type Buf struct {
	// buf is the data contained in the buffer
	buf *[]byte

	// ptr is the current read and write position in the buffer
	ptr int

	// max is the maximum number of bytes the buffer can hold
	max int `default:"0"` // 0 means no limit
}

// NewBuf creates a new Buf with unlimited size.
func NewBuf() *Buf {
	return &Buf{buf: &[]byte{}}
}

// NewBufMax creates a new Buf with a maximum size of max.
func NewBufMax(max int) *Buf {
	// if less than 0, set to 0
	if max < 0 {
		max = 0
	}
	return &Buf{buf: &[]byte{}, max: max}
}

// NewBufFromBytes creates a new Buf from a byte slice with unlimited size.
func NewBufFromBytes(buf *[]byte) *Buf {
	return &Buf{buf: buf, max: 0}
}

// NewBufMax creates a new Buf with a maximum size of max.
func NewBufMaxFromBytes(buf *[]byte, max int) *Buf {
	return &Buf{buf: buf, max: max}
}

// ResetPtr resets the read and write position to the beginning of the buffer.
func (b *Buf) ResetPtr() {
	b.ptr = 0
}

// ReadByte reads a single byte from the buffer.
func (b *Buf) ReadByte() (byte, error) {
	// Check if we have reached the end of the buffer
	if b.ptr >= len(*b.buf) {
		return 0, ErrAdvanceTooFar
	}

	// Read the byte
	c := (*b.buf)[b.ptr]

	// Advance the pointer
	err := b.IncrementPtr()
	if err != nil {
		return 0, err
	}

	// Return the byte
	return c, nil
}

// WriteByte writes a single byte to the buffer.
func (b *Buf) WriteByte(c byte) error {
	// Check if we have reached the max size of the buffer
	if b.max > 0 && b.ptr >= b.max {
		return ErrBufferFull
	}

	// if we have reached size, expand the buffer
	if b.ptr >= len(*b.buf) {
		*b.buf = append(*b.buf, 0)
	}

	// Write the byte
	(*b.buf)[b.ptr] = c

	// Advance the pointer
	err := b.IncrementPtr()
	if err != nil {
		return err
	}

	// Return nil
	return nil
}

// Read reads up to len(p) bytes from the buffer into p.
func (b *Buf) ReadFromSlice(p []byte, i int) error {
	// check if i is 0
	if i == 0 {
		return ErrInvalidLength
	}

	// check if i is greater than the length of p
	if i > len(p) {
		return ErrInvalidLength
	}

	// iterate i times
	for j := 0; j < i; j++ {
		// read single byte
		c, err := b.ReadByte()

		// check if we have an error
		if err != nil {
			return err
		}

		// write byte to p
		p[j] = c
	}

	// return nil
	return nil
}

// Write writes i bytes from p to the buffer.
func (b *Buf) WriteFromSlice(p []byte, i int) error {
	// iterate i times
	for j := 0; j < i; j++ {
		// write single byte
		err := b.WriteByte(p[j])

		// check if we have an error
		if err != nil {
			return err
		}
	}

	// return nil
	return nil
}

// Len returns the number of written bytes in the buffer.
func (b *Buf) Len() int {
	return len(*b.buf)
}

// SetPtr sets the pointer to the given position.
func (b *Buf) SetPtr(i int) error {
	// check if i is greater than the buffer length
	if i > len(*b.buf) {
		return ErrAdvanceTooFar
	}

	// set the pointer
	b.ptr = i

	// return nil
	return nil
}

// GetPtr returns the current pointer position.
func (b *Buf) GetPtr() int {
	return b.ptr
}

// SetMax sets the maximum number of bytes the buffer can hold.
func (b *Buf) SetMax(i int) {
	b.max = i
}

// GetMax returns the maximum number of bytes the buffer can hold.
func (b *Buf) GetMax() int {
	return b.max
}

// DecrementPtr decrements the pointer by 1.
func (b *Buf) DecrementPtr() error {
	// check if we are at the beginning of the buffer
	if b.ptr == 0 {
		return ErrAdvanceTooFar
	}
	b.ptr--
	return nil
}

// IncrementPtr increments the pointer by 1.
func (b *Buf) IncrementPtr() error {
	// check if we are at the end of the buffer
	if b.ptr >= len(*b.buf) {
		return ErrAdvanceTooFar
	}
	b.ptr++
	return nil
}

// String returns the string representation of the buffer.
func (b *Buf) String() string {
	return fmt.Sprintf("%v", *b.buf)
}
