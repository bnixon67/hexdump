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
	"strings"
)

// IsPrintable determines if b is a printable ASCII characters.
func IsPrintable(b byte) bool {
	if b >= ' ' && b <= '~' {
		return true
	}
	return false
}

// HexDump reads a file and returns its contents as a formatted hex dump.
func HexDump(filename string) (string, error) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a string builder to store the formatted output
	var sb strings.Builder

	// Create a buffer to read the file in chunks
	// The small size may not be performant, but simplifies program logic
	const bufferSize = 16
	buffer := make([]byte, bufferSize)

	// Initialize offset for tracking the number of bytes read
	offset := 0

	for {
		// Read a chunk of data from the file
		// use ReadFull to avoid incomplete reads like /dev/random
		bytesRead, err := io.ReadFull(file, buffer)
		if err != nil {
			if err == io.EOF {
				// If EOF is encountered, break the loop
				break
			}

			if err != io.ErrUnexpectedEOF {
				// Return error if it's not an UnexpectedEOF
				return "", err
			}

			// fall through on UnexpectedEOF
		}

		// Format and write the offset
		sb.WriteString(fmt.Sprintf("%08x  ", offset))

		// Write the bytes in hexadecimal format
		for i := 0; i < len(buffer); i++ {
			if i < bytesRead {
				sb.WriteString(fmt.Sprintf("%02x ", buffer[i]))
			} else {
				sb.WriteString("   ")
			}
		}

		// Write the printable bytes
		for i := 0; i < len(buffer); i++ {
			if i < bytesRead {
				if IsPrintable(buffer[i]) {
					sb.WriteByte(buffer[i])
				} else {
					sb.WriteByte('.')
				}
			}
		}

		sb.WriteString("\n")

		// Update the offset
		offset += bytesRead
	}

	// Write the total number of bytes read
	sb.WriteString(fmt.Sprintf("%08x\n", offset))

	return sb.String(), nil
}
