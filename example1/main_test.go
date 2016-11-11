package main

import (
	"net/http/httptest"
	"testing"
)

func TestRequestSucceeds(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	helloWorld(w, r)
	if w.Code != 200 {
		t.Errorf("Oh no, I failed with code %d", w.Code)
	}
}
