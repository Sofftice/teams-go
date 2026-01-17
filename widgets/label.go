package widgets

import (
	qt "github.com/mappu/miqt/qt6"
)

type SLabel struct {
	*qt.QLabel
}

type SHeader struct {
	*qt.QLabel
}

func NewLabel(text string) *SLabel {
	label := &SLabel{qt.NewQLabel3(text)}
	label.SetProperty("class", qt.NewQVariant11("label"))
	return label
}

func NewHeader(text string) *SHeader {
	header := &SHeader{qt.NewQLabel3(text)}
	header.SetProperty("class", qt.NewQVariant11("header"))
	return header
}
