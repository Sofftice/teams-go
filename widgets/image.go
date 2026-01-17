package widgets

import (
	"github.com/Sofftice/teams-go/publ"
	qt "github.com/mappu/miqt/qt6"
)

type SImageWidget struct {
	*qt.QLabel
}

func (i *SImageWidget) Constrain(w int, h int) *SImageWidget {
	i.SetPixmap(i.Pixmap2().Scaled3(w, h, qt.IgnoreAspectRatio, qt.SmoothTransformation))
	return i
}

func NewImage(path string) *SImageWidget {
	label := &SImageWidget{qt.NewQLabel3("")}
	label.SetPixmap(qt.NewQPixmap4(publ.GetResourcePath(path, true)))

	return label
}
