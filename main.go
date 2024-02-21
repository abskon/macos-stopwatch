package main

import (
	"time"

	"github.com/altsko/macos-stopwatch/ui"
	u "github.com/altsko/macos-stopwatch/utils"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

// key/mod list: utils/utils.go
// MOD uint64, KEY int
var KEYS map[State][]any = map[State][]any{
	Ready: {u.MOD_SHIFT, u.MOD_CTRL, u.MOD_OPT}, // shift+ctrl+opt
	Start: {u.MOD_CTRL, u.MOD_OPT, u.MOD_CMD},   // ctrl+opt+cmd
	Stop:  {nil},                                // none
}

type State int

const (
	Ready State = iota
	Start
	Stop
)

func main() {
	fontName := "Jetbrains Mono"

	state := make(chan State, 1)
	ks := u.NewKeySeq()
	sw := u.NewStopwatch()

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		item := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		item.Retain()

		item.Button().SetFont_(customFont(fontName, 14))
		item.Button().SetTitle(sw.Str())

		tb := ui.NewTextBox(sw.Str(), fontName)

		updateUI := func() {
			core.Dispatch(func() {
				item.Button().SetTitle(sw.Str())
				tb.SetString(sw.Str())
			})
		}

		quit := make(chan struct{})
		go func() {
			ticker := time.NewTicker(3 * time.Millisecond) // refresh ui every 3ms
			defer ticker.Stop()                            // ui is also updated when state changes

			for {
				select {
				case <-ticker.C:
					if sw.IsRunning() {
						updateUI()
					}
				case newState := <-state:
					switch newState {
					case Ready:
						sw.Reset()
					case Start:
						sw.Start()
					case Stop:
						sw.Stop()
					}

					updateUI()
				case <-quit:
					return
				}
			}
		}()

		eventMonitor(func(e cocoa.NSEvent) {
			updateState(state, ks, e)
		})

		menu := cocoa.NSMenu_New()
		quitItem := cocoa.NSMenuItem_New()
		quitItem.SetTitle("Quit")
		quitItem.SetAction(objc.Sel("terminate:"))

		openItem := cocoa.NSMenuItem_New()
		openItem.SetTitle("Open")
		openItem.SetAction(objc.Sel("open:"))

		menu.AddItem(quitItem)
		menu.AddItem(openItem)
		item.SetMenu(menu)

		cocoa.NSApp().SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
		cocoa.NSApp().ActivateIgnoringOtherApps(true)
		cocoa.NSApp().Run()
	})
	app.Run()
}

func updateState(s chan<- State, ks *u.KeySeq, e cocoa.NSEvent) {
	switch e.Type() {
	case cocoa.NSEventTypeKeyDown:
		key, _ := e.KeyCode()
		keyInt := int(key)
		ks.Update(u.Mods_Key{Key: &keyInt})
	case cocoa.NSEventTypeFlagsChanged:
		mod := e.Get("modifierFlags").Uint()
		ks.Update(u.Mods_Key{Mods: &mod})
	case cocoa.NSEventTypeKeyUp:
		none := u.KEY_NONE
		ks.Update(u.Mods_Key{Key: &none})
	}

	// fmt.Println(ks.Str())

	for state, keys := range KEYS {
		if ks.Pressed(keys...) {
			s <- state
		}
	}
}

func eventMonitor(callback func(e cocoa.NSEvent)) {
	var mask uint64 = cocoa.NSEventMaskKeyDown | cocoa.NSEventMaskFlagsChanged | cocoa.NSEventMaskKeyUp
	ch := make(chan cocoa.NSEvent)

	cocoa.NSEvent_GlobalMonitorMatchingMask(mask, ch)

	go func() {
		for e := range ch {
			callback(e)
		}
	}()
}

func customFont(name string, size float64) cocoa.NSFontRef {
	nsFontClass := cocoa.NSFont_Init(name, size)
	return cocoa.NSFont_fromRef(nsFontClass.Send("retain"))
}
