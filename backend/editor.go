package backend

import (
	"fmt"
	"os"
)

func (a *App) GetContent(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		Log("Error", fmt.Sprintf("Error getting file content: %v", err))
	}

	str := string(b)

	return str
}
func (a *App) SaveContent(path, content string) bool {
	data := []byte(content)
	err := os.WriteFile(path, data, 0777)
	if err != nil {
		Log("Error", fmt.Sprintf("Error when %v file saving: %v", path, err))
		return false
	}
	return true
}
