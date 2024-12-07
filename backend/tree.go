package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ncruces/zenity"
)

func (a *App) AskDirectory(w http.ResponseWriter, r *http.Request) {

	dir, err := zenity.SelectFile(
		zenity.Filename(``),
		zenity.Directory(),
	)
	if err != nil {
		LogWithDetails(err.Error())
	}

	a.Configuration.Workspace = strings.ReplaceAll(dir, "\\", "/")
	a.Configuration.WriteConf()

	json.NewEncoder(w).Encode(a.Configuration.Workspace)

}

func (a *App) GetDirectory(w http.ResponseWriter, r *http.Request) {

	var path string
	err := json.NewDecoder(r.Body).Decode(&path)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	LogWithDetails("Info - GetDirectory fonskiyonu çalıştı")
	var files []string
	var filetypes []bool
	root := path

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			filetypes = append(filetypes, false)
		} else {
			filetypes = append(filetypes, true)
		}

		files = append(files, strings.ReplaceAll(path, "\\", "/"))
		return nil
	})

	if err != nil {
		fmt.Printf("Hata: %v\n", err)
	}

	result := map[string]interface{}{
		"files":     files,
		"filetypes": filetypes,
	}

	json.NewEncoder(w).Encode(result)
}
func (a *App) IsFile(w http.ResponseWriter, r *http.Request) {
	var response bool
	var path string
	err := json.NewDecoder(r.Body).Decode(&path)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		response = false
	case mode.IsRegular():
		response = true
	}

	json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
}

func (a *App) CheckWorkspace(w http.ResponseWriter, r *http.Request) {

	exist, err := exists(a.Configuration.Workspace)
	if err != nil {
		LogWithDetails(fmt.Sprintf("Error - Error checking workspace: %v", err))
		return
	}
	if !exist {
		a.Configuration.Workspace = ""
		a.Configuration.WriteConf()
	}
	json.NewEncoder(w).Encode(a.Configuration.Workspace)

}
func (a *App) CloseConfWorkspace() {
	a.Configuration.Workspace = ""
	a.Configuration.WriteConf()
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
