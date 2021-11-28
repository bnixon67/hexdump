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

// IsPrintable returns true if the byte b is in the range of printable ASCII
// characters, otherwise returns false.
func IsPrintable(b byte) bool {
	if b >= ' ' && b <= '~' {
		return true
	}
	return false
}

// HexDump outputs a file provided on the command line
func HexDump(filename string) error {
	// open filename provided
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// define buffer to read the file in chunks
	// small size may not be performant, but simplifies program
	const bufferSize = 16
	buffer := make([]byte, bufferSize)

	// offset stores the current number of bytes read
	offset := 0

	for {
		// read a buffer size chunk of the file
		// use ReadFull to avoid incomplete reads
		// for files like /dev/random
		bytesRead, err := io.ReadFull(file, buffer)
		if err != nil {
			// exit loop on EOF
			if err == io.EOF {
				break
			}

			// display error if not UnexpectedEOF
			if err != io.ErrUnexpectedEOF {
				return err
			}

			// fall through on UnexpectedEOF
		}

		// print a line for each chunk read
		// offset
		fmt.Printf("%08x  ", offset)
		// bytes values in hex
		for n := 0; n < len(buffer); n++ {
			if n < bytesRead {
				fmt.Printf("%02x ", buffer[n])
			} else {
				fmt.Print("   ")
			}
		}
		// printable_bytes
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

	// total bytes read
	fmt.Printf("%08x\n", offset)

	return nil
}
