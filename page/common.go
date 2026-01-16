package page

import qt "github.com/mappu/miqt/qt6"

func prepare(window *qt.QMainWindow) *qt.QWidget {
	central := window.CentralWidget()
	if central != nil {
		central.DeleteLater()
	}

	widget := qt.NewQWidget2()
	window.SetCentralWidget(widget)
	return widget
}
