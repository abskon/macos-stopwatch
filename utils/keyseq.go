package utils

import (
	"strings"
)

type Seq struct {
	ModState map[uint64]bool
	Keys     []int
}

type KeySeq struct {
	mods uint64
	seq  Seq
}

func NewKeySeq() *KeySeq {
	return &KeySeq{
		seq: Seq{
			ModState: map[uint64]bool{
				MOD_SHIFT: false,
				MOD_CTRL:  false,
				MOD_OPT:   false,
				MOD_CMD:   false,
			},
			Keys: []int{},
		},
	}
}

func (ks *KeySeq) IsPressed(key any) bool {
	switch k := key.(type) {
	case uint64:
		pressed, exists := ks.seq.ModState[k]
		return pressed && exists
	case int:
		for _, key := range ks.seq.Keys {
			if key == k {
				return true
			}
		}
	}
	return false
}

func (ks *KeySeq) ArePressed(keys ...any) bool {
	pressed := []any{}
	for _, key := range keys {
		if ks.IsPressed(key) {
			pressed = append(pressed, key)
		}
	}
	return len(pressed) == len(keys)
}

type Mods_Key struct {
	Mods *uint64
	Key  *int
}

func (ks *KeySeq) Update(mk Mods_Key) {
	if mk.Mods != nil {
		ks.mods = *mk.Mods
		for modKey := range ks.seq.ModState {
			ks.seq.ModState[modKey] = ks.mods&modKey != 0
		}
	}

	if mk.Key != nil {
		inKeys := false
		for i, key := range ks.seq.Keys {
			if key == *mk.Key {
				ks.seq.Keys = append(ks.seq.Keys[:i], ks.seq.Keys[i+1:]...)
				inKeys = true
				break
			}
		}
		if !inKeys {
			ks.seq.Keys = append(ks.seq.Keys, *mk.Key)
		}
	}
}

func (ks *KeySeq) Str() string {
	var parts []string

	for modKey, pressed := range ks.seq.ModState {
		if pressed {
			parts = append(parts, modNames[modKey])
		}
	}

	if len(parts) == 0 {
		return "None"
	}

	return strings.Join(parts, "+")
}
