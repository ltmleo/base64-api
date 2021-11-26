package main

import (
	b64 "encoding/base64"
	"errors"
)

var invalidStringError = errors.New("Invalid String")

func Encode(stringToBeEncoded string) string {
	return b64.StdEncoding.EncodeToString([]byte(stringToBeEncoded))
}

func Decode(stringToBeDecoded string) (string, error) {
	decoded, err := b64.StdEncoding.DecodeString(stringToBeDecoded)
	if err != nil {
		return "", invalidStringError
	}
	return string(decoded), nil
}