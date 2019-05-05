package sc2

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func getTestClient(status int, body []byte) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write(body)
		r.Body.Close()
	}))

	client := NewClient(&ClientConfig{
		Region:  "us",
		Test:    true,
		TestUrl: server.URL + "/",
	})

	return client, server
}

func getBodyFromFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

func byteSlicesMatch(a, b []byte) bool {
	return bytes.Equal(a, b)
}

func writeTestDebugOutput(a, b []byte, testname string) {
	out := "test_output/" + testname + "_output.json"
	fix := "test_output/" + testname + "_fixture.json"
	err := ioutil.WriteFile(out, a, 0644)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fix, b, 0644)
	if err != nil {
		panic(err)
	}

}
