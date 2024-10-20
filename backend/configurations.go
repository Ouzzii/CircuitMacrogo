package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Conf struct {
	Workspace     string            `json:"workspace"`
	PdflatexPaths map[string]string `json:"pdflatex-paths"`
	LastDistro    string            `json:"last-distro"`
}

func ReadConf() Conf {
	plan, _ := ioutil.ReadFile("configuration.json")
	var data Conf
	if err := json.Unmarshal(plan, &data); err != nil {
		log.Printf("Veriyi çözümlerken hata oluştu: %v", err)
	}
	return data
}
func (c *Conf) WriteConf() {
	jsonData, err := json.MarshalIndent(c, "", "  ") // Güzel bir format için MarshalIndent kullanıyoruz
	if err != nil {
		fmt.Println("JSON formatına çevirme hatası:", err)
	}

	// JSON verisini bir dosyaya yaz
	err = ioutil.WriteFile("configuration.json", jsonData, os.ModePerm)
	if err != nil {
		fmt.Println("Dosya yazma hatası:", err)
	}
}
func (c *Conf) AddPdflatexPath(key, path string) {
	//if c.PdflatexPaths == nil {
	//	c.PdflatexPaths = make(map[string]string)
	//}
	c.PdflatexPaths[key] = path
}

func CheckInternet() bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	_, err := client.Get("http://www.google.com")
	return err == nil
}
