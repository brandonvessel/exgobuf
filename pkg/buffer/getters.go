package buffer

import (
	"fmt"
)

// ReadInt64 reads an int64 from the buffer.
func (b *Buf) ReadInt64() (int64, error) {
	var buf [8]byte
	var err error

	for i := 0; i < 8; i++ {
		buf[i], err = b.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// return int64
	return int64(buf[0]) | int64(buf[1])<<8 | int64(buf[2])<<16 | int64(buf[3])<<24 | int64(buf[4])<<32 | int64(buf[5])<<40 | int64(buf[6])<<48 | int64(buf[7])<<56, nil
}

// ReadUint64 reads a uint64 from the buffer.
func (b *Buf) ReadUint64() (uint64, error) {
	var buf [8]byte
	var err error

	for i := 0; i < 8; i++ {
		buf[i], err = b.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// return uint64
	return uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 | uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56, nil
}

// ReadInt32 reads an int32 from the buffer.
func (b *Buf) ReadInt32() (int32, error) {
	var buf [4]byte
	var err error

	for i := 0; i < 4; i++ {
		buf[i], err = b.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// return int32
	return int32(buf[0]) | int32(buf[1])<<8 | int32(buf[2])<<16 | int32(buf[3])<<24, nil
}

// ReadUint32 reads a uint32 from the buffer.
func (b *Buf) ReadUint32() (uint32, error) {
	var buf [4]byte
	var err error

	for i := 0; i < 4; i++ {
		buf[i], err = b.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// return uint32
	return uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24, nil
}

// ReadInt16 reads an int16 from the buffer.
func (b *Buf) ReadInt16() (int16, error) {
	var buf [2]byte
	var err error

	for i := 0; i < 2; i++ {
		buf[i], err = b.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// return int16
	return int16(buf[0]) | int16(buf[1])<<8, nil
}

// ReadUint16 reads a uint16 from the buffer.
func (b *Buf) ReadUint16() (uint16, error) {
	var buf [2]byte
	var err error

	for i := 0; i < 2; i++ {
		buf[i], err = b.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// return uint16
	return uint16(buf[0]) | uint16(buf[1])<<8, nil
}

// ReadInt8 reads an int8 from the buffer.
func (b *Buf) ReadInt8() (int8, error) {
	var buf [1]byte
	var err error

	for i := 0; i < 1; i++ {
		buf[i], err = b.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// return int8
	return int8(buf[0]), nil
}

// ReadUint8 reads a uint8 from the buffer.
func (b *Buf) ReadUint8() (uint8, error) {
	var buf [1]byte
	var err error

	for i := 0; i < 1; i++ {
		buf[i], err = b.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// return uint8
	return uint8(buf[0]), nil
}

// ReadString reads a string from the buffer.
func (b *Buf) ReadString(length int) (string, error) {
	// read buffer from b
	buf, err := b.ReadBytes(length)

	// check if we have an error
	if err != nil {
		return "", err
	}

	// return string
	return string(buf), nil
}

// ReadBytes reads a byte slice from the buffer.
func (b *Buf) ReadBytes(length int) ([]byte, error) {
	// create buffer
	buf := make([]byte, length)

	// read buffer from b
	err := b.ReadFromSlice(buf, length)

	// check if we have an error
	if err != nil {
		return nil, err
	}

	// return buffer
	return buf, nil
}

// GetLength reads the length of a blob from the buffer.
func (b *Buf) ReadLength(lengthType LenType) (int, error) {
	// get length
	var length int

	switch lengthType {
	case Uint8:
		l, err := b.ReadUint8()
		if err != nil {
			return 0, err
		}
		length = int(l)
	case Uint16:
		l, err := b.ReadUint16()
		if err != nil {
			return 0, err
		}
		length = int(l)
	case Uint32:
		l, err := b.ReadUint32()
		if err != nil {
			return 0, err
		}
		length = int(l)
	case Uint64:
		l, err := b.ReadUint64()
		if err != nil {
			return 0, err
		}
		length = int(l)
	case Int8:
		l, err := b.ReadInt8()
		if err != nil {
			return 0, err
		}
		length = int(l)
	case Int16:
		l, err := b.ReadInt16()
		if err != nil {
			return 0, err
		}
		length = int(l)
	case Int32:
		l, err := b.ReadInt32()
		if err != nil {
			return 0, err
		}
		length = int(l)
	case Int64:
		l, err := b.ReadInt64()
		if err != nil {
			return 0, err
		}
		length = int(l)
	default:
		return 0, fmt.Errorf("invalid length type: %d", lengthType)
	}
	return length, nil
}

// GetBlob reads a blob from the buffer by decoding the length information and then reading the blob length.
func (b *Buf) ReadBlob() ([]byte, error) {
	// types
	var val byte
	var err error

	// get length type
	val, err = b.ReadByte()
	if err != nil {
		return nil, err
	}

	var lengthType LenType = LenType(uint8(val))

	// get length
	length, err := b.ReadLength(lengthType)

	// check if we have an error
	if err != nil {
		return nil, err
	}

	// read blob
	return b.ReadBytes(length)
}

// ReadBlobAt reads the nth blob from the buffer.
// WARNING: This will reset the pointer to after the blob
/*func (b *Buf) ReadBlobAt(n int) ([]byte, error) {
	// types
	var val byte
	var err error

	// get length type
}*/
