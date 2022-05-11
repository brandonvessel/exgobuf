package buffer

// LenType is the type of the length field in the buffer for a blob.
type LenType uint8

const (
	Int8 LenType = iota
	Int16
	Int32
	Int64
	Uint8
	Uint16
	Uint32
	Uint64
)

func (l LenType) String() string {
	switch l {
	case Int8:
		return "int8"
	case Int16:
		return "int16"
	case Int32:
		return "int32"
	case Int64:
		return "int64"
	case Uint8:
		return "uint8"
	case Uint16:
		return "uint16"
	case Uint32:
		return "uint32"
	case Uint64:
		return "uint64"
	default:
		return "unknown"
	}
}
