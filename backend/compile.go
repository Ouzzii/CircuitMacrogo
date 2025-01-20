package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func tolatex(m4 string) (string, string) {
	pgfFile := "pgf.m4"
	m4Input := m4
	outputFile := strings.Replace(m4, filepath.Ext(m4), ".tex", 1)

	var m4out, dpicout bytes.Buffer

	var m4Cmd, dpicCmd *exec.Cmd
	if runtime.GOOS == "linux" {
		m4Cmd = exec.Command("m4", pgfFile, m4Input)
		dpicCmd = exec.Command("dpic", "-g")
	} else if runtime.GOOS == "windows" {
		m4Cmd = exec.Command("cmd", "/C", "chdir", os.Getenv("USERPROFILE")+"\\CMEditor\\circuit_macros-master", "&&", "m4", pgfFile, m4Input)
		dpicCmd = exec.Command("cmd", "/C", "chdir", os.Getenv("USERPROFILE")+"\\CMEditor\\circuit_macros-master", "&&", "dpic.exe", "-g")
	}

	m4Cmd.Stdout = &m4out
	dpicCmd.Stdout = &dpicout
	dpicCmd.Stdin = &m4out

	// Her iki komutu çalıştırıyoruz
	if err := m4Cmd.Start(); err != nil {
		LogWithDetails(fmt.Sprintf("Error - m4 komutunu başlatma hatası: %v", err))
		return "", fmt.Sprintf("m4 komutunu başlatma hatası: %v", err)
	}
	m4Cmd.Wait()

	if err := dpicCmd.Start(); err != nil {
		LogWithDetails(fmt.Sprintf("Error - dpic komutunu başlatma hatası: %v", err))
		return "", fmt.Sprintf("dpic komutunu başlatma hatası: %v", err)
	}
	dpicCmd.Wait()

	ioutil.WriteFile(outputFile, dpicout.Bytes(), 0644)

	LogWithDetails(fmt.Sprintf("Info - Latex'e derleme işlemi başarıyla gerçekleşti: %v", m4))
	return outputFile, ""
}

func (a *App) toPDF(latex string) (bool, string) {
	fmt.Println("compiling to pdf")
	var message string
	var stdout, stderr bytes.Buffer

	pdfCmd := exec.Command(a.Configuration.PdflatexPaths[a.Configuration.LastDistro], fmt.Sprintf(`-output-directory=%v`, a.Configuration.Workspace), fmt.Sprintf(`-aux-directory=%v`, a.Configuration.Workspace), "-interaction=nonstopmode", latex)

	pdfCmd.Stdout = &stdout
	pdfCmd.Stderr = &stderr

	if err := pdfCmd.Start(); err != nil {
		LogWithDetails(fmt.Sprintf("Error - pdflatex komutunu başlatma hatası: %v", err))
		fmt.Println(err)
		return false, fmt.Sprintf("pdflatex komutunu başlatma hatası: %v", err)
	}

	/*if err := pdfCmd.Wait(); err != nil {
		fmt.Println(err)
		LogWithDetails(fmt.Sprintf("Error - pdflatex komutunun beklenmesi sırasında hata: %v", err))
		return false, fmt.Sprintf("pdflatex komutunun beklenmesi sırasında hata: %v", err)
	}*/

	pdfCmd.Wait()
	output_message := string(stdout.Bytes())
	if strings.Contains(output_message, "Output written on") {
		message = "success"
	} else if strings.Contains(output_message, ".sty' not found") {
		lib := strings.Split(output_message, "! LaTeX Error: File `")[1]
		lib = strings.Split(lib, "'")[0]
		fmt.Println(lib, "Kutuphanesi bulunamadi")
		//
	}
	LogWithDetails(fmt.Sprintf("Info - pdflatex ile başarıyla dosya derlendi, %v", latex))
	return true, message
}

type CompileTarget struct {
	Target string `json:"target"`
	Path   string `json"path"`
}

func (a *App) Compile(w http.ResponseWriter, r *http.Request) {
	//target, path
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	var compiledata CompileTarget
	if err := json.Unmarshal(body, &compiledata); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if compiledata.Target == "latex" {
		_, err := tolatex(compiledata.Path)
		if err != "" {
			w.Write([]byte(err))
		}
		w.Write([]byte(""))
	} else if compiledata.Target == "pdf" {

		tltx, err := tolatex(compiledata.Path)
		if err != "" {
			w.Write([]byte(err))
		}

		_, err = a.toPDF(tltx)
		w.Write([]byte(err))
	}
	w.Write([]byte(""))
}
