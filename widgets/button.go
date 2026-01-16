package widgets

import qt "github.com/mappu/miqt/qt6"

type SPrimaryButton struct {
	*qt.QPushButton
}

func NewPrimaryButton(text string) *SPrimaryButton {
	btn := &SPrimaryButton{
		qt.NewQPushButton3(text),
	}

	btn.SetCursor(qt.NewQCursor2(qt.PointingHandCursor))
	btn.SetProperty("class", qt.NewQVariant11("primary"))

	return btn
}
