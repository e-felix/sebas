package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/math32"
	"cogentcore.org/core/styles"
)

var APP_TITLE = "Sebas - Your Tech Butler"
var mainContainer *core.Frame

func main() {
	body := core.NewBody(APP_TITLE)
	mainContainer = core.NewFrame(body)
	mainContainer.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Direction = styles.Column
		s.Justify.Content = styles.Start
		s.Align.Items = styles.Start
		s.Grow = math32.Vector2{X: 1, Y: 1}
	})

	createHeader()
	createBody()

	body.RunMainWindow()
}

func createHeader() {
	menu := core.NewFrame(mainContainer)
	menu.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Justify.Content = styles.SpaceBetween
		s.Align.Items = styles.Center
		s.Grow = math32.Vector2{X: 1, Y: 0}
		s.Border.Width.Bottom.Value = 1
		s.Padding.Bottom.Value = 12
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

func createBody() {
	body := core.NewFrame(mainContainer)
	body.Styler(func(s *styles.Style) {
		s.Display = styles.Grid
		s.Columns = 12
		s.Grow = math32.Vector2{X: 1, Y: 1}
	})

	sideBar := core.NewFrame(body)
	sideBar.Styler(func(s *styles.Style) {
		s.Columns = 1
		s.Border.Width.Right.Value = 1
		s.Grow = math32.Vector2{X: 1, Y: 1}
		s.Display = styles.Flex
		s.Direction = styles.Column
		s.Justify.Content = styles.Start
	})

	menu := core.NewFrame(sideBar)
	menu.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Direction = styles.Column
		s.Grow = math32.Vector2{X: 1, Y: 1}
	})

	commandsButton := core.NewButton(menu).SetText("Commands")
	commandsButton.Styler(func(s *styles.Style) {
		s.Grow = math32.Vector2{X: 1}
		radius := float32(5)
		s.Border.Radius.Top.Value = radius
		s.Border.Radius.Right.Value = radius
		s.Border.Radius.Bottom.Value = radius
		s.Border.Radius.Left.Value = radius
	})
	envsButton := core.NewButton(menu).SetText("Envs")
	envsButton.Styler(func(s *styles.Style) {
		s.Grow = math32.Vector2{X: 1}
		radius := float32(5)
		s.Border.Radius.Top.Value = radius
		s.Border.Radius.Right.Value = radius
		s.Border.Radius.Bottom.Value = radius
		s.Border.Radius.Left.Value = radius
	})
	sebasVersion := core.NewText(sideBar)
	sebasVersion.SetText("Sebas v0.1.0")
	sebasVersion.Styler(func(s *styles.Style) {
		s.Font.Style = styles.Italic
		s.Align.Self = styles.Center
		s.Grow = math32.Vector2{X: 1}
	})

	mainContent := core.NewFrame(body)
	mainContent.Styler(func(s *styles.Style) {
		s.Columns = 2
		s.Grow = math32.Vector2{X: 10, Y: 1}
	})

	/*
		commandsContent := core.NewFrame(mainContent)
		commandsContent.Styler(func(s *styles.Style) {
			s.Display = styles.Flex
			s.Direction = styles.Column
			s.Grow = math32.Vector2{X: 1, Y: 1}
		})

		core.NewText(commandsContent).SetText("COMMANDS")

		envsContent := core.NewFrame(mainContent)
		envsContent.Styler(func(s *styles.Style) {
			s.Display = styles.Flex
			s.Direction = styles.Column
			s.Grow = math32.Vector2{X: 1, Y: 1}
		})

		core.NewText(envsContent).SetText("ENVS")
	*/

	// Menu events
	commandsButton.OnClick(func(e events.Event) {
		for i := range mainContent.Children {
			mainContent.Child(i).Destroy()
		}
		commandsContent := core.NewFrame(mainContent)
		commandsContent.Styler(func(s *styles.Style) {
			s.Display = styles.Flex
			s.Direction = styles.Column
			s.Grow = math32.Vector2{X: 1, Y: 1}
		})
		core.NewText(commandsContent).SetText("COMMANDS")
		mainContent.Update()
	})
	envsButton.OnClick(func(e events.Event) {
		for i := range mainContent.Children {
			mainContent.Child(i).Destroy()
		}
		envsContent := core.NewFrame(mainContent)
		envsContent.Styler(func(s *styles.Style) {
			s.Display = styles.Flex
			s.Direction = styles.Column
			s.Grow = math32.Vector2{X: 1, Y: 1}
		})

		core.NewText(envsContent).SetText("ENVS")
		mainContent.Update()
	})
}
