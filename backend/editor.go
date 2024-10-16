package backend

import (
	"fmt"
	"os"
)

func (a *App) GetContent(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Getting Content: ", path)
		fmt.Println("Error getting file content:", err)
	}

	str := string(b)

	return str
}
