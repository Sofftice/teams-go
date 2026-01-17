package publ

import (
	"encoding/json"
	qt "github.com/mappu/miqt/qt6"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type resourceRegistry struct {
	Version     int      `json:"version"`
	Fonts       []string `json:"fonts"`
	StyleSheets []string `json:"stylesheets"`
}

const ProcessRegistryVersion = 0

var ResourceRegistry resourceRegistry
var CompiledStyleSheets = ""

func GetResourcePath(path string, absolute bool) string {
	path = strings.ReplaceAll(path, "\\", string(filepath.Separator))
	path = strings.ReplaceAll(path, "/", string(filepath.Separator))

	path = filepath.Join("resources", path)

	if absolute {
		wd, _ := os.Getwd()
		path = filepath.Join(wd, path)
	}

	return path
}

// ReadResource reads a resource from the resources/ directory. It accepts only relative paths from within the resources directory with forward slashes.
func ReadResource(path string) ([]byte, error) {
	return os.ReadFile(GetResourcePath(path, false))
}

// InitializeResources is the method called when the program starts
func InitializeResources() {
	log.Println("Initializing all resources")

	initializeRegistry()
	compileStyleSheets()

	log.Println("Resources ready!")
}

func initializeRegistry() {
	data, err := ReadResource("registry.json")
	if err != nil {
		log.Fatalf("Failed to get registry.json: %s", err)
	}

	err = json.Unmarshal(data, &ResourceRegistry)
	if err != nil {
		log.Fatalf("Failed to unmarshal registry.json: %s", err)
	}
}

// compileStyleSheets sets the CompiledStyleSheets variable
func compileStyleSheets() {
	for _, path := range ResourceRegistry.StyleSheets {
		data, err := ReadResource(path)
		if err != nil {
			log.Printf("Error: cannot read data resource at path %s: %s\n", data, err)
		}

		CompiledStyleSheets += "\n" + string(data)
	}
}

func LoadFonts() {
	for _, path := range ResourceRegistry.Fonts {
		qt.QFontDatabase_AddApplicationFont(GetResourcePath(path, true))
	}
}
