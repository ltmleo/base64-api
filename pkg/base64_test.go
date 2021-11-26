package main

import (
	"testing"
)


func TestEncode(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			assert(t, tt.stringToBeEncoded, tt.stringEncoded, Encode(tt.stringToBeEncoded))
		})
	}
}
func TestDecode(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			decodedString, _ := Decode(tt.stringEncoded)
			assert(t, tt.stringEncoded, tt.stringToBeEncoded, decodedString)

		})
	}
}
func TestDecodeErrors(t *testing.T) {
	for _, tt := range testErrorCases {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := Decode(tt.stringToBeDecoded)
			assertError(t, err, tt.decodedError)
		})
	}
}

func assert(t testing.TB, stringToBeEncoded string, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Encode(%q) = %q, want %q", stringToBeEncoded, got, want)
	}
}

func assertError (t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}