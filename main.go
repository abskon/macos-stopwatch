package main

import (
	u "github.com/altsko/speedrun-timer/utils"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

// key/mod list: utils/const.go
var KEYS map[State]any = map[State]any{ // MOD uint64, KEY int
	Ready: u.MOD_SHIFT,
	Start: u.MOD_CMD,
	Stop:  u.KEY_SPACE,
}

type State int

const (
	Ready State = iota
	Start
	Stop
)

func main() {
	modFlags := u.NewMods() // manager for shift, ctrl, opt, cmd
	state := make(chan State, 1)

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		item := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		item.Retain()
		item.Button().SetTitle("Ready")

		go func() {
			for newState := range state {
				var newTitle string
				switch newState {
				case Ready:
					newTitle = "Ready"
				case Start:
					newTitle = "Start"
				case Stop:
					newTitle = "Stop"
				}

				core.Dispatch(func() {
					item.Button().SetTitle(newTitle)
				})
			}
		}()

		eventMonitor(func(e cocoa.NSEvent) {
			updateState(state, modFlags, e)
		})
	})
	app.Run()
}

func updateState(s chan<- State, m *u.Mods, e cocoa.NSEvent) {
	var k any

	switch e.Type() {
	case cocoa.NSEventTypeKeyDown:
		keyCode, _ := e.KeyCode()
		k = keyCode // int64
	case cocoa.NSEventTypeFlagsChanged:
		_modFlags := e.Get("modifierFlags").Uint()
		m.Update(_modFlags)
		k = m // *ModFlags
	}

	for targetState, key := range KEYS {
		switch key := key.(type) {
		case int:
			if keyCode, ok := k.(int64); ok && keyCode == int64(key) {
				s <- targetState
				return
			}
		case uint64:
			if modFlags, ok := k.(*u.Mods); ok && modFlags.IsPressed(key) {
				s <- targetState
				return
			}
		}
	}
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
