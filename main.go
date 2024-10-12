package main

import (
	"api/facts"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const PORT = ":8080"

func getAllFacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := facts.GetFacts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseMessage{Message: "Error getting facts"})
		return
	}

	w.WriteHeader(http.StatusOK)
	response := make(map[string]interface{})
	response["data"] = data
	json.NewEncoder(w).Encode(response)
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func getFacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idValue := r.PathValue("id")
	id, err := strconv.Atoi(idValue)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseMessage{Message: "Error getting ID from path"})
		return
	}

	data, err := facts.GetFactByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseMessage{Message: "Error getting facts"})
		return
	}

	w.WriteHeader(http.StatusOK)
	response := make(map[string]interface{})
	response["data"] = data
	json.NewEncoder(w).Encode(response)
}

func checkAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		token := query.Get("token")
		if token != "123" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	err := facts.LoadFacts()
	if err != nil {
		fmt.Println("error loading facts", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /{$}", checkAuth(http.HandlerFunc(getAllFacts)))
	mux.Handle("GET /{id}", checkAuth(http.HandlerFunc(getFacts)))

	fmt.Println("Server is running on 8080")
	http.ListenAndServe(PORT, mux)
}
