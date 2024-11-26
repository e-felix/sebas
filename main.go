package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/math32"
	"cogentcore.org/core/styles"
)

var APP_TITLE = "Sebas - Your Tech Butler"
var mainContainer *core.Frame

func main() {
	body := core.NewBody(APP_TITLE)
	mainContainer = core.NewFrame(body)
	mainContainer.Styler((func(s *styles.Style) {
		s.Display = styles.Stacked
		s.Grow = math32.Vector2{X: 1, Y: 1}
	}))

	createTopPanel()

	body.RunMainWindow()
}

func createTopPanel() {
	menu := core.NewFrame(mainContainer)
	menu.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Justify.Content = styles.SpaceBetween
		s.Align.Items = styles.Center
		s.Grow = math32.Vector2{X: 1, Y: 0}
	})

	left := core.NewFrame(menu)
	core.NewText(left).SetText("Sebas - Your Tech Butler")
	left.Styler(func(s *styles.Style) {
		s.Grow = math32.Vector2{X: 1, Y: 0}
	})

	right := core.NewFrame(menu)
	core.NewChooser(right).SetStrings("GambettesBox", "MyLittleBox")
	core.NewChooser(right).SetStrings("Production", "Staging")
	right.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Justify.Content = styles.End
		s.Grow = math32.Vector2{X: 1, Y: 0}
	})
}
