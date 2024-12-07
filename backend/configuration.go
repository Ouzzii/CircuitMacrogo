package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func ReadConf() Conf {
	plan, _ := ioutil.ReadFile("configuration.json")
	var data Conf
	if err := json.Unmarshal(plan, &data); err != nil {
		LogWithDetails(fmt.Sprintf("Veriyi çözümlerken hata oluştu: %v", err))
	}
	return data
}
func (c *Conf) WriteConf() {
	if c.PdflatexPaths == nil {
		c.PdflatexPaths = make(map[string]string)
	}
	jsonData, err := json.MarshalIndent(c, "", "  ") // Güzel bir format için MarshalIndent kullanıyoruz
	if err != nil {
		LogWithDetails(fmt.Sprintf("Error - JSON formatına çevirme hatası: %v", err))
	}

	// JSON verisini bir dosyaya yaz
	err = ioutil.WriteFile("configuration.json", jsonData, os.ModePerm)
	if err != nil {
		LogWithDetails(fmt.Sprintf("Error - Dosya yazma hatası: %v", err))
	}
}
func (c *Conf) AddPdflatexPath(key, path string) {
	if c.PdflatexPaths == nil {
		c.PdflatexPaths = make(map[string]string)
	}
	c.PdflatexPaths[key] = path
}

func CheckInternet() bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	_, err := client.Get("http://www.google.com")
	return err == nil
}
