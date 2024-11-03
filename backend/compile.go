package backend

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func tolatex(m4 string) (string, string) {
	// Girdi dosyaları
	pgfFile := "pgf.m4"
	m4Input := m4
	outputFile := strings.Replace(m4, filepath.Ext(m4), ".tex", 1)

	m4Cmd := exec.Command("m4", pgfFile, m4Input)

	dpicCmd := exec.Command("dpic", "-g")

	dpicCmd.Stdin, _ = m4Cmd.StdoutPipe()

	// dpic komutunun çıktısını dosyaya yönlendiriyoruz
	output, err := os.Create(outputFile)
	if err != nil {
		LogWithDetails(fmt.Sprintf("Error - Çıktı dosyasını oluşturma hatası: %v", err))
		return "", fmt.Sprintf("Çıktı dosyasını oluşturma hatası: %v", err)
	}
	defer output.Close()
	dpicCmd.Stdout = output

	// Her iki komutu çalıştırıyoruz
	if err := m4Cmd.Start(); err != nil {
		LogWithDetails(fmt.Sprintf("Error - m4 komutunu başlatma hatası: %v", err))
		return "", fmt.Sprintf("m4 komutunu başlatma hatası: %v", err)
	}

	if err := dpicCmd.Start(); err != nil {
		LogWithDetails(fmt.Sprintf("Error - dpic komutunu başlatma hatası: %v", err))
		return "", fmt.Sprintf("dpic komutunu başlatma hatası: %v", err)
	}

	// Her iki komutun bitmesini bekliyoruz
	if err := m4Cmd.Wait(); err != nil {
		LogWithDetails(fmt.Sprintf("Error - m4 komutunun beklenmesi sırasında hata: %v", err))
		return "", fmt.Sprintf("m4 komutunun beklenmesi sırasında hata: %v", err)
	}

	if err := dpicCmd.Wait(); err != nil {
		LogWithDetails(fmt.Sprintf("Error - dpic komutunun beklenmesi sırasında hata: %v", err))
		return "", fmt.Sprintf("dpic komutunun beklenmesi sırasında hata: %v", err)
	}
	LogWithDetails(fmt.Sprintf("Info - Latex'e derleme işlemi başarıyla gerçekleşti: %v", m4))
	return outputFile, ""
}

func toPDF(latex string) (bool, string) {
	conf := ReadConf()
	pdfCmd := exec.Command(conf.PdflatexPaths[conf.LastDistro], fmt.Sprintf(`-output-directory=%v`, conf.Workspace), fmt.Sprintf(`-aux-directory=%v`, conf.Workspace), "-interaction=nonstopmode", latex)

	if err := pdfCmd.Start(); err != nil {
		LogWithDetails(fmt.Sprintf("Error - pdflatex komutunu başlatma hatası: %v", err))
		fmt.Println(err)
		return false, fmt.Sprintf("pdflatex komutunu başlatma hatası: %v", err)
	}
	if err := pdfCmd.Wait(); err != nil {
		fmt.Println(err)
		LogWithDetails(fmt.Sprintf("Error - pdflatex komutunun beklenmesi sırasında hata: %v", err))
		return false, fmt.Sprintf("pdflatex komutunun beklenmesi sırasında hata: %v", err)
	}

	LogWithDetails(fmt.Sprintf("Info - pdflatex ile başarıyla dosya derlendi, %v", latex))
	return true, ""
}

func (a *App) Compile(target, path string) string {

	if target == "latex" {
		_, err := tolatex(path)
		if err != "" {
			return err
		}
		return ""
	} else if target == "pdf" {

		tltx, err := tolatex(path)
		if err != "" {
			return err
		}

		_, err = toPDF(tltx)
		return err
	}
	return ""
}
