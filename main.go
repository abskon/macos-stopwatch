package main

import (
	"fmt"
	"strings"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

const characterToListen = " "

func main() {
    cocoa.TerminateAfterWindowsClose = false
    app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
        // Initialize the menu bar item with a starting title of "0".
        item := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
        item.Retain()
        item.Button().SetTitle("0")

        setupEventMonitor(func(e cocoa.NSEvent) {
            characters, err := e.Characters()
            if err != nil {
                fmt.Println(err)
            }

            if len(characters) > 0 {
                fmt.Println(characters)
            }

            if strings.Contains(characters, characterToListen) {
                core.Dispatch(func() {
                    currentTitle := item.Button().Title()
                    var count int
                    fmt.Sscanf(currentTitle, "%d", &count)
                    count++
                    item.Button().SetTitle(fmt.Sprintf("%d", count))
                })
            }
        })
    })

    app.Run()
}

func setupEventMonitor(callback func(e cocoa.NSEvent)) {
    var mask uint64 = cocoa.NSEventMaskKeyDown
    ch := make(chan cocoa.NSEvent)

    cocoa.NSEvent_GlobalMonitorMatchingMask(mask, ch)

    go func() {
        for e := range ch {
            callback(e)
        }
    }()
}
