package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Conf struct {
	Workspace string `json:"workspace"`
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
