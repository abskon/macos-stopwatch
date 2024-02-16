package utils

import (
	"strings"
)

var modNames = map[uint64]string{
	MOD_SHIFT: "Shift",
	MOD_CTRL:  "Ctrl",
	MOD_OPT:   "Opt",
	MOD_CMD:   "Cmd",
}

type ModStroke struct {
	State map[uint64]bool
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

	for modKey := range m.stroke.State {
		m.stroke.State[modKey] = m.flags&modKey != 0
	}
}

func (m *ModFlags) Pressed(modKey uint64) bool {
	return m.flags&modKey != 0
}

func (m *ModFlags) Str() string {
	var parts []string

	for modKey, pressed := range m.stroke.State {
		if pressed {
			parts = append(parts, modNames[modKey])
		}
	}

	if len(parts) == 0 {
		return "None"
	}

	return strings.Join(parts, "+")
}
