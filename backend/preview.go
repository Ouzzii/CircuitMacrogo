package backend

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (a *App) GetPDF(w http.ResponseWriter, r *http.Request) {
	// İstek gövdesini oku
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var PostData struct {
		Path string `json:"path"`
	}
	if err := json.Unmarshal(body, &PostData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	data, err := ioutil.ReadFile(PostData.Path)
	if err != nil {
		LogWithDetails(fmt.Sprintf("Error - error encountered while reading pdf file: %v", err))
		http.Error(w, "Failed to read PDF file", http.StatusInternalServerError)
		return
	}

	base64Data := base64.StdEncoding.EncodeToString(data)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"base64": base64Data}); err != nil {
		http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
		return
	}
}
