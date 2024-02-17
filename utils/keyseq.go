package utils

import (
	"strings"
)

type Seq struct {
	modState map[uint64]bool
	key      int
}

type KeySeq struct {
	mods uint64
	seq  Seq
}

func NewKeySeq() *KeySeq {
	return &KeySeq{
		seq: Seq{
			modState: map[uint64]bool{
				MOD_SHIFT: false,
				MOD_CTRL:  false,
				MOD_OPT:   false,
				MOD_CMD:   false,
			},
			key: KEY_NONE,
		},
	}
}

type Mods_Key struct {
	Mods *uint64
	Key  *int
}

func (ks *KeySeq) Update(mk Mods_Key) {
	if mk.Mods != nil {
		ks.mods = *mk.Mods
		for modKey := range ks.seq.modState {
			ks.seq.modState[modKey] = ks.mods&modKey != 0
		}
	}

	if mk.Key != nil {
		ks.seq.key = *mk.Key
	}
}

func (ks *KeySeq) Pressed(keys ...any) bool {
	pressed := []any{}
	for _, key := range keys {
		if ks.isPressed(key) {
			pressed = append(pressed, key)
		}
	}
	return len(pressed) == len(keys)
}

func (ks *KeySeq) isPressed(key any) bool {
	switch k := key.(type) {
	case uint64:
		pressed, exists := ks.seq.modState[k]
		return pressed && exists
	case int:
		return ks.seq.key == k
	}
	return false
}

func (ks *KeySeq) Str() string {
	var parts []string

	modOrder := []uint64{MOD_SHIFT, MOD_CTRL, MOD_OPT, MOD_CMD}

	for _, mod := range modOrder {
		if pressed, exists := ks.seq.modState[mod]; exists && pressed {
			if modName, ok := modNames[mod]; ok {
				parts = append(parts, modName)
			}
		}
	}

	if ks.seq.key != KEY_NONE {
		if keyName, exists := keyNames[ks.seq.key]; exists {
			parts = append(parts, keyName)
		}
	}

	if len(parts) == 0 {
		return keyNames[KEY_NONE]
	}

	return strings.Join(parts, "+")
}
