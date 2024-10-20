package backend

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func tolatex(m4 string) {
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
		fmt.Fprintf(os.Stderr, "Çıktı dosyasını oluşturma hatası: %v\n", err)
		return
	}
	defer output.Close()
	dpicCmd.Stdout = output

	// Her iki komutu çalıştırıyoruz
	if err := m4Cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "m4 komutunu başlatma hatası: %v\n", err)
		return
	}

	if err := dpicCmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "dpic komutunu başlatma hatası: %v\n", err)
		return
	}

	// Her iki komutun bitmesini bekliyoruz
	if err := m4Cmd.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "m4 komutunun beklenmesi sırasında hata: %v\n", err)
		return
	}

	if err := dpicCmd.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "dpic komutunun beklenmesi sırasında hata: %v\n", err)
		return
	}

	fmt.Println("İşlem başarıyla tamamlandı.")
	toPDF(outputFile)
}

func toPDF(latex string) {

}
