package widgets

import (
	"github.com/Sofftice/teams-go/publ"
	qt "github.com/mappu/miqt/qt6"
	"log"
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
	label.SetPixmap(NewPixmap(path))

	return label
}

func NewPixmap(path string) *qt.QPixmap {
	data, err := publ.ReadResource(path)
	if err != nil {
		log.Printf("Error: cannot read image %s: %s\n", path, err)
		return qt.NewQPixmap()
	}

	pixmap := qt.NewQPixmap()
	pixmap.LoadFromData(&data[0], uint(len(data)))
	return pixmap
}
