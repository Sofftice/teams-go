package page

import qt "github.com/mappu/miqt/qt6"

// prepare prepares a window to change its contents by removing the old content and providing a central widget
func prepare(window *qt.QMainWindow) *qt.QWidget {
	central := window.CentralWidget()
	if central != nil {
		central.DeleteLater()
	}

	widget := qt.NewQWidget2()
	window.SetCentralWidget(widget)
	return widget
}

// NewCenterVLayout allows for quick and streamlined creation of VBox layouts which have its contents centered
func NewCenterVLayout(widget *qt.QWidget, content func(layout *qt.QVBoxLayout)) *qt.QVBoxLayout {
	layout := qt.NewQVBoxLayout2()

	layout.AddStretch()
	content(layout)
	layout.AddStretch()

	widget.SetLayout(layout.QLayout)
	return layout
}

// NewCenterVLayout2 NewCenterVLayout with a layout instead of widget
func NewCenterVLayout2(p *qt.QBoxLayout, content func(layout *qt.QVBoxLayout)) *qt.QVBoxLayout {
	layout := qt.NewQVBoxLayout2()

	layout.AddStretch()
	content(layout)
	layout.AddStretch()

	p.AddLayout(layout.QLayout)
	return layout
}

// NewCenterHLayout allows for quick and streamlined creation of HBox layouts which have its contents centered
func NewCenterHLayout(widget *qt.QWidget, content func(layout *qt.QHBoxLayout)) *qt.QHBoxLayout {
	layout := qt.NewQHBoxLayout2()

	layout.AddStretch()
	content(layout)
	layout.AddStretch()

	widget.SetLayout(layout.QLayout)
	return layout
}

// NewCenterHLayout2 NewCenterHLayout with a layout instead of widget
func NewCenterHLayout2(p *qt.QBoxLayout, content func(layout *qt.QHBoxLayout)) *qt.QHBoxLayout {
	layout := qt.NewQHBoxLayout2()

	layout.AddStretch()
	content(layout)
	layout.AddStretch()

	p.AddLayout(layout.QLayout)
	return layout
}
