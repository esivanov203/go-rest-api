package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/esivanov203/go-rest-api/internal/model"
	"github.com/esivanov203/go-rest-api/internal/store/teststore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := model.GetTestUser(t)
	store.GetUserRepo().Create(u)

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	sKey := []byte("secret")
	s := newServer(store, sessions.NewCookieStore([]byte(sKey)))
	sc := securecookie.New([]byte(sKey), nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookie, _ := sc.Encode(sessionName, testCase.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookie))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, testCase.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))

	cases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{name: "valid", payload: map[string]string{"email": "test@test.com", "password": "pwd"}, expectedCode: http.StatusCreated},
		{name: "invalid payload", payload: "bad", expectedCode: http.StatusBadRequest},
		{name: "invalid params", payload: map[string]string{"email": "test", "password": ""}, expectedCode: http.StatusUnprocessableEntity},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			body := &bytes.Buffer{}
			json.NewEncoder(body).Encode(c.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", body)
			s.ServeHTTP(rec, req)
			assert.Equal(t, c.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	store := teststore.New()
	u := model.GetTestUser(t)
	pwd := u.Password
	store.GetUserRepo().Create(u)
	server := newServer(store, sessions.NewCookieStore([]byte("secret")))

	cases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{name: "valid", payload: map[string]string{"email": u.Email, "password": pwd}, expectedCode: http.StatusOK},
		{name: "invalid payload", payload: "mypayload", expectedCode: http.StatusBadRequest},
		{name: "invalid username", payload: map[string]string{"email": "iuser", "password": u.Password}, expectedCode: http.StatusUnauthorized},
		{name: "invalid password", payload: map[string]string{"email": u.Email, "password": "wrong"}, expectedCode: http.StatusUnauthorized},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			body := &bytes.Buffer{}
			json.NewEncoder(body).Encode(c.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", body)
			server.ServeHTTP(rec, req)
			assert.Equal(t, c.expectedCode, rec.Code)
		})
	}
}
