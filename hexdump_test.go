package main

import (
	"os"
	"testing"
)

func TestHexDump(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		output      string
		expectedErr error
	}{
		{
			name:        "Empty File",
			input:       "testdata/empty.txt",
			output:      "testdata/empty.txt.out",
			expectedErr: nil,
		},
		{
			name:        "Small File",
			input:       "testdata/small.txt",
			output:      "testdata/small.txt.out",
			expectedErr: nil,
		},
		{
			name:        "Binary File",
			input:       "testdata/binary",
			output:      "testdata/binary.out",
			expectedErr: nil,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := HexDump(tc.input)

			if err != tc.expectedErr {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}

			data, err := os.ReadFile(tc.output)
			if err != nil {
				t.Fatal(err)
			}

			if actual != string(data) {
				t.Errorf("unexpected output:\nexpected:\n%s\ngot:\n%s", data, actual)
			}
		})
	}
}
