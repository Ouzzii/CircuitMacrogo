package backend

import (
	"encoding/base64"
	"io/ioutil"
	"log"
)

func (a *App) GetPDF(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("error encountered while reading pdf file:", err)
		return ""
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}
