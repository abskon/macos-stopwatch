package main

import (
	"time"

	u "github.com/altsko/speedrun-timer/utils"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

// key/mod list: utils/utils.go
var KEYS map[State][]any = map[State][]any{ // MOD uint64, KEY int
	Ready: {u.MOD_SHIFT, u.MOD_CTRL, u.MOD_OPT},
	Start: {u.MOD_CMD},
	Stop:  {u.KEY_SPACE},
}

type State int

const (
	Ready State = iota
	Start
	Stop
)

func main() {
	state := make(chan State, 1)
	ks := u.NewKeySeq()
	timer := u.NewTimer()

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		item := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		item.Retain()
		item.Button().SetFont_(customFont("Menlo", 16))
		item.Button().SetTitle(timer.Str())

		quit := make(chan struct{})

		go func() {
			ticker := time.NewTicker(11 * time.Millisecond) // refresh ui every 11ms (90fps)
			defer ticker.Stop()                             // ui is also updated when state changes

			for {
				select {
				case <-ticker.C:
					if timer.IsRunning() {
						core.Dispatch(func() {
							item.Button().SetTitle(timer.Str())
						})
					}
				case newState := <-state:
					switch newState {
					case Ready:
						timer.Reset()
					case Start:
						timer.Start()
					case Stop:
						timer.Stop()
					}

					core.Dispatch(func() {
						item.Button().SetTitle(timer.Str())
					})
				case <-quit:
					return
				}
			}
		}()

		eventMonitor(func(e cocoa.NSEvent) {
			updateState(state, ks, e)
		})
		setupMenu(item, quit)
	})

	app.Run()
}

func setupMenu(item cocoa.NSStatusItem, quit chan<- struct{}) {
	menu := cocoa.NSMenu_New()
	quitItem := cocoa.NSMenuItem_New()
	quitItem.SetTitle("Quit")
	quitItem.SetAction(objc.Sel("terminate:"))

	menu.AddItem(quitItem)
	item.SetMenu(menu)
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
