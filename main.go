package main

import (
	"embed"
	"github.com/Sofftice/teams-go/page"
	"github.com/Sofftice/teams-go/publ"
	qt "github.com/mappu/miqt/qt6"
	"log"
	"os"
)

//go:embed resources
var res embed.FS

func main() {
	log.Println("Preparing Sofftice Teams")
	publ.InitializeResources(res)

	qt.NewQApplication(os.Args)
	publ.LoadFonts()

	window := qt.NewQMainWindow2()
	window.SetWindowTitle("Sofftice Teams")
	window.SetMinimumSize2(816, 520)
	window.SetStyleSheet(publ.CompiledStyleSheets)

	page.DisplayLoginPage(window)

	window.Show()

	qt.QCoreApplication_SetApplicationName("SFTeams")
	qt.QGuiApplication_SetApplicationDisplayName("Sofftice Teams")

	log.Println("Executing QApplication")
	qt.QApplication_Exec()
}
