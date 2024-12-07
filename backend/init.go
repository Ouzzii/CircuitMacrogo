package backend

type App struct {
	Configuration Conf
}

type Conf struct {
	Workspace     string            `json:"workspace"`
	PdflatexPaths map[string]string `json:"pdflatexPaths"`
	LastDistro    string            `json:"last-distro"`
}
