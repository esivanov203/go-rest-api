package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/esivanov203/go-rest-api/internal/model"
	"github.com/esivanov203/go-rest-api/internal/store/teststore"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
