package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//TestSet Test Set message
func TestSet(t *testing.T) {
	t.Skip() //skip for now
	//Fixture
	w := httptest.NewRecorder()
	r := new(http.Request)
	//r.URL.Path = "/Message"
	r.Method = "POST"
	//TODO Add body payload

	//call
	msg := NewMsgController()
	msg.Set(w, r)

	//assert
	if w.Code != 200 {
		t.Error("Expected 200 but got ", w.Code)
	}
}

//TestGet Test Get message
func TestGet(t *testing.T) {
	//Fixture
	w := httptest.NewRecorder()
	r := new(http.Request)
	//r.URL.Path = "/Message"

	//call
	msg := NewMsgController()
	msg.Get(w, r)

	//assert
	if w.Code != 200 {
		t.Error("Expected 200 but got ", w.Code)
	}
}

// BenchmarkSetMsg Performance metrics
func BenchmarkSetMsg(b *testing.B) {
	//Fixture
	w := httptest.NewRecorder()
	r := new(http.Request)
	r.URL.Path = "/Message"
	r.Method = "POST"
	//TODO Add body payload

	msg := NewMsgController()

	for i := 0; i < b.N; i++ {
		msg.Set(w, r)
	}
}
