package backend

import (
	"log"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
)

/*func init() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Log dosyası açılamadı: %v", err)
	}

	// Log seviyelerini tanımla
	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	defer file.Close()

	// Log'un çıktısını dosyaya yönlendir
	log.SetOutput(file)

	// Log formatı
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Log(level, message string) {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Log dosyası açılamadı: %v", err)
	}

	// Log seviyelerini tanımla
	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(file, "DEBUG", log.Ldate|log.Ltime|log.Lshortfile)
	defer file.Close()

	log.SetOutput(file)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if strings.ToLower(level) == "info" {
		Info.Println(message)
	} else if strings.ToLower(level) == "debug" {

	} else if strings.ToLower(level) == "error" {
		Error.Println(message)
	} else if strings.ToLower(level) == "warning" {
		Warning.Println(message)
	}

}
*/
