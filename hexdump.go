/*
Copyright 2021 Bill Nixon

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

// Usage prints usage information for the program
func Usage() {
	fmt.Printf("usage: %s file\n", os.Args[0])
}

// IsPrintable determines if a byte is in the range of ASCII printable characters
func IsPrintable(c byte) bool {
	if c >= ' ' && c <= '~' {
		return true
	}
	return false
}

func main() {
	// must provide a command line argument for the file name
	if len(os.Args) != 2 {
		Usage()
		return
	}

	// open file provided on command line
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	// buffer used to read the file in chunks
	const bufferSize = 16
	buffer := make([]byte, bufferSize)

	// offset stores the current number of bytes read
	offset := 0

	for {
		// read a buffer size chunk of the file
		// use ReadFull to avoid incomplete reads, e.g. /dev/random
		bytesRead, err := io.ReadFull(file, buffer)
		if err != nil {
			// exit loop on EOF
			if err == io.EOF {
				break
			}

			// display error if not UnexpectedEOF
			if err != io.ErrUnexpectedEOF {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			// fall through on UnexpectedEOF
		}

		// print a line for each chunk read
		// offset  bytes_values_in_hex  printable_bytes
		fmt.Printf("%08x  ", offset)
		for n := 0; n < len(buffer); n++ {
			if n < bytesRead {
				fmt.Printf("%02x ", buffer[n])
			} else {
				fmt.Print("   ")
			}
		}
		for n := 0; n < len(buffer); n++ {
			if n < bytesRead {
				if IsPrintable(buffer[n]) {
					fmt.Printf("%c", buffer[n])
				} else {
					fmt.Printf("%c", '.')
				}
			}
		}
		fmt.Println()

		offset += bytesRead
	}

	// print total bytes read
	fmt.Printf("%08x\n", offset)
}
