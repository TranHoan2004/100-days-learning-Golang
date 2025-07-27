package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Dieu kien can va du de 1 phuong thuc tro thanh 1 endpoint Rest
// Tham so truyen vao phai co ResponseWriter va Request
// Phuong thuc duoc dang ky endpoint bang HandleFunc

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

// GET /books
func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Print("getBooks")
	w.Header().Set("Content-Type", "application/json")

	// ghi log kiem tra du lieu json
	data, err1 := json.Marshal(books)
	if err1 != nil { // nil la gia tri khong ton tai hoac chua khoi tao
		log.Printf("Loi: %s", err1.Error())
		return
	}
	log.Print(string(data))

	err := json.NewEncoder(w).Encode(books) // ghi thang du lieu vao Response roi gui ve client
	if err != nil {
		return
	}
	// ResponseWriter giong nhu 1 doan tau dang chuan bi roi ga. Viec cua minh la chat hang hoa (data) len cac toa
	// Mot khi "con tau" roi ben thi se khong the bat no dung lai hoac lay lai hang
}

// GET /books/{id}
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// _ co nghia la bo qua gia tri tai vi tri do. Trong truong hop ben duoi thi la index
	for _, book := range books {
		if book.ID == params["id"] {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				return
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte(`{"error" : "Book not found"}`))
	if err != nil {
		return
	}
}

// POST /books
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		return
	}
	book.ID = strconv.Itoa(rand.Intn(100000)) // tao id random
	books = append(books, book)
	err1 := json.NewEncoder(w).Encode(book)
	if err1 != nil {
		return
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			var updated Book
			err := json.NewDecoder(r.Body).Decode(&updated)
			if err != nil {
				return
			}
			updated.ID = params["id"]
			books = append(books, updated)
			err1 := json.NewEncoder(w).Encode(updated)
			if err1 != nil {
				return
			}
		}
	}
}

// DELETE /books/{id}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte(`{"error":"Book not found"}`))
	if err != nil {
		return
	}
}

func main() {
	// Mock dữ liệu
	books = append(books, Book{ID: "1", Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: "2", Title: "Clean Code", Author: "Robert C. Martin"})

	r := mux.NewRouter()

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server is running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
