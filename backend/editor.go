package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (a *App) GetContent(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	b, err := os.ReadFile(path)
	if err != nil {
		LogWithDetails(fmt.Sprintf("Error - Error getting file content: %v", err))
	}

	w.Write(b)
}

type SaveContentPost struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func (a *App) SaveContent(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var PostData SaveContentPost
	if err := json.Unmarshal(body, &PostData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(PostData)

	err = os.WriteFile(PostData.Path, []byte(PostData.Content), 0777)
	if err != nil {
		LogWithDetails(fmt.Sprintf("Error - Error when %v file saving: %v", PostData.Path, err))
		w.Write([]byte("nok"))
	}
	w.Write([]byte("ok"))
}
