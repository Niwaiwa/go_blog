package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type RegisterData struct {
// 	Username string `json:"username"`
// 	Account  string `json:"account"`
// 	Password string `json:"password"`
// }

func TestPingRoute(t *testing.T) {
	router := setRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome Go Blog Server", w.Body.String())
}

func TestRegisterRoute(t *testing.T) {
	// data := &RegisterData{
	// 	Username: "testcase1",
	// 	Account:  "testcase1@gmail.com",
	// 	Password: "123456",
	// }
	// jsonData, _ := json.Marshal(data)

	jsonData := `{
		"username": "testcase1",
		"account": "testcase1@gmail.com",
		"password": "123456"
	}`

	router := setRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(jsonData)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"status":"account created."}`, w.Body.String())
}
