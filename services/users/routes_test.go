package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ChrisBryann/go-ecommerce/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	usersStore := &mockUsersStore{}
	handler := NewHandler(usersStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "",
			Password:  "asdadawd",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUsersStore struct{}

func (m *mockUsersStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUsersStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUsersStore) CreateUser(types.User) error {
	return nil
}
