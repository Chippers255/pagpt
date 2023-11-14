package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Create the layout panels.
	systemInfoPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetText("System Information").
		SetTextAlign(tview.AlignCenter)
	terminalPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetText("Terminal Window").
		SetTextAlign(tview.AlignCenter)
	notificationsPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetText("Notifications").
		SetTextAlign(tview.AlignCenter)

	// Create a flex layout that arranges panels vertically.
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	// Add the system information panel with fixed height.
	flex.AddItem(systemInfoPanel, 3, 1, false)

	// Create a horizontal flex for the terminal and notifications panels.
	horizontalFlex := tview.NewFlex().
		AddItem(terminalPanel, 0, 2, false).
		AddItem(notificationsPanel, 0, 1, false)

	// Add the horizontal flex to the main flex layout.
	flex.AddItem(horizontalFlex, 0, 3, false)

	// Set the root and start the application.
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
