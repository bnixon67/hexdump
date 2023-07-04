/*
Copyright 2023 Bill Nixon

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
	"os"
)

const (
	EXIT_OK        = iota // Exit status for successful execution
	EXIT_ERR_USAGE        // Exit status for incorrect usage
	EXIT_ERR              // Exit status for general error
)

// Usage prints usage information for the program.
func Usage(name string) {
	fmt.Printf("usage: %s file\n", name)
}

func main() {
	// ensure filename provided on command line
	if len(os.Args) != 2 {
		Usage(os.Args[0])
		os.Exit(EXIT_ERR_USAGE)
	}

	s, err := HexDump(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(EXIT_ERR)
	}
	fmt.Print(s)

	os.Exit(EXIT_OK)
}
