package backend

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func (a *App) Detect_tex_distros() Conf {
	var conf Conf
	//last distro check

	if runtime.GOOS == "linux" {
		//miktex-pdflatex
		detectPdfCmd := exec.Command("sh", "-c", "find / -name miktex-pdflatex 2>&1 | grep -v 'Permission denied' | grep -v 'Invalid argument' | grep -v 'No such file or directory'")
		output, err := detectPdfCmd.CombinedOutput()
		if err != nil {
			Log("Error", fmt.Sprintf("miktex - detect pdflatex komutunu başlatma hatası: %v", err))
			return conf
		}
		conf = ReadConf()
		conf.AddPdflatexPath("miktex", strings.Replace(string(output), "\n", "", 1))
		conf.WriteConf()

		//pdflatex
		detectPdfCmd = exec.Command("sh", "-c", "which pdflatex")

		output, err = detectPdfCmd.CombinedOutput()
		if err != nil {
			Log("Error", fmt.Sprintf("texlive - detect pdflatex komutunu başlatma hatası: %v", err))
			return conf
		}

		for _, pdflatexpath := range strings.Split(string(output), "\n") {
			if pdflatexpath != "" {
				pdflatexDistroCmd := exec.Command("sh", "-c", fmt.Sprintf("%v --version", pdflatexpath))
				output, err = pdflatexDistroCmd.CombinedOutput()
				if err != nil {
					Log("Error", fmt.Sprintf("pdflatex versiyonu kontrol edilirken hata ile karşılaşıldı: %v\n", err))
					return conf
				}
				if strings.Contains(string(output), "TeX Live") {
					conf.AddPdflatexPath("texlive", pdflatexpath)
				}
			}
		}
		conf.WriteConf()

	} else if runtime.GOOS == "windows" {

	}
	return conf

}

func (a *App) Boxdims_is_installed() Boxdims {
	conf := ReadConf()
	boxdims := Boxdims{}
	if runtime.GOOS == "linux" {
		if _, exists := conf.PdflatexPaths["miktex"]; exists {
			miktexpath := filepath.Join(os.Getenv("HOME"), ".miktex", "texmfs", "install", "tex", "latex", "circuit_macros")
			if !Exists(miktexpath) {
				if CheckInternet() {
					Download(miktexpath+".zip", "https://mirrors.ctan.org/graphics/circuit_macros.zip")
					Unzip(miktexpath+".zip", filepath.Dir(miktexpath))
					//update database
					exec.Command("sh", "-c", fmt.Sprintf("%v --update-fndb", get_initexmf()))
					Log("Info", "Miktex: boxdims.sty başarılı bir şekilde indirildi")
					boxdims.Miktex = "exist"
				} else {
					Log("Error", "Miktex: internet bağlantısı olmadığından boxdims.sty indirilemedi")
					boxdims.Miktex = "not exist"
				}
			} else {
				Log("Info", "Miktex: boxdims.sty zaten kurulu")
				boxdims.Miktex = "exist"
			}
		} else {
			Log("Warning", "miktex: cihazda miktex bulunmamaktadır")
		}
		if _, exists := conf.PdflatexPaths["texlive"]; exists {
			texlivepathCmd := exec.Command("sh", "-c", `find / -wholename "*/texmf-dist/tex/latex" 2>&1 | grep -v "Permission denied" | grep -v "Invalid argument" | grep -v "No such file or directory"`)
			output, _ := texlivepathCmd.CombinedOutput()
			texlivePath := strings.Trim(string(output), "\n")
			if Exists(texlivePath) {
				commandbash := fmt.Sprintf(
					"wget -P %s https://mirrors.ctan.org/graphics/circuit_macros.zip && unzip %s -d %s && rm %s && mv %s %s && echo successful",
					texlivePath,
					filepath.Join(texlivePath, "circuit_macros.zip"),
					texlivePath,
					filepath.Join(texlivePath, "circuit_macros.zip"),
					filepath.Join(texlivePath, "circuit_macros"),
					filepath.Join(texlivePath, "boxdims"),
				)
				if !Exists(filepath.Join(texlivePath, "boxdims")) {
					if CheckInternet() {
						downloadTexliveCmd := exec.Command("pkexec", "bash", "-c", commandbash)
						output, _ := downloadTexliveCmd.CombinedOutput()
						result := strings.Trim(string(output), "\n")

						if strings.HasSuffix(result, "successful") {
							Log("Info", "Texlive için boxdims başarıyla indirildi")
							boxdims.Texlive = "exist"
						} else {
							Log("Error", fmt.Sprintf("Texlive için boxdims indirilirken hata ile karşılaşıldı: %v\n", result))
						}
					} else {
						Log("Error", "texlive: internet bağlantısı olmadığından boxdims.sty indirilemedi")
					}
				} else {
					Log("Info", "texlive: boxdims.sty zaten kurulu")
					boxdims.Texlive = "exist"
				}
			} else {
				Log("Error", "texlive: Dizin bulunamadı")
			}
		} else {
			Log("Warning", "texlive: cihazda texlive bulunmamaktadır")
		}

	} else if runtime.GOOS == "windows" {

	}

	return boxdims
}

type Boxdims struct {
	Miktex  string `json:"miktex"`
	Texlive string `json:"texlive"`
}

func get_initexmf() string {
	var initexmf string
	if runtime.GOOS == "linux" {
		InitexmfCmd := exec.Command("sh", "-c", `find / -name initexmf 2>&1 | grep -v "Permission denied" | grep -v "Invalid argument" | grep -v "No such file or directory"`)
		output, _ := InitexmfCmd.CombinedOutput()
		initexmf = strings.Trim(string(output), "\n")
	} else if runtime.GOOS == "windows" {
		InitexmfCmd := exec.Command("where", "initexmf")
		output, _ := InitexmfCmd.CombinedOutput()
		initexmf = strings.Trim(string(output), "\n")
	}
	if initexmf == "" {
		Log("Error", "initexmf komutu bulunamadı")
	}
	return initexmf
}

func (a *App) ChooseDistro(distro string) {
	conf := ReadConf()

	conf.LastDistro = distro
	conf.WriteConf()

}
