package flex

import (
	"fmt"

	"github.com/gookit/color"
)

type (
	direction struct {
		Row    string
		Column string
	}

	wrap struct {
		NoWrap string
		Wrap   string
	}

	justifyContent struct {
		Start        string
		End          string
		Center       string
		SpaceBetween string
		SpaceAround  string
		SpaceEvenly  string
	}

	alignItems struct {
		Start   string
		End     string
		Center  string
		Stretch string
	}

	alignContent struct {
		Start        string
		End          string
		Center       string
		Stretch      string
		SpaceBetween string
		SpaceAround  string
		SpaceEvenly  string
	}
)

var (
	Direction      = newDirection()
	Wrap           = newWrap()
	JustifyContent = newJustifyContent()
	AlignItems     = newAlignItems()
	AlignContent   = newAlignContent()
)

func newDirection() *direction {
	return &direction{
		Row:    "Row",
		Column: "Column",
	}
}

func newWrap() *wrap {
	return &wrap{
		NoWrap: "NoWrap",
		Wrap:   "Wrap",
	}
}

func newJustifyContent() *justifyContent {
	return &justifyContent{
		Start:        "Start",
		End:          "End",
		Center:       "Center",
		SpaceBetween: "SpaceBetween",
		SpaceAround:  "SpaceAround",
		SpaceEvenly:  "SpaceEvenly",
	}
}

func newAlignItems() *alignItems {
	return &alignItems{
		Start:   "Start",
		End:     "End",
		Center:  "Center",
		Stretch: "Stretch",
	}
}

func newAlignContent() *alignContent {
	return &alignContent{
		Start:        "Start",
		End:          "End",
		Center:       "Center",
		Stretch:      "Stretch",
		SpaceBetween: "SpaceBetween",
		SpaceAround:  "SpaceAround",
		SpaceEvenly:  "SpaceEvenly",
	}
}

type Flex struct {
	Direction      string
	Wrap           string
	JustifyContent string
	AlignItems     string
	AlignContent   string
}

type pixel struct {
	char    rune
	fgColor color.Color
	bgColor color.Color
}

type Matrix struct {
	Height int
	Width  int
	Data   [][]pixel
}

var MainContainer Matrix

func (m *Matrix) SetSizes(height int, width int) {
	m.Height = height
	m.Width = width
	m.Data = make([][]pixel, height)

	for i := 0; i < height; i++ {
		m.Data[i] = make([]pixel, width)
		for j := 0; j < width; j++ {
			m.Data[i][j] = pixel{' ', color.FgGray, color.BgBlack}
		}
	}
}

func (m *Matrix) Render() {
	fmt.Print("\033[H\033[2J")

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			color.New(m.Data[i][j].fgColor, m.Data[i][j].bgColor).Printf("%c", m.Data[i][j].char)
		}
		fmt.Println()
	}
}

func (m *Matrix) Place(w *Widget) {
	w.AdjustChildrenPlacement()

	for i := w.Y; i < w.Y+w.Height && i < m.Height; i++ {
		for j := w.X; j < w.X+w.Width && j < m.Width; j++ {
			m.Data[i][j] = pixel{' ', w.FgColor, w.BgColor}
		}

	}
	if w.Children != nil {
		for _, child := range w.Children {
			m.Place(&child)
		}
	}

}

type Widget struct {
	X            int
	Y            int
	Height       int
	Width        int
	BgColor      color.Color
	FgColor      color.Color
	Children     []Widget
	FlexSettings Flex
}

func (w *Widget) AdjustChildrenPlacement() {
	if w.Children != nil {

		var (
			contentWidth          int // sum of children width
			contentHeight         int // sum of children height
			coordAlong            int
			coordAcross           int
			containerLengthAlong  int
			containerLengthAcross int
			contentLengthAlong    int
			//contentLengthAcross     int
			paddingAlongStart       int
			paddingAlongEnd         int
			paddingAcrossStart      int
			paddingAcrossEnd        int
			spaceTotal              int
			spaceBorder             int
			spaceObject             int
			positionAlongStart      int
			alignmentPointAcross    int
			objectLengthNumerator   int
			objectLengthDenominator int = 1
			objectPositionAlong     int
			objectPositionAcross    int
			objectLengthAlong       int
			objectLengthAcross      int
		)

		for _, child := range w.Children {
			contentWidth += child.Width
			contentHeight += child.Height
		}

		switch w.FlexSettings.Direction {

		case Direction.Row:
			coordAlong = w.X
			coordAcross = w.Y
			containerLengthAlong = w.Width
			containerLengthAcross = w.Height
			contentLengthAlong = contentWidth
			//contentLengthAcross = contentHeight
			paddingAlongStart = 0
			paddingAlongEnd = 0
			paddingAcrossStart = 0
			paddingAcrossEnd = 0

		case Direction.Column:
			coordAlong = w.Y
			coordAcross = w.X
			containerLengthAlong = w.Height
			containerLengthAcross = w.Width
			contentLengthAlong = contentHeight
			//contentLengthAcross = contentWidth
			paddingAlongStart = 0
			paddingAlongEnd = 0
			paddingAcrossStart = 0
			paddingAcrossEnd = 0
		}

		spaceTotal = containerLengthAlong - paddingAlongStart - paddingAlongEnd - contentLengthAlong

		switch w.FlexSettings.JustifyContent {

		case JustifyContent.Start:
			positionAlongStart = 0

		case JustifyContent.End:
			positionAlongStart = containerLengthAlong - contentLengthAlong

		case JustifyContent.Center:
			spaceBorder = spaceTotal / 2
			positionAlongStart = coordAlong + paddingAcrossStart + spaceBorder

		case JustifyContent.SpaceAround:
			spaceBorder = spaceTotal / (len(w.Children) * 2)
			positionAlongStart = coordAlong + paddingAlongStart + spaceBorder
			spaceObject = spaceBorder * 2

		case JustifyContent.SpaceBetween:
			spaceBorder = spaceTotal / (len(w.Children) - 1)
			positionAlongStart = coordAlong + paddingAlongStart
			spaceObject = spaceBorder

		case JustifyContent.SpaceEvenly:
			spaceBorder = spaceTotal / (len(w.Children) + 1)
			positionAlongStart = coordAlong + paddingAlongStart + spaceBorder
			spaceObject = spaceBorder
		}

		switch w.FlexSettings.AlignItems {

		case AlignItems.Start:
			alignmentPointAcross = coordAcross + paddingAcrossStart
			objectLengthNumerator = 0

		case AlignItems.End:
			alignmentPointAcross = coordAcross + containerLengthAcross - paddingAcrossEnd
			objectLengthNumerator = -1

		case AlignItems.Center:
			alignmentPointAcross = coordAcross + (containerLengthAcross-paddingAcrossStart-paddingAcrossEnd)/2
			objectLengthNumerator = -1
			objectLengthDenominator = 2

			//case AlignItems.Stretch:

		}

		objectPositionAlong = positionAlongStart

		for i := 0; i < len(w.Children); i++ {
			var child *Widget = &w.Children[i]

			switch w.FlexSettings.Direction {

			case Direction.Row:
				objectLengthAlong = child.Width
				objectLengthAcross = child.Height

			case Direction.Column:
				objectLengthAlong = child.Height
				objectLengthAcross = child.Width
			}

			objectPositionAcross = alignmentPointAcross + objectLengthAcross*objectLengthNumerator/objectLengthDenominator

			switch w.FlexSettings.Direction {

			case Direction.Row:
				child.X = objectPositionAlong
				child.Y = objectPositionAcross

			case Direction.Column:
				child.X = objectPositionAcross
				child.Y = objectPositionAlong
			}

			objectPositionAlong += objectLengthAlong + spaceObject

			for i := 0; i < len(w.Children); i++ {
				var child *Widget = &w.Children[i]
				child.AdjustChildrenPlacement()
			}
		}
	}
}
