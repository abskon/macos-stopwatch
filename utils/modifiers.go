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

type ModSeq struct {
	State map[uint64]bool
}

type Mods struct {
	mods uint64
	seq  ModSeq
}

func NewMods() *Mods {
	return &Mods{
		seq: ModSeq{
			State: map[uint64]bool{
				MOD_SHIFT: false,
				MOD_CTRL:  false,
				MOD_OPT:   false,
				MOD_CMD:   false,
			},
		},
	}
}

func (m *Mods) IsPressed(modKey uint64) bool {
	pressed, exists := m.seq.State[modKey]
	return pressed && exists
}

func (m *Mods) GetPressed() []uint64 {
	var pressed []uint64

	for modKey, isPressed := range m.seq.State {
		if isPressed {
			pressed = append(pressed, modKey)
		}
	}

	return pressed
}

func (m *Mods) Update(mods uint64) {
	m.mods = mods

	for modKey := range m.seq.State {
		m.seq.State[modKey] = m.mods&modKey != 0
	}
}

func (m *Mods) Str() string {
	var parts []string

	for modKey, pressed := range m.seq.State {
		if pressed {
			parts = append(parts, modNames[modKey])
		}
	}

	if len(parts) == 0 {
		return "None"
	}

	return strings.Join(parts, "+")
}
