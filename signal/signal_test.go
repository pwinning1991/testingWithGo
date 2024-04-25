package signal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("http.NewRequest: %v", err)
	}
	Handler(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Handler() status = %v, want %v", resp.StatusCode, http.StatusOK)

	}
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Handler() Content-Type = %q, want %q", contentType, "application/json")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll: %v", err)
	}

	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if p.Age != 30 {
		t.Errorf("person.Age = %v, want %v", p.Age, 21)
	}
	if p.Name != "Bob Jones" {
		t.Errorf("person.Name = %v, want %v", p.Name, "Bob Jones")
	}
}
