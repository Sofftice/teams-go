package page

import (
	"github.com/Sofftice/teams-go/widgets"
	qt "github.com/mappu/miqt/qt6"
)

func DisplayLoginPage(window *qt.QMainWindow) {
	widget := prepare(window)

	NewCenterHLayout(widget, func(layout *qt.QHBoxLayout) {
		NewCenterVLayout2(layout.QBoxLayout, func(layout *qt.QVBoxLayout) {
			title := widgets.NewHeader("Welcome to Teams")
			layout.AddWidget(title.QWidget)

			subtitle := widgets.NewLabel("Chat and collaborate in one app.")
			subtitle.SetAlignment(qt.AlignCenter)
			layout.AddWidget(subtitle.QWidget)

			//image := widgets.NewImage("img/person.png").Constrain(200, 200)
			//layout.AddWidget(image.QWidget)

			layout.AddSpacing(50 + 200)

			signIn := widgets.NewPrimaryButton("Sign in")
			layout.AddWidget(signIn.QWidget)
		})
	})
}
