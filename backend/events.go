package backend

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) Startup(ctx context.Context) {
	fmt.Println("App started!")

	runtime.EventsOn(ctx, "RunCheckDirectory", func(data ...interface{}) {
		if dirInfo, ok := data[0].(map[string]interface{}); ok {
			files := dirInfo["files"].([]interface{})
			directories := dirInfo["directories"].([]interface{})

			// Dosya ve klasör kontrolü
			result := checkFilesAndDirectories(files, directories)

			// Sonucu frontend'e gönder
			runtime.EventsEmit(ctx, "DirectoryCheckResult", result)
		}

	})

	/*runtime.EventsOn(ctx, "CheckDirectoryOnChange", func(data ...interface{}) {
		fmt.Println("CheckDirectoryOnChange received from frontend:", data)

		// Burada gerekli işlemleri yapabilirsiniz
		// Örneğin, belirli bir dizindeki dosya ve klasörleri kontrol etme
	})

	// Frontend'den gelen dosya ve dizinleri kontrol et
	runtime.EventsOn(ctx, "VerifyDirectories", func(data ...interface{}) {
		if dirInfo, ok := data[0].(map[string]interface{}); ok {
			files := dirInfo["files"].([]interface{})
			directories := dirInfo["directories"].([]interface{})

			// Dosya ve klasör kontrolü
			result := checkFilesAndDirectories(files, directories)
			// Sonucu frontend'e gönder
			runtime.EventsEmit(ctx, "DirectoryCheckResult", result)
		}
	})*/

	a.ctx = ctx
}

func checkFilesAndDirectories(files []interface{}, directories []interface{}) bool {
	conf := ReadConf()
	localDir := conf.Workspace

	localFiles := make([]string, 0)
	localDirectories := make([]string, 0)

	// Yerel dosyaları ve dizinleri oku
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
		fmt.Println("Error walking the path:", err)
		return false
	}

	// Sıralama
	sort.Strings(localFiles)
	sort.Strings(localDirectories)

	// Kontrol
	filesMatch := compareSlices(files, localFiles)
	dirsMatch := compareSlices(directories, localDirectories)
	return (filesMatch && dirsMatch)
}

// İki dilim arasındaki eşleşmeyi kontrol eden fonksiyon
func compareSlices(slice1 []interface{}, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice2 {
		if v != slice1[i].(string) {
			return false
		}
	}
	return true
}
