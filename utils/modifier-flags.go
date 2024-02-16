package utils

import "strings"

const (
	ShiftKeyModifier   uint64 = 0x20000
	ControlKeyModifier uint64 = 0x40000
	OptionKeyModifier  uint64 = 0x80000
	CommandKeyModifier uint64 = 0x100000
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
	if m.Pressed(ShiftKeyModifier) {
		parts = append(parts, "Shift")
	}
	if m.Pressed(ControlKeyModifier) {
		parts = append(parts, "Control")
	}
	if m.Pressed(OptionKeyModifier) {
		parts = append(parts, "Option")
	}
	if m.Pressed(CommandKeyModifier) {
		parts = append(parts, "Command")
	}
	if len(parts) == 0 {
		return "None"
	}
	return strings.Join(parts, "+")
}
