package backend

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	Runtime "runtime"
	"sort"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var logFile *os.File

// Log setup fonksiyonu
func SetupLogger() error {
	var err error
	logFile, err = os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	log.SetOutput(logFile)
	//LogWithDetails("Log sistemi başlatıldı.")
	return nil
}

// Log fonksiyonu
func LogWithDetails(message string) {
	_, file, line, ok := Runtime.Caller(1) // 1, bu fonksiyonu çağıran fonksiyonun konumunu alır
	if ok {
		// Sadece dosya adı ve satır numarasını alıyoruz
		fileName := filepath.Base(file)                     // Dosya adını al
		log.Printf("- %s:%d - %s", fileName, line, message) // Log formatını ayarlıyoruz
	} else {
		LogWithDetails(message)
	}
}

// Cleanup fonksiyonu
func Cleanup() {
	if logFile != nil {
		logFile.Close()
		LogWithDetails("Log sistemi kapatıldı.")
	}
}

func (a *App) Startup(ctx context.Context) {

	a.configuration = ReadConf()

	err := SetupLogger()
	if err != nil {
		fmt.Println("Log dosyası oluşturulamadı:", err)
		return
	}

	LogWithDetails("Startup")
	runtime.EventsOn(ctx, "RunCheckDirectory", func(data ...interface{}) {

		if dirInfo, ok := data[0].(map[string]interface{}); ok {
			files := dirInfo["files"].([]interface{})
			directories := dirInfo["directories"].([]interface{})

			// Dosya ve klasör kontrolü
			result := a.checkFilesAndDirectories(files, directories)

			// Sonucu frontend'e gönder
			runtime.EventsEmit(ctx, "DirectoryCheckResult", result)
		}

	})

	a.ctx = ctx
}

func (a *App) checkFilesAndDirectories(files []interface{}, directories []interface{}) bool {

	localDir := a.configuration.Workspace

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

	// Sıralama
	sort.Strings(localFiles)
	sort.Strings(localDirectories)

	// Kontrol
	filesMatch := compareSlices(files, localFiles)
	dirsMatch := compareSlices(directories, localDirectories)
	fmt.Println(filesMatch, dirsMatch, "events")
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
