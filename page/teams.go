package page

import (
	"github.com/Sofftice/teams-go/widgets"
	qt "github.com/mappu/miqt/qt6"
	"log"
)

func DisplayTeamsPage(window *qt.QMainWindow) {
	log.Println("Preparing teams page...")
	widget := prepare(window)

	NewCenterVLayout(widget, func(layout *qt.QVBoxLayout) {
		layout.AddWidget(widgets.NewHeader("My Teams").QWidget)
		layout.AddStretch()
	})

	log.Println("Teams page is now displaying")
}
