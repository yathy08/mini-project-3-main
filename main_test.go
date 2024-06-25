package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yathy08/mini-project3/internal/domain"
	"github.com/yathy08/mini-project3/internal/handler"
	"gopkg.in/h2non/gock.v1"
)

func TestGetAll(t *testing.T) {
	defer gock.Off()

	gock.New("https://reqres.in").
		Get("/api/users").
		Reply(200).
		JSON(map[string]interface{}{
			"data": []domain.User{
				{ID: 1, Email: "garzao@e.o.cara"},
			},
		})

	router := gin.Default()
	router.GET("/", handler.GetAll)

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %v; got %v", http.StatusOK, rr.Code)
	}

	var res handler.Users
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expected := []domain.User{{ID: 1, Email: "garzao@e.o.cara"}}
	if len(res.Data) != len(expected) {
		t.Fatalf("expected %d items; got %d items", len(expected), len(res.Data))
	}

	for i := range res.Data {
		if res.Data[i].ID != expected[i].ID || res.Data[i].Email != expected[i].Email {
			t.Fatalf("expected %v; got %v", expected[i], res.Data[i])
		}
	}
}

func TestGetByID(t *testing.T) {
	defer gock.Off()

	gock.New("https://reqres.in").
		Get("/api/users").
		Reply(200).
		JSON(map[string]interface{}{
			"data": []domain.User{
				{ID: 1, Email: "garzao@e.o.cara"},
			},
		})

	router := gin.Default()
	router.GET("/:id", handler.GetByID)

	req, _ := http.NewRequest("GET", "/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %v; got %v", http.StatusOK, rr.Code)
	}

	var res domain.User
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expected := domain.User{ID: 1, Email: "garzao@e.o.cara"}
	if res.ID != expected.ID || res.Email != expected.Email {
		t.Fatalf("expected %v; got %v", expected, res)
	}
}

func TestCreate(t *testing.T) {
	defer gock.Off()

	newUser := domain.User{Email: "newuser@example.com"}

	// Mock API response for creating a user
	gock.New("https://reqres.in").
		Post("/api/users").
		Reply(201).
		JSON(domain.User{ID: 2, Email: "newuser@example.com"})

	router := gin.Default()
	router.POST("/users", handler.Create)

	body, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected %v; got %v", http.StatusCreated, rr.Code)
	}

	var res domain.User
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expected := domain.User{ID: 2, Email: "newuser@example.com"}
	if res.ID != expected.ID || res.Email != expected.Email {
		t.Fatalf("expected %v; got %v", expected, res)
	}
}
func TestUpdate(t *testing.T) {
	defer gock.Off()

	updatedUser := domain.User{Email: "updateduser@example.com"}
	gock.New("https://reqres.in").
		Put("/api/users/1").
		Reply(200).
		JSON(domain.User{ID: 1, Email: "updateduser@example.com"})

	router := gin.Default()
	router.PUT("/users/:id", handler.Update)

	body, _ := json.Marshal(updatedUser)
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %v; got %v", http.StatusOK, rr.Code)
	}

	var res domain.User
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expected := domain.User{ID: 1, Email: "updateduser@example.com"}
	if res.ID != expected.ID || res.Email != expected.Email {
		t.Fatalf("expected %v; got %v", expected, res)
	}
}

func TestDelete(t *testing.T) {
	defer gock.Off()

	gock.New("https://reqres.in").
		Delete("/api/users/1").
		Reply(204)

	router := gin.Default()
	router.DELETE("/users/:id", handler.Delete)

	req, _ := http.NewRequest("DELETE", "/users/1", nil)

	
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %v; got %v", http.StatusOK, rr.Code)
	}

	expected := `{"message":"User deleted successfully"}`
	if rr.Body.String() != expected {
		t.Fatalf("expected %v; got %v", expected, rr.Body.String())
	}
}
