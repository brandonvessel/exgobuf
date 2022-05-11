package main

import (
	"fmt"
	"log"

	buflib "github.com/brandonvessel/ezgobuf/pkg/buffer"
)

func main() {
	// Create a new buffer
	buf := buflib.NewBuf()

	// create a byte slice with the values 1 to 10
	nums := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// write the array to the buffer as a blob
	err := buf.WriteBlob(nums)

	// check if the write operation failed
	if err != nil {
		// print the error
		log.Panicln("error:", err.Error())
	}

	// reset the pointer
	buf.ResetPtr()

	// read blob bytes with header
	blob, err := buf.ReadBlobWithHeader()

	// check if the read operation failed
	if err != nil {
		// print the error
		log.Panicln("error:", err.Error())
	}

	// print the input
	fmt.Printf("%v\n", nums)

	// print the output
	fmt.Printf("%v\n", blob)

	// print the length of the buffer
	fmt.Println("length:", buf.Len())

	// print the current pointer position
	fmt.Println("pointer:", buf.GetPtr())
}
