package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/14-bits/crowd/pkg/sink"
	"github.com/julienschmidt/httprouter"
)

func TestCrowd(t *testing.T) {
	c := Crowd{
		S: sink.NewVoid(),
	}

	router := httprouter.New()
	router.POST("/api/foo", c.Handle)

	data := url.Values{}
	data.Add("foo", "bar")

	req, _ := http.NewRequest("POST", "/api/foo", bytes.NewBufferString(data.Encode()))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Returned with wrong status. Got %d, expected %d", status, http.StatusOK)
	}
}
