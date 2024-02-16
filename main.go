package main

import (
	"fmt"

	"github.com/altsko/speedrun-timer/utils"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

func main() {
	modFlags := utils.NewModFlags() // manager for shift, ctrl, opt, cmd

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		item := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		item.Retain()
		item.Button().SetTitle("Listening...")

		setupEventMonitor(func(e cocoa.NSEvent) {
			core.Dispatch(func() {
				switch e.Type() {
				case cocoa.NSEventTypeKeyDown:
					keyCode, _ := e.KeyCode()
					item.Button().SetTitle(fmt.Sprintf("Key Down: %d", keyCode))
				case cocoa.NSEventTypeFlagsChanged:
					_modFlags := e.Get("modifierFlags").Uint()
					modFlags.Update(_modFlags)
					item.Button().SetTitle(fmt.Sprintf("Modifier Flags: %s", modFlags.Str()))
				default:
				}
			})
		})
	})

	app.Run()
}

func setupEventMonitor(callback func(e cocoa.NSEvent)) {
	var mask uint64 = cocoa.NSEventMaskKeyDown | cocoa.NSEventMaskFlagsChanged
	ch := make(chan cocoa.NSEvent)

	cocoa.NSEvent_GlobalMonitorMatchingMask(mask, ch)

	go func() {
		for e := range ch {
			callback(e)
		}
	}()
}
