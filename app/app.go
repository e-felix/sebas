package app

import (
	App "fyne.io/fyne/v2/app"
	Container "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	APP_TITLE = "Sebas - Your personal butler"
)

func Run() {
	app := App.New()
	window := app.NewWindow(APP_TITLE)

	hello := widget.NewLabel("Greetings master.")
	whatCanIDoButton := widget.NewButton("Greetings Sebas.", func() {
		hello.SetText("How can I make your journey the most comfortable ?")
	})
	container := Container.NewVBox(hello, whatCanIDoButton)
	window.SetContent(container)

	window.ShowAndRun()
}
