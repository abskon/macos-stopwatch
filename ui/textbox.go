package ui

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type TextBox struct {
	view     cocoa.NSView
	textView cocoa.NSTextView
	window   cocoa.NSWindow
}

func NewTextBox(initText string, fontName string) *TextBox {
	fontSize := 48.0

	tv := cocoa.NSTextView_Init(core.Rect(0, 0, 0, 0))
	tv.SetString(initText)
	tv.SetFont(cocoa.NSFont_Init(fontName, fontSize))
	tv.LayoutManager().EnsureLayoutForTextContainer(tv.TextContainer())
	tr := tv.LayoutManager().UsedRectForTextContainer(tv.TextContainer())

	h := tr.Size.Height * 1.25
	tr.Origin.Y = (h / 2) - (tr.Size.Height / 2)
	t := cocoa.NSTextView_Init(tr)
	t.SetString(initText)
	t.SetFont(cocoa.NSFont_Init(fontName, fontSize))
	t.SetEditable(false)
	t.SetImportsGraphics(false)
	t.SetDrawsBackground(false)
	t.SetAlignment(cocoa.NSTextAlignmentRight) // align center

	c := cocoa.NSView_Init(core.Rect(0, 0, 0, 0))
	c.SetWantsLayer(true)
	c.Layer().SetCornerRadius(h / 8)
	bgColor := cocoa.NSColor_Init(0, 0, 0, 0.75)
	c.SetBackgroundColor(bgColor)
	c.AddSubviewPositionedRelativeTo(t, cocoa.NSWindowAbove, nil)

	tr.Size.Height = h
	tr.Origin.X = 64
	tr.Origin.Y = 64

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

	return &TextBox{
		view:     c,
		textView: t,
		window:   w,
	}
}

func (tb *TextBox) SetString(text string) {
	tb.textView.SetString(text)
}

func (tb *TextBox) Open(sender objc.Object) {
	tb.window.MakeKeyAndOrderFront(sender)
}

func (tb *TextBox) Close(sender objc.Object) {
	tb.window.OrderOut(sender)
}
