package backend

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func (a *App) GetPDF(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		LogWithDetails(fmt.Sprintf("Error - error encountered while reading pdf file: %v", err))
		return ""
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}
