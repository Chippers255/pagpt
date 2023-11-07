package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Create the layout panels
	systemInfoPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetText("System Information and Resource Usage will be displayed here.")

	messagePanel := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetText("Messages will be displayed here. This is a placeholder for now.")

	aiOutputPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetText("AI Output will be displayed here.")

	userInputPanel := tview.NewInputField().
		SetLabel("User Input: ").
		SetFieldWidth(100).
		SetFieldBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	instructionsBar := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("Hotkeys: [Ctrl-C] Quit [Ctrl-R] Refresh")

	// Create the ASCII logo for the bootup screen
	logo := `
	______   __  __   ____     ______   ____     ____     ______   ______   __ __
	/ ____/   \ \/ /  / __ )   / ____/  / __ \   / __ \   / ____/  / ____/  / //_/
   / /         \  /  / __  |  / __/    / /_/ /  / / / /  / __/    / /      / ,<   
  / /___       / /  / /_/ /  / /___   / _, _/  / /_/ /  / /___   / /___   / /| |  
  \____/      /_/  /_____/  /_____/  /_/ |_|  /_____/  /_____/   \____/  /_/ |_|  
																				  
`

	// Create a flex layout to arrange panels
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText(logo).SetTextAlign(tview.AlignCenter), 0, 1, false).
		AddItem(systemInfoPanel, 0, 3, false).
		AddItem(messagePanel, 0, 2, false).
		AddItem(aiOutputPanel, 0, 2, false).
		AddItem(userInputPanel, 0, 1, true). // Set to true to capture input
		AddItem(instructionsBar, 1, 1, false)

	// Handler to update AI output panel when Enter is pressed
	userInputPanel.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			// Get the text from the input field
			message := userInputPanel.GetText()
			// Clear the input field for new input
			userInputPanel.SetText("")

			// Safely update the UI in the main thread
			app.QueueUpdateDraw(func() {
				// Append the message to the AI output panel
				currentText := aiOutputPanel.GetText(true)
				newText := currentText + message + "\n"
				aiOutputPanel.SetText(newText)
			})
		}
	})

	// Set the root and run the application
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
