package main

var testCases = []struct {
	testName string
	stringToBeEncoded string
	stringEncoded string

}{
	{"String", "Leonardo", "TGVvbmFyZG8="},
	{"Integer", "123456", "MTIzNDU2"},
	{"Special Characters", "!@#$%", "IUAjJCU="},
	{"Multiple Characters", "Hello World!! 123 #$ :)", "SGVsbG8gV29ybGQhISAxMjMgIyQgOik="},
	{"Empty String", "", ""},
}
var testErrorCases = []struct {
	testName string
	stringToBeDecoded string
	decodedError error
}{
	{"Not Base64 format", "this_is_not_base64", invalidStringError},
}
