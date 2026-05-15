package publ

import (
	"embed"
	"encoding/json"
	qt "github.com/mappu/miqt/qt6"
	"log"
	"os"
)

var embeddedResources embed.FS

type resourceRegistry struct {
	Version     int      `json:"version"`
	Fonts       []string `json:"fonts"`
	StyleSheets []string `json:"stylesheets"`
}

var ResourceRegistry resourceRegistry
var CompiledStyleSheets = ""

// ReadResource reads a file from the embedded resources/ directory.
// Accepts only relative paths with forward slashes.
func ReadResource(path string) ([]byte, error) {
	return embeddedResources.ReadFile("resources/" + path)
}

// InitializeResources is the method called when the program starts
func InitializeResources(fs embed.FS) {
	log.Println("Initializing all resources")
	embeddedResources = fs

	initializeRegistry()
	compileStyleSheets()
	log.Println("Resources ready!")
}

func initializeRegistry() {
	data, err := ReadResource("registry.json")
	if err != nil {
		log.Fatalf("Failed to read registry.json: %s", err)
	}

	if err = json.Unmarshal(data, &ResourceRegistry); err != nil {
		log.Fatalf("Failed to unmarshal registry.json: %s", err)
	}
}

func compileStyleSheets() {
	for _, path := range ResourceRegistry.StyleSheets {
		data, err := ReadResource(path)
		if err != nil {
			log.Printf("Error: cannot read stylesheet %s: %s\n", path, err)
			continue
		}

		CompiledStyleSheets += "\n" + string(data)
	}
}

func LoadFonts() {
	for _, path := range ResourceRegistry.Fonts {
		data, err := ReadResource(path)
		if err != nil {
			log.Printf("Error: cannot read font %s: %s\n", path, err)
			continue
		}

		tmp, err := os.CreateTemp("", "sfteams-font-*.ttf")
		if err != nil {
			log.Printf("Error: cannot create temp file for font %s: %s\n", path, err)
			continue
		}

		if _, err = tmp.Write(data); err != nil {
			log.Printf("Error: cannot write temp font %s: %s\n", path, err)
			tmp.Close()
			continue
		}
		tmp.Close()

		qt.QFontDatabase_AddApplicationFont(tmp.Name())
	}
}
