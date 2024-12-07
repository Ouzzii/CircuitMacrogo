package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Tüm istekleri kabul etmek için true döndür
		return true
	},
}

func (a *App) RunCheckDirectory(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket connection established!")

	// WebSocket üzerinden veri alışverişi
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		var res [][]string

		// Gelen mesajı çözümle
		err = json.Unmarshal(message, &res)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			break
		}
		checked := a.checkFilesAndDirectories(res[0], res[1])
		checkedJson, err := json.Marshal(checked)
		if err != nil {
			fmt.Println("Error marshalling response:", err)
			break
		}
		err = conn.WriteMessage(websocket.TextMessage, checkedJson)

		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}

}

func (a *App) checkFilesAndDirectories(files []string, directories []string) bool {

	localDir := a.Configuration.Workspace

	localFiles := make([]string, 0)
	localDirectories := make([]string, 0)

	err := filepath.Walk(localDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.ReplaceAll(path, "\\", "/") != localDir {
			localDirectories = append(localDirectories, strings.ReplaceAll(path, "\\", "/"))
		} else if !info.IsDir() {
			localFiles = append(localFiles, strings.ReplaceAll(path, "\\", "/"))
		}
		return nil
	})

	if err != nil {
		LogWithDetails(fmt.Sprintf("Error walking the path: %v", err.Error()))
		return false
	}

	sort.Strings(localFiles)
	sort.Strings(localDirectories)

	filesMatch := compareSlices(files, localFiles)
	dirsMatch := compareSlices(directories, localDirectories)
	return (filesMatch && dirsMatch)
}

func compareSlices(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice2 {
		if v != slice1[i] {
			return false
		}
	}
	return true
}
