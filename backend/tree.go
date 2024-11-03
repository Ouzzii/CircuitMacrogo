package backend

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ncruces/zenity"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) AskDirectory() string {
	a.CheckWorkspace()
	dir, err := zenity.SelectFile(
		zenity.Filename(``),
		zenity.Directory(),
	)
	if err != nil {
		LogWithDetails(err.Error())
	}
	conf := ReadConf()
	conf.Workspace = strings.ReplaceAll(dir, "\\", "/")
	conf.WriteConf()
	return strings.ReplaceAll(dir, "\\", "/")
}

func (a *App) GetDirectory(path string) []string {
	LogWithDetails("Info - GetDirectory fonskiyonu çalıştı")
	var files []string
	root := path

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
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

func (a *App) CheckWorkspace() string {
	conf := ReadConf()
	exist, err := exists(conf.Workspace)
	if err != nil {
		LogWithDetails(fmt.Sprint("Error - Error checking workspace: %v", err))
		return ""
	}
	if !exist {
		conf.Workspace = ""
		conf.WriteConf()
	}
	return conf.Workspace
}
func (a *App) CloseConfWorkspace() {
	conf := ReadConf()
	conf.Workspace = ""
	conf.WriteConf()
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
