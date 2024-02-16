package main

import (
	"fmt"
	"unsafe"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

func main() {
    cocoa.TerminateAfterWindowsClose = false
    app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
        item := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
        item.Retain()
        item.Button().SetTitle("0") // Start with 0

        setupGlobalEventMonitorForControlKeyPress(func() {
            currentTitle := item.Button().Title()
            var count int
            fmt.Sscanf(currentTitle, "%d", &count)
            count++
            core.Dispatch(func() {
                item.Button().SetTitle(fmt.Sprintf("%d", count))
            })
        })
    })

    app.Run()
}

func setupGlobalEventMonitorForControlKeyPress(onControlKeyPress func()) {
    var mask uint64 = cocoa.NSEventMaskFlagsChanged
    ch := make(chan cocoa.NSEvent)
    cocoa.NSEvent_GlobalMonitorMatchingMask(mask, ch)

    go func() {
        var controlKeyDown bool
        for e := range ch {
            flags := uint64(e.Get("modifierFlags").Uint())
            controlFlag := uint64(1 << 18) // Control key flag
            
            // Check if the Control key flag is present in the flags
            if flags&controlFlag == controlFlag {
                if !controlKeyDown { // Control key was pressed down
                    controlKeyDown = true
                    onControlKeyPress()
                }
            } else {
                if controlKeyDown { // Control key was released
                    controlKeyDown = false
                }
            }
        }
    }()
}

// ModifierFlags returns the modifier flags of an NSEvent.
func (e cocoa.NSEvent) ModifierFlags() uint64 {
    return *(*uint64)(unsafe.Pointer(e.Get("modifierFlags").Pointer()))
}
