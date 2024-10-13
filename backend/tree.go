package backend

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ncruces/zenity"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AskDirectory() string {
	dir, err := zenity.SelectFile(
		zenity.Filename(``),
		zenity.Directory(),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a.GetDirectory(dir))
	return strings.ReplaceAll(dir, "\\", "/")
}

func (a *App) GetDirectory(path string) []string {
	var files []string
	root := path // Klasörün yolunu buraya gir

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Dosya veya klasörün yolunu slice'a ekle
		files = append(files, strings.ReplaceAll(path, "\\", "/"))
		return nil
	})

	if err != nil {
		fmt.Printf("Hata: %v\n", err)
	}

	return files
}
func (a *App) IsFile(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		return false
	case mode.IsRegular():
		return true
	}
	return false
}
