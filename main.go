package main

import (
	"fmt"
	qt "github.com/mappu/miqt/qt6"
	"log"
	"os"
)

var counter = 0

func main() {
	log.Println("Preparing Sofftice Teams")
	InitializeResources()

	qt.NewQApplication(os.Args)

	window := qt.NewQMainWindow2()
	window.SetWindowTitle("Sofftice Teams")
	window.SetMinimumSize2(816, 520)
	window.SetStyleSheet(CompiledStyleSheets)

	btn := qt.NewQPushButton(window.Window())
	btn.SetText("Hoi")
	btn.OnPressed(func() {
		counter++
		btn.SetText(fmt.Sprintf("Hi: %d", counter))
	})

	window.Show()

	log.Println("Executing QApplication")
	qt.QApplication_Exec()
}
