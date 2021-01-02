/*
Copyright 2021 Bill Nixon

This program is free software: you can redistribute it and/or modify it
under the terms of the GNU General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your option)
any later version.

This program is distributed in the hope that it will be useful, but WITHOUT
ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with
this program.  If not, see <http://www.gnu.org/licenses/>.
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
		fmt.Println(err)
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
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			// exit loop on EOF
			break
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
