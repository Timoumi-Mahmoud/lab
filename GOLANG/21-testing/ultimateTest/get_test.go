package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const succeed = "\u2717"
const failed = "\u2717"

var feed = "Pong ......\n"

//	func mockServer() *httptest.Server {
//		f := func(w http.ResponseWriter, r *http.Request) {
//			w.WriteHeader(200)
//			fmt.Fprintln(w, feed)
//		}
//
//		return httptest.NewServer(http.HandlerFunc(f))
//	}
func TestHttp(t *testing.T) {

	testCases := []struct {
		name         string
		url          string
		responseCode int
		responseBody string
	}{
		{
			fmt.Sprintf("Should make a positive request %v", succeed),
			"http://localhost:8080/ping",
			200,
			"Pong ......\n",
		},
		{
			fmt.Sprintf("Should make a negative request: %v", failed),
			"http://localhost:8080/Wrongping",
			404,
			"404 page not found\n",
		},
	}

	for _, v := range testCases {
		t.Log(v.name)
		resp, err := http.Get(v.url)
		defer resp.Body.Close()
		if err != nil {
			t.Error(err)
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != v.responseCode {
			t.Errorf("Status Code: Want %v, Got %v\n", v.responseCode, resp.StatusCode)
		}

		returnedResponse := string(data)
		if returnedResponse != v.responseBody {
			t.Errorf("Response: Want %v, Got %v\n", v.responseBody, returnedResponse)
		}

	}
}
