package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEncode(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			Assert(t, tt.stringToBeEncoded, tt.stringEncoded, "/encode", GetEncode)
		})
	}
}
func TestGetDecode(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			Assert(t, tt.stringEncoded, tt.stringToBeEncoded, "/decode", GetDecode)
		})
	}
}

func Assert(t testing.TB, stringToArg string, expected string, path string, function func(w http.ResponseWriter, r *http.Request)) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("string", stringToArg)
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(function)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	s := rr.Body.String()
	if s[1 : len(s)-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			s, expected)
	}
}
