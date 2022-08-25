package main

import (
	"Flex/flex"
	"fmt"

	"github.com/gookit/color"
)

func main() {
	// MainContainer is output with each change
	flex.MainContainer.SetSizes(29, 120)

	// some flex settings
	flexSettings := []flex.Flex{
		{ // row Top Left
			Direction:      flex.Direction.Row,
			Wrap:           flex.Wrap.Wrap,
			JustifyContent: flex.JustifyContent.Start,
			AlignItems:     flex.AlignItems.Start,
			AlignContent:   flex.AlignContent.Start,
		},
		{ // column Center
			Direction:      flex.Direction.Column,
			Wrap:           flex.Wrap.Wrap,
			JustifyContent: flex.JustifyContent.Center,
			AlignItems:     flex.AlignItems.Center,
			AlignContent:   flex.AlignContent.Start,
		},
		{ // row Bottom space between
			Direction:      flex.Direction.Row,
			Wrap:           flex.Wrap.Wrap,
			JustifyContent: flex.JustifyContent.SpaceBetween,
			AlignItems:     flex.AlignItems.End,
			AlignContent:   flex.AlignContent.Start,
		},
		{ // column Center Space evenly
			Direction:      flex.Direction.Column,
			Wrap:           flex.Wrap.Wrap,
			JustifyContent: flex.JustifyContent.SpaceEvenly,
			AlignItems:     flex.AlignItems.Center,
			AlignContent:   flex.AlignContent.Start,
		},
	}

	// create a 5x5 blue square
	var SquareBlue = flex.Widget{
		X:            0,
		Y:            0,
		Height:       5,
		Width:        10, // for visibility divide by 2, due to the proportions of the cell 1 to 2
		BgColor:      color.BgBlue,
		FgColor:      color.FgBlue,
		Children:     nil,
		FlexSettings: flex.Flex{},
	}

	// create a 4x8 white square
	var SquareWhite = flex.Widget{
		X:            10,
		Y:            0,
		Height:       8,
		Width:        8,
		BgColor:      color.BgWhite,
		FgColor:      color.FgBlue,
		Children:     nil,
		FlexSettings: flex.Flex{},
	}

	// create a 6x3 red square
	var SquareRed = flex.Widget{
		X:            10,
		Y:            0,
		Height:       3,
		Width:        12,
		BgColor:      color.BgRed,
		FgColor:      color.FgBlue,
		Children:     nil,
		FlexSettings: flex.Flex{},
	}

	// create 40x20 gray container
	var Row = flex.Widget{
		X:       0,
		Y:       0,
		Height:  20,
		Width:   80,
		BgColor: color.BgDarkGray,
		FgColor: color.FgBlue,

		// the container contains three children - those squares described above
		Children: []flex.Widget{SquareWhite, SquareBlue, SquareRed},

		// flex settings of the container are configured here
		FlexSettings: flexSettings[0],
	}

	for _, next := range flexSettings {
		Row.FlexSettings = next

		// MainContainer occupies the entire width of the console and is output with each change
		flex.MainContainer.Place(&Row)

		// start of rendering
		flex.MainContainer.Render()

		// Misha fifth grade
		var m int
		fmt.Scan(&m)
	}
}
