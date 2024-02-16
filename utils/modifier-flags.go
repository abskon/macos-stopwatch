package utils

import (
	"strings"
)

type ModFlags struct {
	modFlags uint64
}

func NewModFlags() *ModFlags {
	return &ModFlags{}
}

func (m *ModFlags) Update(modFlags uint64) {
	m.modFlags = modFlags
}

func (m *ModFlags) Pressed(modKey uint64) bool {
	return m.modFlags&modKey != 0
}

func (m *ModFlags) Str() string {
	var parts []string
	if m.Pressed(MOD_SHIFT) {
		parts = append(parts, "Shift")
	}
	if m.Pressed(MOD_CTRL) {
		parts = append(parts, "Control")
	}
	if m.Pressed(MOD_OPT) {
		parts = append(parts, "Option")
	}
	if m.Pressed(MOD_CMD) {
		parts = append(parts, "Command")
	}
	if len(parts) == 0 {
		return "None"
	}
	return strings.Join(parts, "+")
}
