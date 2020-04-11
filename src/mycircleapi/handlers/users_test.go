package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	//Fixture
	w := httptest.NewRecorder()
	r := new(http.Request)
	r.URL.Path = "/users/test"

	//call
	usr := NewUserController()
	usr.Set(w, r)

	//assert
	if w.Code != 200 {
		t.Error("Expected 200 but got ", w.Code)
	}
}
