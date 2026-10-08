package latihanapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoAPI(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", getTodos)
	mux.HandleFunc("GET /todos/{id}", getTodoByID)
	mux.HandleFunc("POST /todos", createTodo)
	mux.HandleFunc("PUT /todos/{id}", updateTodo)
	mux.HandleFunc("DELETE /todos/{id}", deleteTodo)

	// Test 1: GET /todos
	t.Run("GET /todos", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status code salah: dapat %d, mau %d", rec.Code, http.StatusOK)
		}
	})

	// Test 2: POST /todos
	var createdTodo Todo
	t.Run("POST /todos", func(t *testing.T) {
		body := `{"title": "Belajar Unit Test Go", "is_done": false}`
		req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(body))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Fatalf("status code salah: dapat %d, mau %d", rec.Code, http.StatusCreated)
		}

		if err := json.NewDecoder(rec.Body).Decode(&createdTodo); err != nil {
			t.Fatalf("gagal decode response json: %v", err)
		}

		if createdTodo.Title != "Belajar Unit Test Go" {
			t.Errorf("title salah: dapat %s", createdTodo.Title)
		}
	})

	// Test 3: GET /todos/{id}
	t.Run("GET /todos/{id}", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status code salah: dapat %d, mau %d", rec.Code, http.StatusOK)
		}
	})

	// Test 4: PUT /todos/{id}
	t.Run("PUT /todos/{id}", func(t *testing.T) {
		body := `{"title": "Belajar Unit Test Go (Updated)", "is_done": true}`
		req := httptest.NewRequest(http.MethodPut, "/todos/1", bytes.NewBufferString(body))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status code salah: dapat %d, mau %d", rec.Code, http.StatusOK)
		}

		var updated Todo
		if err := json.NewDecoder(rec.Body).Decode(&updated); err != nil {
			t.Fatalf("gagal decode response json: %v", err)
		}

		if updated.Title != "Belajar Unit Test Go (Updated)" || !updated.IsDone {
			t.Errorf("data tidak sesuai update: %+v", updated)
		}
	})

	// Test 5: DELETE /todos/{id}
	t.Run("DELETE /todos/{id}", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status code salah: dapat %d, mau %d", rec.Code, http.StatusOK)
		}
	})

	// Test 6: GET deleted todo should return 404
	t.Run("GET non-existing /todos/{id}", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Fatalf("status code salah: dapat %d, mau %d", rec.Code, http.StatusNotFound)
		}
	})
}
