package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// type RegisterData struct {
// 	Username string `json:"username"`
// 	Account  string `json:"account"`
// 	Password string `json:"password"`
// }

type CreateArticleResponseData struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type ArticleContent struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetArticleResponseData struct {
	Article ArticleContent `json:"article"`
	Status  string         `json:"status"`
}

type GetArticlesResponseData struct {
	Articles []ArticleContent `json:"articles"`
	Status   string           `json:"status"`
}

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

func TestLoginGetUpdateUserLogoutRoute(t *testing.T) {
	jsonData := `{
		"account": "testcase1@gmail.com",
		"password": "123456"
	}`

	router := setRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(string(jsonData)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"status":"you are logged in"}`, w.Body.String())
	assert.NotEqual(t, "", w.Header().Get("Authorization"))

	authToken := w.Header().Get("Authorization")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/user", nil)
	req.Header.Add("Authorization", authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"account":"testcase1@gmail.com","status":"success","username":"testcase1"}`, w.Body.String())

	updateData := `{"username": "testcase2"}`
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/user", strings.NewReader(string(updateData)))
	req.Header.Add("Authorization", authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"account":"testcase1@gmail.com","status":"success","username":"testcase2"}`, w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/logout", nil)
	req.Header.Add("Authorization", authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"status":"you are logged out"}`, w.Body.String())
}

func TestCreateUpdateGetGetManyDeleteArticleRoute(t *testing.T) {
	jsonData := `{
		"account": "testcase1@gmail.com",
		"password": "123456"
	}`

	router := setRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(string(jsonData)))
	router.ServeHTTP(w, req)

	authToken := w.Header().Get("Authorization")

	// create article test
	createArticleData := `{
		"title": "testcase",
		"content": "testcase content"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/article", strings.NewReader(string(createArticleData)))
	req.Header.Add("Authorization", authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var createResponse CreateArticleResponseData
	_ = json.Unmarshal([]byte(w.Body.Bytes()), &createResponse)
	assert.Equal(t, "Article created.", createResponse.Status)

	// get article test
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/article/"+fmt.Sprint(createResponse.ID), nil)
	req.Header.Add("Authorization", authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var getArticleResponse GetArticleResponseData
	_ = json.Unmarshal([]byte(w.Body.Bytes()), &getArticleResponse)
	assert.Equal(t, "success", getArticleResponse.Status)
	assert.Equal(t, createResponse.ID, getArticleResponse.Article.ID)
	assert.Equal(t, "testcase", getArticleResponse.Article.Title)
	assert.Equal(t, "testcase content", getArticleResponse.Article.Content)

	// update article test
	updateArticleData := `{
		"title": "testcase9",
		"content": "testcase9 content"
	}`

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/article/"+fmt.Sprint(createResponse.ID), strings.NewReader(string(updateArticleData)))
	req.Header.Add("Authorization", authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var updateResponse CreateArticleResponseData
	_ = json.Unmarshal([]byte(w.Body.Bytes()), &updateResponse)
	assert.Equal(t, "success", updateResponse.Status)
	assert.Equal(t, createResponse.ID, updateResponse.ID)

	// get articles test
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/article", nil)
	req.Header.Add("Authorization", authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var getArticlesResponse GetArticlesResponseData
	_ = json.Unmarshal([]byte(w.Body.Bytes()), &getArticlesResponse)
	assert.Equal(t, "success", getArticlesResponse.Status)
	assert.Equal(t, createResponse.ID, getArticlesResponse.Articles[0].ID)
	assert.Equal(t, "testcase9", getArticlesResponse.Articles[0].Title)
	assert.Equal(t, "testcase9 content", getArticlesResponse.Articles[0].Content)

	// delete article test
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/article/"+fmt.Sprint(createResponse.ID), nil)
	req.Header.Add("Authorization", authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"status":"success"}`, w.Body.String())
}

func TestLoginDeleteUserRoute(t *testing.T) {
	jsonData := `{
		"account": "testcase1@gmail.com",
		"password": "123456"
	}`

	router := setRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(string(jsonData)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"status":"you are logged in"}`, w.Body.String())
	assert.NotEqual(t, "", w.Header().Get("Authorization"))

	authToken := w.Header().Get("Authorization")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/user", nil)
	req.Header.Add("Authorization", authToken)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"status":"success"}`, w.Body.String())
}
