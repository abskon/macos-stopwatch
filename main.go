package main

import (
	"github.com/altsko/speedrun-timer/utils"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

type State int

const (
	Ready    State = iota // 0
	Working               // 1
	Finished              // 2
)

var KEYS map[State]any = map[State]any{ // MOD-uint64 KEY-int
	Ready:    utils.MOD_SHIFT,
	Working:  utils.MOD_CMD,
	Finished: utils.KEY_SPACE,
}

func main() {
	// modFlags := utils.NewModFlags() // manager for shift, ctrl, opt, cmd
	// state := Ready

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		item := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		item.Retain()
		item.Button().SetTitle("Listening...")

		eventMonitor(func(e cocoa.NSEvent) {
		})
	})
	app.Run()
}

func eventMonitor(callback func(e cocoa.NSEvent)) {
	var mask uint64 = cocoa.NSEventMaskKeyDown | cocoa.NSEventMaskFlagsChanged
	ch := make(chan cocoa.NSEvent)

	cocoa.NSEvent_GlobalMonitorMatchingMask(mask, ch)

	go func() {
		for e := range ch {
			callback(e)
		}
	}()
}

func updateState(state *State, m *utils.ModFlags, e cocoa.NSEvent) {
	var k any

	switch e.Type() {
	case cocoa.NSEventTypeKeyDown:
		keyCode, _ := e.KeyCode()
		k = keyCode
	case cocoa.NSEventTypeFlagsChanged:
		_modFlags := e.Get("modifierFlags").Uint()
		m.Update(_modFlags)
		k = m
	}

	for targetState, key := range KEYS {
		switch key := key.(type) {
		case int:
			if keyCode, ok := k.(int64); ok && keyCode == int64(key) {
				*state = targetState
				return
			}
		case uint64:
			if modFlags, ok := k.(*utils.ModFlags); ok && modFlags.IsPressed(key) {
				*state = targetState
				return
			}
		}
	}
}
