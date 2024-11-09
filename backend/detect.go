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
	//last distro check

	if runtime.GOOS == "linux" {
		//miktex-pdflatex
		detectPdfCmd := exec.Command("sh", "-c", "find / -name miktex-pdflatex 2>&1 | grep -v 'Permission denied' | grep -v 'Invalid argument' | grep -v 'No such file or directory'")
		output, err := detectPdfCmd.CombinedOutput()
		if err != nil {
			//Log("Error", fmt.Sprintf("miktex - detect pdflatex komutunu başlatma hatası: %v", err))
			return a.configuration
		}
		a.configuration.AddPdflatexPath("miktex", strings.Replace(string(output), "\n", "", 1))
		a.configuration.WriteConf()

		//pdflatex
		detectPdfCmd = exec.Command("sh", "-c", "which pdflatex")

		output, err = detectPdfCmd.CombinedOutput()
		if err != nil {
			//Log("Error", fmt.Sprintf("texlive - detect pdflatex komutunu başlatma hatası: %v", err))
			return a.configuration
		}

		for _, pdflatexpath := range strings.Split(string(output), "\n") {
			if pdflatexpath != "" {
				pdflatexDistroCmd := exec.Command("sh", "-c", fmt.Sprintf("%v --version", pdflatexpath))
				output, err = pdflatexDistroCmd.CombinedOutput()
				if err != nil {
					//Log("Error", fmt.Sprintf("pdflatex versiyonu kontrol edilirken hata ile karşılaşıldı: %v\n", err))
					return a.configuration
				}
				if strings.Contains(string(output), "TeX Live") {
					a.configuration.AddPdflatexPath("texlive", pdflatexpath)
				}
			}
		}
		a.configuration.WriteConf()

	} else if runtime.GOOS == "windows" {

		fmt.Println("abcjdbjkfds", a.configuration.Workspace)
		detectPdfCmd := exec.Command("where", "pdflatex")
		output, err := detectPdfCmd.CombinedOutput()
		for _, pdflatexpath := range strings.Split(string(output), "\n") {
			pdflatexpath = strings.TrimSpace(pdflatexpath)
			if pdflatexpath == "" {
				break
			}
			fmt.Println(pdflatexpath, "--version")
			pdflatexCmd := exec.Command(pdflatexpath, "--version")
			output, err = pdflatexCmd.CombinedOutput()
			if err != nil {
				LogWithDetails("Error - " + fmt.Sprintf("%v - detect pdflatex komutunu başlatma hatası: %v", pdflatexpath, err))
			}
			if strings.Contains(string(output), "TeX Live") {
				a.configuration.AddPdflatexPath("texlive", pdflatexpath)
			} else if strings.Contains(string(output), "MiKTeX") {
				a.configuration.AddPdflatexPath("miktex", pdflatexpath)
			}
			a.configuration.WriteConf()

		}
		if err != nil {
			LogWithDetails("Error - " + fmt.Sprintf("miktex - detect pdflatex komutunu başlatma hatası: %v", err))
			return a.configuration
		}
	}
	fmt.Println("abcjdbjkfds", ReadConf().Workspace)
	fmt.Println("abcjdbjkfds", a.configuration.Workspace)
	return a.configuration

}

func (a *App) Boxdims_is_installed() Boxdims {
	UpdateEnv()
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
					//Log("Info", "Miktex: boxdims.sty başarılı bir şekilde indirildi")
					boxdims.Miktex = "exist"
				} else {
					//Log("Error", "Miktex: internet bağlantısı olmadığından boxdims.sty indirilemedi")
					boxdims.Miktex = "not exist"
				}
			} else {
				//Log("Info", "Miktex: boxdims.sty zaten kurulu")
				boxdims.Miktex = "exist"
			}
		} else {
			//Log("Warning", "miktex: cihazda miktex bulunmamaktadır")
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
							//Log("Info", "Texlive için boxdims başarıyla indirildi")
							boxdims.Texlive = "exist"
						} else {
							//Log("Error", fmt.Sprintf("Texlive için boxdims indirilirken hata ile karşılaşıldı: %v\n", result))
						}
					} else {
						//Log("Error", "texlive: internet bağlantısı olmadığından boxdims.sty indirilemedi")
					}
				} else {
					//Log("Info", "texlive: boxdims.sty zaten kurulu")
					boxdims.Texlive = "exist"
				}
			} else {
				//Log("Error", "texlive: Dizin bulunamadı")
			}
		} else {
			//Log("Warning", "texlive: cihazda texlive bulunmamaktadır")
		}

	} else if runtime.GOOS == "windows" {
		if _, exists := conf.PdflatexPaths["miktex"]; exists {
			miktexpath := filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "MiKTeX", "tex", "latex", "circuit_macros")
			if !Exists(miktexpath) {
				if CheckInternet() {
					Download(miktexpath+".zip", "https://mirrors.ctan.org/graphics/circuit_macros.zip")
					Unzip(miktexpath+".zip", filepath.Dir(miktexpath))
					exec.Command(get_initexmf(), "--update-fndb")
					//Log("Info", "Miktex: boxdims.sty başarılı bir şekilde indirildi")
					boxdims.Miktex = "exist"
				} else {
					//Log("Error", "Miktex: internet bağlantısı olmadığından boxdims.sty indirilemedi")
					boxdims.Miktex = "not exist"
				}
			} else {
				//Log("Info", "Miktex: boxdims.sty zaten kurulu")
				boxdims.Miktex = "exist"
			}
		} else {
			//Log("Warning", "miktex: cihazda miktex bulunmamaktadır")
		}
		if _, exists := conf.PdflatexPaths["texlive"]; exists {
			texlivepathCmd := exec.Command("cmd", "/C", "dir", "C:\\texlive\\texmf-dist", "/s")
			output, _ := texlivepathCmd.CombinedOutput()
			texlivepath := filepath.Join(strings.TrimSpace(strings.Split(strings.Split(string(output), " Directory of ")[1], "\n")[0]), "texmf-dist", "tex", "latex")
			if Exists(texlivepath) {
				if !Exists(filepath.Join(texlivepath, "circuit-macros")) {
					if CheckInternet() {
						curlCmd := exec.Command("curl", "-L", "-o", filepath.Join(texlivepath, "circuit_macros.zip"), "https://mirrors.ctan.org/graphics/circuit_macros.zip")
						if err := curlCmd.Run(); err != nil {
							fmt.Println("Error downloading zip:", err)
						}

						// Adım 2: zip dosyasını unzip ile çıkartma
						unzipCmd := exec.Command("unzip", filepath.Join(texlivepath, "circuit_macros.zip"), "-d", texlivepath)
						if err := unzipCmd.Run(); err != nil {
							fmt.Println("Error unzipping file:", err)
						}

						// Adım 3: zip dosyasını silme
						delCmd := exec.Command("del", filepath.Join(texlivepath, "circuit_macros.zip"))
						if err := delCmd.Run(); err != nil {
							fmt.Println("Error deleting zip:", err)
						}

						// Adım 4: dosya taşımak (move)
						moveCmd := exec.Command("move", filepath.Join(texlivepath, "circuit_macros"), filepath.Join(texlivepath, "circuit-macros"))
						if err := moveCmd.Run(); err != nil {
							fmt.Println("Error moving directory:", err)
						}
						boxdims.Texlive = "exist"
					}
				} else {
					boxdims.Texlive = "exist"
				}
			}
		}
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
		//Log("Error", "initexmf komutu bulunamadı")
	}
	return initexmf
}

func (a *App) ChooseDistro(distro string) {
	conf := ReadConf()

	conf.LastDistro = distro
	conf.WriteConf()

}
