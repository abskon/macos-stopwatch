package utils

import (
	"strings"
)

type ModStroke struct {
	Shift bool
	Ctrl  bool
	Opt   bool
	Cmd   bool
}

type ModFlags struct {
	flags  uint64
	stroke ModStroke
}

func NewModFlags() *ModFlags {
	return &ModFlags{}
}

func (m *ModFlags) Update(flags uint64) {
	m.flags = flags

	m.stroke.Shift = m.Pressed(MOD_SHIFT)
	m.stroke.Ctrl = m.Pressed(MOD_CTRL)
	m.stroke.Opt = m.Pressed(MOD_OPT)
	m.stroke.Cmd = m.Pressed(MOD_CMD)
}

func (m *ModFlags) Pressed(modKey uint64) bool {
	return m.flags&modKey != 0
}

func (m *ModFlags) Str() string {
	var parts []string
	if m.stroke.Shift {
		parts = append(parts, "Shift")
	}
	if m.stroke.Ctrl {
		parts = append(parts, "Ctrl")
	}
	if m.stroke.Opt {
		parts = append(parts, "Opt")
	}
	if m.stroke.Cmd {
		parts = append(parts, "Cmd")
	}
	if len(parts) == 0 {
		return "None"
	}

	return strings.Join(parts, "+")
}
