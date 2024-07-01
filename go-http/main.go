package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("my-header", "contoh header")

		q := r.URL.Query().Get("q")
		w.Header().Set("q", q)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hai cuy"))
	})

	http.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("my-header", "contoh header"))
		byteBody, err := io.ReadAll(r.Body)

		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "error", http.StatusInternalServerError)
		}

		var m map[string]any

		if err := json.Unmarshal(byteBody, &m); err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "error", http.StatusInternalServerError)
		}

		fmt.Println(m)
		w.Write([]byte(byteBody))
	})

	http.ListenAndServe(":8081", nil)
}
