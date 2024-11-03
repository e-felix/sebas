package app

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SebasTheme struct{}

var (
	APP_TITLE string     = "Sebas - Your personal butler"
	_         fyne.Theme = (*SebasTheme)(nil)
)

func Run() {
	sebasApp := app.New()
	sebasApp.Settings().SetTheme(&SebasTheme{})
	window := sebasApp.NewWindow(APP_TITLE)

	// header
	var projectToLoad string
	var envToLoad string
	loadedProject := binding.NewString()
	loadedEnv := binding.NewString()

	// left side
	headerTitle := canvas.NewText("Sebas - Your personal butler", color.White)
	titlesContainer := container.NewGridWithRows(1, headerTitle)

	projectChoices := widget.NewSelect([]string{"Project 1", "Project 2"}, func(value string) {
		log.Println("Project set to", value)
		projectToLoad = value
	})
	projectChoices.PlaceHolder = "Select a project"
	envChoices := widget.NewSelect([]string{"production", "staging"}, func(value string) {
		log.Println("Environment set to", value)
		envToLoad = value
	})
	envChoices.PlaceHolder = "Select a environment"
	choicesContainer := container.NewVBox(projectChoices, envChoices)

	// right side
	loadAction := widget.NewButton("Load", func() {
		if projectToLoad == "" || envToLoad == "" {
			log.Println("Please select a project and an environment to load.")
			return
		}
		log.Println(fmt.Sprintf("Loading project %v in %v context", projectToLoad, envToLoad))
		loadedProject.Set(projectToLoad)
		loadedEnv.Set(envToLoad)
	})
	choicesAndActionsContainer := container.NewGridWithColumns(2, choicesContainer, loadAction)

	header := container.NewAdaptiveGrid(3, titlesContainer, layout.NewSpacer(), choicesAndActionsContainer)

	// body
	cmdTitle := widget.NewLabel("Commands")
	var cmds = [][]string{{"cmd", "args"}, {"ls", "-la"}, {"echo", "Hello ", "World"}}
	cmdsTable := widget.NewTable(
		func() (int, int) {
			return len(cmds), len(cmds[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Commands")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(cmds[tci.Row][tci.Col])
		},
	)
	cmdsTable.StickyRowCount = len(cmds)

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 1
	line2 := canvas.NewLine(color.White)
	line2.StrokeWidth = 1

	body := container.NewVBox(line, cmdTitle, cmdsTable, line2)

	// footer
	version := canvas.NewText("Version 0.1.0", color.White)
	version.TextStyle.Italic = true
	envLabel := widget.NewLabelWithData(loadedEnv)
	separator := canvas.NewText("|", color.White)
	projectLabel := widget.NewLabelWithData(loadedProject)
	loadedContainer := container.NewHBox(envLabel, separator, projectLabel)
	footer := container.NewHBox(version, layout.NewSpacer(), loadedContainer)

	// layout
	container := container.NewVBox(header, body, footer)
	window.SetContent(container)

	window.ShowAndRun()
}

func (m SebasTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m SebasTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m SebasTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m SebasTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
