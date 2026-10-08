package latihanapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// Todo merepresentasikan struktur data untuk satu item todo
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

// Penyimpanan sementara di memory (In-Memory slice)
var (
	todos = []Todo{
		{ID: 1, Title: "Belajar Go dasar", IsDone: true},
		{ID: 2, Title: "Membuat REST API sederhana", IsDone: false},
	}
	nextID = 3
	mu     sync.Mutex // Mutex untuk mencegah data race saat banyak request masuk bersamaan
)

// GET /todos - Mengambil semua todo
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mu.Lock()
	defer mu.Unlock()

	json.NewEncoder(w).Encode(todos)
}

// GET /todos/{id} - Mengambil detail satu todo berdasarkan ID
func getTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "ID harus berupa angka"}`, http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, t := range todos {
		if t.ID == id {
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	http.Error(w, `{"error": "Todo tidak ditemukan"}`, http.StatusNotFound)
}

// POST /todos - Menambahkan todo baru
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, `{"error": "Format JSON tidak valid"}`, http.StatusBadRequest)
		return
	}

	if newTodo.Title == "" {
		http.Error(w, `{"error": "Title tidak boleh kosong"}`, http.StatusBadRequest)
		return
	}

	mu.Lock()
	newTodo.ID = nextID
	nextID++
	todos = append(todos, newTodo)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// PUT /todos/{id} - Mengubah/memperbarui todo berdasarkan ID
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "ID harus berupa angka"}`, http.StatusBadRequest)
		return
	}

	var updatedData Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, `{"error": "Format JSON tidak valid"}`, http.StatusBadRequest)
		return
	}

	if updatedData.Title == "" {
		http.Error(w, `{"error": "Title tidak boleh kosong"}`, http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = updatedData.Title
			todos[i].IsDone = updatedData.IsDone
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}

	http.Error(w, `{"error": "Todo tidak ditemukan"}`, http.StatusNotFound)
}

// DELETE /todos/{id} - Menghapus todo berdasarkan ID
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "ID harus berupa angka"}`, http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, t := range todos {
		if t.ID == id {
			// Menghapus elemen ke-i dari slice
			todos = append(todos[:i], todos[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Todo berhasil dihapus",
			})
			return
		}
	}

	http.Error(w, `{"error": "Todo tidak ditemukan"}`, http.StatusNotFound)
}

// StartServer mendaftarkan rute dan menjalankan HTTP server
func StartServer() {
	mux := http.NewServeMux()

	// Routing bawaan Go standard library
	mux.HandleFunc("GET /todos", getTodos)
	mux.HandleFunc("GET /todos/{id}", getTodoByID)
	mux.HandleFunc("POST /todos", createTodo)
	mux.HandleFunc("PUT /todos/{id}", updateTodo)
	mux.HandleFunc("DELETE /todos/{id}", deleteTodo)

	port := ":8080"
	fmt.Println("🚀 Server TODO API berjalan di http://localhost" + port)
	fmt.Println("Tekan Ctrl+C untuk menghentikan server.")

	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Printf("Gagal menjalankan server: %s\n", err)
	}
}
