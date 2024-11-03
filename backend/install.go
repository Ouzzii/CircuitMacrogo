package backend

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func UpdateEnv() {
	// Yeni eklemek istediğiniz path
	cmPath := os.Getenv("HOME") + "/CMEditor/Circuit_macros"
	if !Exists(cmPath) {
		InstallCM()
		if runtime.GOOS == "windows" {
			if !Exists(cmPath + "/m4.exe") {
				Download(cmPath+"/m4.exe", "https://ece.uwaterloo.ca/~aplevich/dpic/Windows/m4.exe")
			}
			if !Exists(cmPath + "/dpic.exe") {
				Download(cmPath+"/dpic.exe", "https://ece.uwaterloo.ca/~aplevich/dpic/Windows/dpic.exe")
			}

		}
	}

	currentPath := os.Getenv("PATH")

	updatedPath := fmt.Sprintf("%s%c%s", currentPath, filepath.ListSeparator, cmPath)

	err := os.Setenv("PATH", updatedPath)
	if err != nil {
		LogWithDetails(fmt.Sprintf("Error - PATH ayarlanamadı: %v", err))
		return
	}
	LogWithDetails(fmt.Sprintf("Success - Güncellenmiş PATH: %v", os.Getenv("PATH")))
}

func InstallCM() {
	CMEditorPath := fmt.Sprintf("%v/CMEditor", os.Getenv("HOME"))
	if !Exists(CMEditorPath) {
		err := os.MkdirAll(CMEditorPath, 0755)
		if err != nil {
			LogWithDetails(fmt.Sprintf("Error - Klasör oluşturulamadı: %v", err))
			return
		}
		if err := Download(CMEditorPath+"/Circuit_macros.zip", "https://gitlab.com/aplevich/circuit_macros/-/archive/master/circuit_macros-master.zip"); err != nil {
			LogWithDetails(fmt.Sprintf("Error - İndirilirken hata ile karşılaşıldı: %v", err))
		}

		zipFile := CMEditorPath + "/Circuit_macros.zip"
		outputDir := CMEditorPath + "/Circuit_macros"
		err = Unzip(zipFile, outputDir)
		if err != nil {
			LogWithDetails(fmt.Sprintf("Error - Dosya çıkartılamadı: %v", err))
			return
		}
		os.Remove(CMEditorPath + "/Circuit_macros.zip")
		LogWithDetails(fmt.Sprintf("Info - Dosya başarıyla çıkartıldı: %v", outputDir))

	}

}

func Exists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else {
			return false
		}
	}
	return true
}

func GetCMPath() string {
	return ""
}

func Download(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, file := range r.File {
		filePath := filepath.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return err
		}

		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}
	}
	return nil
}
