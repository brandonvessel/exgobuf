package buffer

import (
	"reflect"
	"testing"
)

func TestNewBuf(t *testing.T) {
	tests := []struct {
		name string
		want *Buf
	}{
		// test cases
		{"test 1", &Buf{buf: &[]byte{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBufMax(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name    string
		args    args
		want    *Buf
		wantErr bool
	}{
		// test cases
		{"test 1", args{10}, &Buf{buf: &[]byte{}, max: 10}, false},
		{"test 2", args{0}, &Buf{buf: &[]byte{}, max: 0}, false},
		{"test 3", args{-1}, &Buf{buf: &[]byte{}, max: 0}, false},
		{"test 4", args{-10}, &Buf{buf: &[]byte{}, max: 0}, false},
		{"test 5", args{-0}, &Buf{buf: &[]byte{}, max: 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBufMax(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBufMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBufFromBytes(t *testing.T) {
	type args struct {
		buf *[]byte
	}
	tests := []struct {
		name string
		args args
		want *Buf
	}{
		// test cases
		{"test 1", args{&[]byte{}}, &Buf{buf: &[]byte{}}},
		{"test 2", args{&[]byte{1, 2, 3}}, &Buf{buf: &[]byte{1, 2, 3}}},
		{"test 3", args{&[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, &Buf{buf: &[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBufFromBytes(tt.args.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBufFromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBufMaxFromBytes(t *testing.T) {
	type args struct {
		buf *[]byte
		max int
	}
	tests := []struct {
		name string
		args args
		want *Buf
	}{
		// test cases
		{"test 1", args{&[]byte{}, 10}, &Buf{buf: &[]byte{}, max: 10}},
		{"test 2", args{&[]byte{1, 2, 3}, 10}, &Buf{buf: &[]byte{1, 2, 3}, max: 10}},
		{"test 3", args{&[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10}, &Buf{buf: &[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, max: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBufMaxFromBytes(tt.args.buf, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBufMaxFromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuf_ResetPtr(t *testing.T) {
	tests := []struct {
		name string
		b    *Buf
	}{
		// test cases
		{"test 1", &Buf{buf: &[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}},
		{"test 2", &Buf{buf: &[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, max: 10}},
		{"test 3", &Buf{buf: &[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, max: 10, ptr: 5}},
		{"test 4", &Buf{buf: &[]byte{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.ResetPtr()
		})
	}
}

func TestBuf_ReadByte(t *testing.T) {
	tests := []struct {
		name    string
		b       *Buf
		want    byte
		wantErr bool
	}{
		// test cases
		{"test 1", &Buf{buf: &[]byte{1, 2, 3}}, 1, false},
		{"test 2", &Buf{buf: &[]byte{2, 1, 3}}, 2, false},
		{"test 3", &Buf{buf: &[]byte{3, 2, 1}}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.ReadByte()
			if (err != nil) != tt.wantErr {
				t.Errorf("Buf.ReadByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Buf.ReadByte() = %v, want %v", got, tt.want)
			}
		})
	}

	// multi-byte read test
	t.Run("test 4", func(t *testing.T) {
		b := &Buf{buf: &[]byte{1, 2, 3}}
		want := []byte{1, 2, 3}

		for i := 0; i < 3; i++ {
			got, err := b.ReadByte()
			if err != nil {
				t.Errorf("Buf.ReadByte() error = %v, wantErr %v", err, false)
				return
			}
			if got != want[i] {
				t.Errorf("Buf.ReadByte() = %v, want %v", got, want[i])
			}
		}
	})
}

func TestBuf_WriteByte(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name    string
		b       *Buf
		args    args
		wantErr bool
	}{
		// test cases
		{"test 1", &Buf{buf: &[]byte{}}, args{1}, false},
		{"test 2", &Buf{buf: &[]byte{1, 2, 3}}, args{4}, false},
		{"test 3", &Buf{buf: &[]byte{1, 2, 3}}, args{5}, false},
		{"test 4", &Buf{buf: &[]byte{1, 2, 3}}, args{6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.WriteByte(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Buf.WriteByte() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// multi-byte write test followed by reset and multi-byte read test
	t.Run("test 5", func(t *testing.T) {
		b := &Buf{buf: &[]byte{}}
		input := []byte{1, 2, 3}
		want := []byte{1, 2, 3}

		for i := 0; i < 3; i++ {
			err := b.WriteByte(input[i])
			if err != nil {
				t.Errorf("Buf.WriteByte() error = %v, wantErr %v", err, false)
				return
			}
		}

		b.ResetPtr()

		for i := 0; i < 3; i++ {
			got, err := b.ReadByte()
			if err != nil {
				t.Errorf("Buf.ReadByte() error = %v, wantErr %v", err, false)
				return
			}
			if got != want[i] {
				t.Errorf("Buf.ReadByte() = %v, want %v", got, want[i])
			}
		}
	})
}

func TestBuf_ReadFromSlice(t *testing.T) {
	type args struct {
		p []byte
		i int
	}
	tests := []struct {
		name    string
		b       *Buf
		args    args
		wantErr bool
	}{
		// test cases
		{"test 1", &Buf{buf: &[]byte{1, 2, 3}}, args{[]byte{1, 2, 3}, 0}, true},
		{"test 2", &Buf{buf: &[]byte{1, 2, 3}}, args{[]byte{1, 2, 3}, 1}, false},
		{"test 3", &Buf{buf: &[]byte{1, 2, 3}}, args{[]byte{1, 2, 3}, 2}, false},
		{"test 4", &Buf{buf: &[]byte{1, 2, 3}}, args{[]byte{1, 2, 3}, 3}, false},
		{"test 5", &Buf{buf: &[]byte{1, 2, 3}}, args{[]byte{1, 2, 3}, 4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.ReadFromSlice(tt.args.p, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Buf.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// multi-byte read test
	t.Run("test 6", func(t *testing.T) {
		b := &Buf{buf: &[]byte{1, 2, 3}}
		output := []byte{0, 0, 0}
		want := []byte{1, 2, 3}

		err := b.ReadFromSlice(output, 3)
		if err != nil {
			t.Errorf("Buf.Read() error = %v, wantErr %v", err, false)
			return
		}

		// determine if output equals want
		for i := 0; i < 3; i++ {
			if output[i] != want[i] {
				t.Errorf("Buf.Read() = %v, want %v", output[i], want[i])
			}
		}
	})
}

func TestBuf_WriteFromSlice(t *testing.T) {
	type args struct {
		p []byte
		i int
	}
	tests := []struct {
		name    string
		b       *Buf
		args    args
		wantErr bool
	}{
		// test cases
		{"test 1", &Buf{buf: &[]byte{}}, args{[]byte{1, 2, 3}, 0}, false},
		{"test 2", &Buf{buf: &[]byte{1, 2, 3}}, args{[]byte{1, 2, 3}, 1}, false},
		{"test 3", &Buf{buf: &[]byte{1, 2, 3}}, args{[]byte{1, 2, 3}, 2}, false},

		// limited buffer tests
		{"test 4", NewBufMaxFromBytes(&[]byte{}, 0), args{[]byte{1, 2, 3}, 3}, false},
		{"test 5", NewBufMaxFromBytes(&[]byte{}, 1), args{[]byte{1, 2, 3}, 3}, true},
		{"test 6", NewBufMaxFromBytes(&[]byte{}, 2), args{[]byte{1, 2, 3}, 3}, true},
		{"test 7", NewBufMaxFromBytes(&[]byte{}, 3), args{[]byte{1, 2, 3}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.WriteFromSlice(tt.args.p, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Buf.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuf_Len(t *testing.T) {
	tests := []struct {
		name string
		b    *Buf
		want int
	}{
		// test
		{"test 0", &Buf{buf: &[]byte{}}, 0},
		{"test 1", &Buf{buf: &[]byte{1}}, 1},
		{"test 2", &Buf{buf: &[]byte{1, 2}}, 2},
		{"test 3", &Buf{buf: &[]byte{1, 2, 3}}, 3},
		{"test 4", &Buf{buf: &[]byte{1, 2, 3, 4}}, 4},
		{"test 5", &Buf{buf: &[]byte{5, 4, 3, 2, 1}}, 5},
		{"test 6", &Buf{buf: &[]byte{1, 2, 3, 4, 5}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Len(); got != tt.want {
				t.Errorf("Buf.Len() = %v, want %v", got, tt.want)
			}
		})
	}

	// multi-byte len test
	t.Run("test 7", func(t *testing.T) {
		b := &Buf{buf: &[]byte{}}
		input := []byte{1, 2, 3}
		want := []int{1, 2, 3}

		for i := 0; i < 3; i++ {
			err := b.WriteByte(input[i])
			if err != nil {
				t.Errorf("Buf.WriteByte() error = %v, wantErr %v", err, false)
				return
			}
			if b.Len() != want[i] {
				t.Errorf("Buf.Len() = %v, want %v", b.Len(), want[i])
			}
		}
	})
}

func TestBuf_SetPtr(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		b       *Buf
		args    args
		wantErr bool
	}{
		// no test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.SetPtr(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Buf.SetPtr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// multi-byte pointer moving tests
	t.Run("test 1", func(t *testing.T) {
		b := &Buf{buf: &[]byte{1, 2, 3}}
		want := []int{1, 2, 3}
		for i := 0; i < 3; i++ {
			err := b.SetPtr(want[i])
			if err != nil {
				t.Errorf("Buf.SetPtr() error = %v, wantErr %v", err, false)
				return
			}
			if b.ptr != want[i] {
				t.Errorf("Buf.Ptr() = %v, want %v", b.ptr, want[i])
			}
		}
	})
}

func TestBuf_GetPtr(t *testing.T) {
	tests := []struct {
		name string
		b    *Buf
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetPtr(); got != tt.want {
				t.Errorf("Buf.GetPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuf_SetMax(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		b    *Buf
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.SetMax(tt.args.i)
		})
	}
}

func TestBuf_GetMax(t *testing.T) {
	tests := []struct {
		name string
		b    *Buf
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetMax(); got != tt.want {
				t.Errorf("Buf.GetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuf_DecrementPtr(t *testing.T) {
	tests := []struct {
		name    string
		b       *Buf
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.DecrementPtr(); (err != nil) != tt.wantErr {
				t.Errorf("Buf.DecrementPtr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
