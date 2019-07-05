package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(Login)

	handler.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	loginSucc := `{"Code":200,"MSg":"Login Success","Data":null}`
	if loginSucc != recorder.Body.String() {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), loginSucc)
	}
	t.Log("Login Handler Success")
}