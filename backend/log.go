package backend

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	Runtime "runtime"
	"time"
)

var logFile *os.File

// Log setup fonksiyonu
func SetupLogger() error {
	var err error
	now := time.Now()
	logFile, err = os.OpenFile(fmt.Sprintf("Log-%v.txt", now.Format("02-01-2006")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	log.SetOutput(logFile)
	LogWithDetails("Log sistemi başlatıldı.")
	return nil
}

// Log fonksiyonu
func LogWithDetails(args ...string) {
	_, file, line, ok := Runtime.Caller(1)
	if ok {
		// Dosya adını al
		fileName := filepath.Base(file)

		// Mesajı oluşturalım
		var message string
		if len(args) == 1 {
			message = args[0]
		} else if len(args) == 2 {
			message = args[0] + " - " + args[1]
		} else {
			log.Println("Invalid number of arguments")
			return
		}

		// Log formatını ayarlıyoruz
		log.Printf("- %s:%d - %s", fileName, line, message)
	} else {
		LogWithDetails(args...)
	}
}
