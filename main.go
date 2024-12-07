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

	// Ana sayfa handler
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func h1(w http.ResponseWriter, r *http.Request) {
	temPlate := template.Must(template.ParseFiles("./templates/index.html"))
	temPlate.Execute(w, nil)
}
