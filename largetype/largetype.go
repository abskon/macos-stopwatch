package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

var (
	fontName *string
	text     string
)

func main() {
	fontName = flag.String("font", "Helvetica", "font to use")
	flag.Parse()

	text = strings.Join(flag.Args(), " ")
	if text == "" {
		text = "Hello world"
	}
	text = fmt.Sprintf(" %s ", text)
	fmt.Println(text)

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		screen := cocoa.NSScreen_Main().Frame().Size

		tr, fontSize := func() (rect core.NSRect, size float64) {
			t := cocoa.NSTextView_Init(core.Rect(0, 0, 0, 0))
			t.SetString(text)

			for s := 70.0; s <= 550; s += 12 {
				t.SetFont(cocoa.NSFont_Init(*fontName, s))
				t.LayoutManager().EnsureLayoutForTextContainer(t.TextContainer())
				rect = t.LayoutManager().UsedRectForTextContainer(t.TextContainer())
				size = s
				if rect.Size.Width >= screen.Width*0.8 {
					break
				}
			}
			return rect, size
		}()

		height := tr.Size.Height * 1.5
		tr.Origin.Y = (height / 2) - (tr.Size.Height / 2)
		t := cocoa.NSTextView_Init(tr)
		t.SetString(text)
		t.SetFont(cocoa.NSFont_Init(*fontName, fontSize))
		t.SetEditable(false)
		t.SetImportsGraphics(false)
		t.SetDrawsBackground(false)

		c := cocoa.NSView_Init(core.Rect(0, 0, 0, 0))

		c.SetWantsLayer(true)
		c.Layer().SetCornerRadius(32.0)

		bgColor := cocoa.NSColor_Init(0, 0, 0, 0.75)
		c.SetBackgroundColor(bgColor)
		c.AddSubviewPositionedRelativeTo(t, cocoa.NSWindowAbove, nil)

		tr.Size.Height = height
		tr.Origin.X = (screen.Width / 2) - (tr.Size.Width / 2)
		tr.Origin.Y = (screen.Height / 2) - (tr.Size.Height / 2)

		w := cocoa.NSWindow_Init(tr, cocoa.NSBorderlessWindowMask, cocoa.NSBackingStoreBuffered, false)
		w.Retain()
		w.SetContentView(c)
		w.SetTitlebarAppearsTransparent(true)
		w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		w.SetOpaque(false)
		w.SetBackgroundColor(cocoa.NSColor_Clear())
		w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
		w.SetFrameDisplay(tr, true)
		w.MakeKeyAndOrderFront(nil)

		cocoa.NSApp().SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
		cocoa.NSApp().ActivateIgnoringOtherApps(true)

		cocoa.NSApp().Run()
	})
	app.Run()
}
