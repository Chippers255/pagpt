package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	app := tview.NewApplication()

	// Fetch system information
	hostname, _ := os.Hostname()
	osInfo, _ := host.Info()
	cpuInfos, _ := cpu.Info()
	memInfo, _ := mem.VirtualMemory()

	// Function to update CPU usage
	updateCPUUsage := func(cpuView *tview.TextView) {
		for {
			percentages, err := cpu.Percent(time.Second, false)
			if err == nil && len(percentages) > 0 {
				app.QueueUpdateDraw(func() {
					cpuText := fmt.Sprintf("[yellow]CPUs: [white]%.2f%% Usage", percentages[0])
					cpuView.SetText(cpuText)
				})
			}
			time.Sleep(time.Second)
		}
	}

	// Prepare system information text components
	hostnameText := fmt.Sprintf("[yellow]Hostname: [white]%s", hostname)
	osText := fmt.Sprintf("[yellow]OS: [white]%s-%s", osInfo.Platform, osInfo.PlatformVersion)
	cpuModelText := fmt.Sprintf("[yellow]CPU: [white]%dx %s", runtime.NumCPU(), cpuInfos[0].ModelName)
	memoryText := fmt.Sprintf("[yellow]Memory: [white]%v MB Free / %v MB Total",
		memInfo.Available/1024/1024,
		memInfo.Total/1024/1024,
	)

	// Create text views for each piece of system information
	hostnameView := tview.NewTextView().SetDynamicColors(true).SetText(hostnameText)
	osView := tview.NewTextView().SetDynamicColors(true).SetText(osText)
	cpuModelView := tview.NewTextView().SetDynamicColors(true).SetText(cpuModelText)
	cpuUsageView := tview.NewTextView().SetDynamicColors(true)
	memoryView := tview.NewTextView().SetDynamicColors(true).SetText(memoryText)

	// Start the CPU usage update goroutine
	go updateCPUUsage(cpuUsageView)

	// Create a horizontal flex layout for the system information panel
	systemInfoPanel := tview.NewFlex().
		AddItem(tview.NewBox().SetBorderPadding(1, 1, 1, 1), 1, 0, false). // Left margin
		AddItem(hostnameView, 0, 1, false).
		AddItem(osView, 0, 1, false).
		AddItem(cpuModelView, 0, 1, false).
		AddItem(cpuUsageView, 0, 1, false).
		AddItem(memoryView, 0, 1, false).
		AddItem(tview.NewBox().SetBorderPadding(1, 1, 1, 1), 1, 0, false) // Right margin

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
