package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
)

var APP_TITLE = "Sebas - The Tech Butler"

func main() {
	body := core.NewBody(APP_TITLE)

	greetings := core.NewButton(body)
	greetings.SetText("Greetings, Sir.").OnClick(func(_ events.Event) {
		core.MessageSnackbar(body, "Greetings, Sebas.")
	})

	body.RunMainWindow()
}
