package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gustvision/backend-interview/pkg/account/dto"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction_Withdraw(t *testing.T) {
	h, err := newTestHandler()
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	req := dto.CreateTransactionReq{
		AccountID: "testaid3",
		Amount:    -50,
	}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	request := httptest.NewRequest(
		"POST",
		"/transaction",
		bytes.NewBuffer(b))
	hf := http.HandlerFunc(h.CreateTransaction)
	hf.ServeHTTP(recorder, request)

	assert.EqualValues(t,
		http.StatusOK,
		recorder.Code,
		"wrong status code")

	result := dto.CreateTransactionResp{}
	err = json.Unmarshal(recorder.Body.Bytes(), &result)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t,
		req.Amount,
		result.Amount,
		"wrong amount")

	assert.Equal(t,
		req.AccountID,
		result.AccountID,
		"wrong account ID")

	assert.Equal(t,
		1750.0,
		result.AccountTotal,
		"wrong account total")

	assert.True(t, result.CreatedAt <= time.Now().Unix())
	assert.True(t, result.CreatedAt > time.Now().Add(-3*time.Second).Unix())
}

func TestCreateTransaction_Deposit(t *testing.T) {
	h, err := newTestHandler()
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	req := dto.CreateTransactionReq{
		AccountID: "testaid3",
		Amount:    50,
	}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	request := httptest.NewRequest(
		"POST",
		"/transaction",
		bytes.NewBuffer(b))
	hf := http.HandlerFunc(h.CreateTransaction)
	hf.ServeHTTP(recorder, request)

	assert.EqualValues(t,
		http.StatusOK,
		recorder.Code,
		"wrong status code")

	result := dto.CreateTransactionResp{}
	err = json.Unmarshal(recorder.Body.Bytes(), &result)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t,
		req.Amount,
		result.Amount,
		"wrong amount")

	assert.Equal(t,
		req.AccountID,
		result.AccountID,
		"wrong account ID")

	assert.Equal(t,
		1800.0,
		result.AccountTotal,
		"wrong account total")

	assert.True(t, result.CreatedAt <= time.Now().Unix())
	assert.True(t, result.CreatedAt > time.Now().Add(-3*time.Second).Unix())
}

func TestGetMaxTransaction(t *testing.T) {
	h, err := newTestHandler()
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	b, err := json.Marshal(dto.GetMaxTransactionReq{
		UserID: "testuid",
		From:   40,
		To:     60,
	})
	if err != nil {
		t.Fatal(err)
	}

	request := httptest.NewRequest(
		"GET",
		"/max-transaction",
		bytes.NewBuffer(b))
	hf := http.HandlerFunc(h.GetMaxTransaction)
	hf.ServeHTTP(recorder, request)

	assert.EqualValues(t,
		http.StatusOK,
		recorder.Code,
		"wrong status code")

	expectedBody := `{"ID":"testtx3","Amount":300,"AccountID":"testaid1","CreatedAt":60}`
	assert.Equal(t,
		expectedBody,
		recorder.Body.String(),
		"wrong user")
}
