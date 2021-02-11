package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gustvision/backend-interview/pkg/user/dto"
	"github.com/stretchr/testify/assert"
)

func TestGetValidUser(t *testing.T) {
	h, err := newTestHandler()
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	b, err := json.Marshal(dto.GetUserReq{ID: "testuid"})
	if err != nil {
		t.Fatal(err)
	}

	request := httptest.NewRequest(
		"GET",
		"/user",
		bytes.NewBuffer(b))
	hf := http.HandlerFunc(h.GetUser)
	hf.ServeHTTP(recorder, request)

	assert.EqualValues(t,
		http.StatusOK,
		recorder.Code,
		"wrong status code")

	expectedBody := `{"ID":"testuid","Name":"testname","Total":2300}`
	assert.Equal(t,
		expectedBody,
		recorder.Body.String(),
		"wrong user")
}
