package page

import (
	"encoding/json"
	"fmt"
	"github.com/Sofftice/teams-go/publ"
	"github.com/Sofftice/teams-go/widgets"
	qt "github.com/mappu/miqt/qt6"
	"log"
)

func DisplayLoginFlowPage(window *qt.QMainWindow) {
	log.Println("Preparing login flow page...")
	widget := prepare(window)

	var statusWidget *widgets.SLabel

	NewCenterHLayout(widget, func(layout *qt.QHBoxLayout) {
		NewCenterVLayout2(layout.QBoxLayout, func(layout *qt.QVBoxLayout) {
			statusWidget = widgets.NewLabel("Please continue in the popup")
			layout.AddWidget(statusWidget.QWidget)
		})

		go startFlow(statusWidget)
	})

	log.Println("Login flow page is now displaying!")
}

func startFlow(label *widgets.SLabel) {
	resp, err := publ.DemandTokenResponse()
	if err != nil {
		log.Println("ERROR", err)
		label.SetText(fmt.Sprintf("Something's not right... %s; try restarting", err))
		return
	}

	dat, _ := json.MarshalIndent(resp, "", "    ")
	label.SetText(string(dat))
}
