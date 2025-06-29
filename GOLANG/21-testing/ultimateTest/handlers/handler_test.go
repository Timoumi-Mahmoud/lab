package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	Routes()
}

func TestSendJSON(t *testing.T) {
	url := "/user/get"
	statusCode := 200
	r := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	t.Logf("\tTest :!\twhen checking %q for status code %d", url, statusCode)
	{
		if w.Code != 200 {
			t.Fatalf("FAILD: \tShould receive a status code of %d for the response. Received %d", statusCode, w.Code)
		}
		t.Logf("PASS: Should receive a status code of %d for the response", statusCode)

		var u struct {
			Name  string
			Email string
		}
		if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
			t.Fatal("FAILD:\tShould be able to decode the response")
		}

		t.Log("Should be able to decode the reponse")
		if u.Name == "mahmoud" {
			t.Log("PASS: \tshould have \"mahmoud\" for Name in the respond")
		} else {
			t.Logf("FAILD: \tshould have \"mahmoud\" for Name in the respond, but got:%q", u.Name)
		}
		if u.Email == "mahmoud@gmail.com" {
			t.Log("PASS: \tshould have \"mahmoud@gmail.com\" for Email in the respond")
		} else {
			t.Logf("FAILD: \tshould have \"mahmoud@gmail.com\" for Email in the respond, but got:%q", u.Email)
		}

	}
}

func TestBlabla(t *testing.T) {
	t.Log("test Blablabla")
}
