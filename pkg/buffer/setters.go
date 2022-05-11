package buffer

import (
	"fmt"
)

// WriteInt64 writes an int64 to the buffer.
func (b *Buf) WriteInt64(i int64) error {
	// create 8 byte buffer
	buf := make([]byte, 8)

	// write int64 to buffer
	buf[0] = byte(i)
	buf[1] = byte(i >> 8)
	buf[2] = byte(i >> 16)
	buf[3] = byte(i >> 24)
	buf[4] = byte(i >> 32)
	buf[5] = byte(i >> 40)
	buf[6] = byte(i >> 48)
	buf[7] = byte(i >> 56)

	// write buffer to b
	return b.WriteBytes(buf)
}

// WriteUint64 writes a uint64 to the buffer.
func (b *Buf) WriteUint64(i uint64) error {
	// create 8 byte buffer
	buf := make([]byte, 8)

	// write uint64 to buffer
	buf[0] = byte(i)
	buf[1] = byte(i >> 8)
	buf[2] = byte(i >> 16)
	buf[3] = byte(i >> 24)
	buf[4] = byte(i >> 32)
	buf[5] = byte(i >> 40)
	buf[6] = byte(i >> 48)
	buf[7] = byte(i >> 56)

	// write buffer to b
	return b.WriteFromSlice(buf, 8)
}

// WriteInt32 writes an int32 to the buffer.
func (b *Buf) WriteInt32(i int32) error {
	// create 4 byte buffer
	buf := make([]byte, 4)

	// write int32 to buffer
	buf[0] = byte(i)
	buf[1] = byte(i >> 8)
	buf[2] = byte(i >> 16)
	buf[3] = byte(i >> 24)

	// write buffer to b
	return b.WriteFromSlice(buf, 4)
}

// WriteUint32 writes a uint32 to the buffer.
func (b *Buf) WriteUint32(i uint32) error {
	// create 4 byte buffer
	buf := make([]byte, 4)

	// write uint32 to buffer
	buf[0] = byte(i)
	buf[1] = byte(i >> 8)
	buf[2] = byte(i >> 16)
	buf[3] = byte(i >> 24)

	// write buffer to b
	return b.WriteFromSlice(buf, 4)
}

// WriteInt16 writes an int16 to the buffer
func (b *Buf) WriteInt16(i int16) error {
	// create 2 byte buffer
	buf := make([]byte, 2)

	// write int16 to buffer
	buf[0] = byte(i)
	buf[1] = byte(i >> 8)

	// write buffer to b
	return b.WriteFromSlice(buf, 2)
}

// WriteUint16 writes a uint16 to the buffer.
func (b *Buf) WriteUint16(i uint16) error {
	// create 2 byte buffer
	buf := make([]byte, 2)

	// write uint16 to buffer
	buf[0] = byte(i)
	buf[1] = byte(i >> 8)

	// write buffer to b
	return b.WriteFromSlice(buf, 2)
}

// WriteInt8 writes an int8 to the buffer.
func (b *Buf) WriteInt8(i int8) error {
	// create 1 byte buffer
	buf := make([]byte, 1)

	// write int8 to buffer
	buf[0] = byte(i)

	// write buffer to b
	return b.WriteFromSlice(buf, 1)
}

// WriteUint8 writes a uint8 to the buffer.
func (b *Buf) WriteUint8(i uint8) error {
	// create 1 byte buffer
	buf := make([]byte, 1)

	// write uint8 to buffer
	buf[0] = byte(i)

	// write buffer to b
	return b.WriteFromSlice(buf, 1)
}

// WriteString writes a string to the buffer.
func (b *Buf) WriteString(s string) error {
	// create buffer
	buf := []byte(s)

	// write buffer to b
	return b.WriteFromSlice(buf, len(buf))
}

// WriteBytes writes a byte slice to the buffer.
func (b *Buf) WriteBytes(buf []byte) error {
	// write buffer to b
	return b.WriteFromSlice(buf, len(buf))
}

// WriteBlob writes a blob to the buffer by setting the lenght information and then writing the blob after it.
func (b *Buf) WriteBlob(blob []byte) error {
	// determine how many bytes are required to represent the length of blob
	lengthb := len(blob)
	length := lengthb

	// bit shift right
	length = length >> 8

	// if length is 0, uint8
	if length == 0 {
		// write uint8
		if err := b.WriteByte(byte(Uint8)); err != nil {
			return err
		}
		if err := b.WriteUint8(uint8(lengthb)); err != nil {
			return err
		}
	} else {
		// bit shift right
		length = length >> 8

		// if length is 0, uint16
		if length == 0 {
			// write uint16
			if err := b.WriteByte(byte(Uint16)); err != nil {
				return err
			}
			if err := b.WriteUint16(uint16(lengthb)); err != nil {
				return err
			}
		} else {
			// bit shift right
			length = length >> 16

			// if length is 0, uint32
			if length == 0 {
				// write uint32
				if err := b.WriteByte(byte(Uint32)); err != nil {
					return err
				}
				if err := b.WriteUint32(uint32(lengthb)); err != nil {
					return err
				}
			} else {
				// bit shift right
				length = length >> 32

				// if length is 0, uint64
				if length == 0 {
					// write uint64
					if err := b.WriteByte(byte(Uint64)); err != nil {
						return err
					}
					if err := b.WriteUint64(uint64(lengthb)); err != nil {
						return err
					}
				} else {
					// unsupported number
					return fmt.Errorf("unsupported number")
				}
			}
		}
	}

	// write blob after length
	return b.WriteBytes(blob)
}
