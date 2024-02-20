package ui

import "github.com/progrium/macdriver/objc"

type AppController struct {
	objc.Object
	Tb *TextBox
}

func (c *AppController) Open(sender objc.Object) {
	c.Tb.Open(sender)
}

func (c *AppController) Close(sender objc.Object) {
	c.Tb.Close(sender)
}
