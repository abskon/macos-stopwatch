package utils

import (
	"strings"
)

type Seq struct {
	ModState map[uint64]bool
	Key      int
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
			Key: -1,
		},
	}
}

func (ks *KeySeq) IsPressed(key any) bool {
	switch k := key.(type) {
	case uint64:
		pressed, exists := ks.seq.ModState[k]
		return pressed && exists
	case int:
		return ks.seq.Key == k
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

type UpdateOpts struct {
	Mods *uint64
	Key  *int
}

// TODO: Update Key when key up (main.go)
func (ks *KeySeq) Update(opts UpdateOpts) {
	if opts.Mods != nil {
		ks.mods = *opts.Mods
		for modKey := range ks.seq.ModState {
			ks.seq.ModState[modKey] = ks.mods&modKey != 0
		}
	}

	if opts.Key != nil {
		ks.seq.Key = *opts.Key
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
