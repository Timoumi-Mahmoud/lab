package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {

	subtest := []struct {
		Url            string
		ResponseBody   string
		ResponseStatus string
	}{
		{"/upper?word=hello", "HELLO", "200 OK"},
		{"/upper?word=mahmoud", "MAHMOUD", "200 OK"},
		{"/upper?word=abc", "ABC", "200 OK"},
		{"/upper?word=ABC", "ABC", "200 OK"},
		{"/upper?word=", "missing word", "400 Bad Request"},
		{"/upper?word=%dkfdjf%", "invalid request", "400 Bad Request"},
	}
	for _, st := range subtest {
		req := httptest.NewRequest("GET", st.Url, nil)
		w := httptest.NewRecorder()
		upperCaseHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil , got %v", err)
		}

		if res.Status != st.ResponseStatus {
			t.Errorf("expected status code to be %v, got %v", st.ResponseStatus, res.Status)
		}
		if string(data) != st.ResponseBody {
			t.Errorf("expected %v, got %v", st.ResponseBody, string(data))
		}

	}
	/*
		req := httptest.NewRequest("GET", "/upper?word=hello", nil)
		w := httptest.NewRecorder()
		upperCaseHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil , got %v", err)
		}
		fmt.Println(res.Status)
		if res.Status != "200 OK" {
			t.Errorf("expected status code to be 200 , got %v", res.Status)
		}
		if string(data) != "HELLO" {
			t.Errorf("expected HELLO , got %v", string(data))
		}
	*/
}
