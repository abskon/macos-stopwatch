package main

import (
	"time"

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
} // TODO: handle case when same key is used for different states

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
		item.Button().SetFont_(customFont("Menlo", 14))
		item.Button().SetTitle(timer.Str())

		quit := make(chan struct{})

		go func() {
			ticker := time.NewTicker(11 * time.Millisecond)
			defer ticker.Stop()

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
	var k any

	switch e.Type() {
	case cocoa.NSEventTypeKeyDown:
		keyCode, _ := e.KeyCode()
		k = keyCode // int64
	case cocoa.NSEventTypeFlagsChanged:
		mods := e.Get("modifierFlags").Uint()
		ks.Update(u.UpdateOpts{Mods: &mods})
		k = ks // *ModFlags
	}

	for targetState, key := range KEYS {
		switch key := key.(type) {
		case int:
			if keyCode, ok := k.(int64); ok && keyCode == int64(key) {
				s <- targetState
				return
			}
		case uint64:
			if modFlags, ok := k.(*u.KeySeq); ok && modFlags.IsPressed(key) {
				s <- targetState
				return
			}
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

func customFont(name string, size int32) cocoa.NSFontRef {
	nsFontClass := objc.GetClass("NSFont")
	sizeNum := core.NSNumber_WithInt(size)
	font := nsFontClass.Send("fontWithName:size:", core.NSString_FromString(name), sizeNum)
	return cocoa.NSFont_fromRef(font)
}
