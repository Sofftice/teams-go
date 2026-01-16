package page

import qt "github.com/mappu/miqt/qt6"

func DisplayLoginPage(window *qt.QMainWindow) {
	widget := prepare(window)

	NewCenterHLayout(widget, func(layout *qt.QHBoxLayout) {
		NewCenterVLayout2(layout.QBoxLayout, func(layout *qt.QVBoxLayout) {
			text := qt.NewQLabel3("Test")
			layout.AddWidget(text.QWidget)

			text2 := qt.NewQLabel3("Test")
			layout.AddWidget(text2.QWidget)
		})
	})
}
