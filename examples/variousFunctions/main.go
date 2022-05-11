package main

import (
	"fmt"
	"log"
	"net"

	buflib "github.com/brandonvessel/ezgobuf/pkg/buffer"
)

// some examples of how to use the buffer lib

func main() {
	// Create a new buffer
	buf := buflib.NewBuf()

	// Write a byte to the buffer
	buf.WriteByte(130)

	// reset the pointer
	buf.ResetPtr()

	// Read a byte from the buffer
	mybyte, err := buf.ReadByte()

	// Check if the read operation failed
	if err != nil {
		// Print the error
		log.Panicln("error:", err.Error())
	}

	// print the output
	println(int(mybyte))

	// create array of numbers 1 to 10
	nums := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums2 := []byte{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	// reset pointer
	buf.ResetPtr()

	// write the array to the buffer
	err = buf.WriteBlob(nums)

	// check if the write operation failed
	if err != nil {
		// print the error
		log.Panicln("error:", err.Error())
	}

	// write second blob
	err = buf.WriteBlob(nums2)

	// check if the write operation failed
	if err != nil {
		// print the error
		log.Panicln("error:", err.Error())
	}

	// reset the pointer
	buf.ResetPtr()

	// read the array from the buffer
	mynums, err := buf.ReadBlob()

	// check if the read operation failed
	if err != nil {
		// Print the error
		log.Panicln("error:", err.Error())
	}

	// print the output
	fmt.Printf("%v\n", mynums)

	// read second blob
	mynums2, err := buf.ReadBlob()

	// check if the read operation failed
	if err != nil {
		// Print the error
		log.Panicln("error:", err.Error())
	}

	// print the output
	fmt.Printf("%v\n", mynums2)

	// print buffer
	fmt.Println(buf)

	// make a new buffer
	buf2 := buflib.NewBuf()

	// ip as a string
	ip := "192.168.1.1"

	// convert ip to bytes
	ipBytes := net.ParseIP(ip)

	// write ip bytes to buf2
	err = buf2.WriteBlob(ipBytes.To4())

	// check if the write operation failed
	if err != nil {
		// print the error
		log.Panicln("error:", err.Error())
	}

	// reset the pointer
	buf2.ResetPtr()

	// read ip bytes from buf2
	myipBytes, err := buf2.ReadBlob()

	// check if the read operation failed
	if err != nil {
		// Print the error
		log.Panicln("error:", err.Error())
	}

	// print the input
	fmt.Printf("%v\n", ipBytes.To4())

	// print the buffer
	fmt.Println(buf2)

	// print the output
	fmt.Printf("%v\n", myipBytes)

	// create new buffer
	buf3 := buflib.NewBuf()

	// create new byte with numbers 1 to 100000
	nums3 := make([]byte, 100000)

	// create a for loop to fill the array
	for i := 0; i < 100000; i++ {
		nums3[i] = byte(i)
	}

	// write the array to the buffer
	err = buf3.WriteBlob(nums3)

	// check if the write operation failed
	if err != nil {
		// print the error
		log.Panicln("error:", err.Error())
	}

	// reset the pointer
	buf3.ResetPtr()

	// read the array from the buffer
	mynums3, err := buf3.ReadBlob()

	// check if the read operation failed
	if err != nil {
		// Print the error
		log.Panicln("error:", err.Error())
	}

	// print the output
	//fmt.Printf("%v\n", mynums3)

	// check if the input buffer equals the output buffer
	match := true
	for i := 0; i < 100000; i++ {
		if nums3[i] != mynums3[i] {
			fmt.Println("error:", "buffers do not match")
			match = false
			break
		}
	}

	// check if the buffers match
	if match {
		fmt.Println("success:", "buffers match")
	}
}
