package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // İstekleri doğrulama (geliştirme ortamında kullanılır)
	},
}

func main() {
	// Statik dosyaları serve et
	fs := http.FileServer(http.Dir("./style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	// WebSocket handler
	http.HandleFunc("/ws", handleWebSocket)

	// Ana sayfa handler
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func h1(w http.ResponseWriter, r *http.Request) {
	temPlate := template.Must(template.ParseFiles("./templates/index.html"))
	temPlate.Execute(w, nil)
}

// WebSocket bağlantısı için handler
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Bağlantı Hatası:", err)
		return
	}
	defer ws.Close()

	log.Println("WebSocket bağlantısı kuruldu.")

	for {

		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Mesaj Okuma Hatası:", err)
			break
		}

		log.Printf("Alınan Mesaj: %s\n", msg)
		time.Sleep(time.Second * 10)

		// Gelen mesajı tekrar istemciye gönder
		err = ws.WriteJSON(fmt.Sprintf("Geri Gönderilen: %s", msg))
		if err != nil {
			log.Println("Mesaj Gönderme Hatası:", err)
			break
		}
	}
}
