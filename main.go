package main

import (
	"html/template"
	"log"
	"net/http"
	"webserver/backend"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // İstekleri doğrulama (geliştirme ortamında kullanılır)
	},
}

func init() {
	backend.UpdateEnv()
}

func main() {
	// Statik dosyaları serve et

	app := backend.App{
		Configuration: backend.Conf{},
	}

	app.Configuration = backend.ReadConf()

	styleDir := http.FileServer(http.Dir("./style"))
	scriptDir := http.FileServer(http.Dir("./scripts"))
	http.Handle("/style/", http.StripPrefix("/style/", styleDir))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", scriptDir))

	// WebSocket handler
	http.HandleFunc("/CheckWorkspace", app.CheckWorkspace)
	http.HandleFunc("/AskDirectory", app.AskDirectory)
	http.HandleFunc("/GetDirectory", app.GetDirectory)
	http.HandleFunc("/IsFile", app.IsFile)
	http.HandleFunc("/RunDirectoryCheck", app.RunCheckDirectory)
	http.HandleFunc("/CloseWorkspace", app.CloseConfWorkspace)

	http.HandleFunc("/GetContent", app.GetContent)
	http.HandleFunc("/SaveContent", app.SaveContent)

	http.HandleFunc("/GetPDF", app.GetPDF)

	http.HandleFunc("/DetectTexDistros", app.Detect_tex_distros)
	http.HandleFunc("/BoxdimsIsInstalled", app.Boxdims_is_installed)
	http.HandleFunc("/ChooseDistro", app.ChooseDistro)

	http.HandleFunc("/Compile", app.Compile)

	// Ana sayfa handler
	http.HandleFunc("/", HomePage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	temPlate := template.Must(template.ParseFiles("./templates/index.html"))
	temPlate.Execute(w, nil)
}
