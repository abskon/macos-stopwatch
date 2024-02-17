package utils

var modNames = map[uint64]string{
	MOD_SHIFT: "Shift",
	MOD_CTRL:  "Ctrl",
	MOD_OPT:   "Opt",
	MOD_CMD:   "Cmd",
}

var keyNames = map[int]string{
	KEY_A:              "A",
	KEY_S:              "S",
	KEY_D:              "D",
	KEY_F:              "F",
	KEY_H:              "H",
	KEY_G:              "G",
	KEY_Z:              "Z",
	KEY_X:              "X",
	KEY_C:              "C",
	KEY_V:              "V",
	KEY_B:              "B",
	KEY_Q:              "Q",
	KEY_W:              "W",
	KEY_E:              "E",
	KEY_R:              "R",
	KEY_Y:              "Y",
	KEY_T:              "T",
	KEY_1:              "1",
	KEY_2:              "2",
	KEY_3:              "3",
	KEY_4:              "4",
	KEY_6:              "6",
	KEY_5:              "5",
	KEY_EQUAL:          "=",
	KEY_9:              "9",
	KEY_7:              "7",
	KEY_MINUS:          "-",
	KEY_8:              "8",
	KEY_0:              "0",
	KEY_O:              "O",
	KEY_U:              "U",
	KEY_I:              "I",
	KEY_P:              "P",
	KEY_L:              "L",
	KEY_J:              "J",
	KEY_K:              "K",
	KEY_N:              "N",
	KEY_M:              "M",
	KEY_COMMA:          ",",
	KEY_PERIOD:         ".",
	KEY_SLASH:          "/",
	KEY_SEMICOLON:      ";",
	KEY_APOSTROPHE:     "'",
	KEY_BACKSLASH:      "\\",
	KEY_SECTION_SIGN:   "§",
	KEY_F1:             "F1",
	KEY_F2:             "F2",
	KEY_F3:             "F3",
	KEY_F4:             "F4",
	KEY_F5:             "F5",
	KEY_F6:             "F6",
	KEY_F7:             "F7",
	KEY_F8:             "F8",
	KEY_F9:             "F9",
	KEY_F10:            "F10",
	KEY_F11:            "F11",
	KEY_F12:            "F12",
	KEY_ENTER:          "Enter",
	KEY_TAB:            "Tab",
	KEY_SPACE:          "Space",
	KEY_DELETE:         "Delete",
	KEY_ESCAPE:         "Escape",
	KEY_COMMAND:        "Command",
	KEY_SHIFT:          "Shift",
	KEY_CAPSLOCK:       "CapsLock",
	KEY_OPTION:         "Option",
	KEY_CONTROL:        "Control",
	KEY_RIGHTSHIFT:     "RightShift",
	KEY_RIGHTOPTION:    "RightOption",
	KEY_RIGHTCONTROL:   "RightControl",
	KEY_FUNCTION:       "Function",
	KEY_RIGHTARROW:     "RightArrow",
	KEY_LEFTARROW:      "LeftArrow",
	KEY_DOWNARROW:      "DownArrow",
	KEY_UPARROW:        "UpArrow",
	KEY_VOLUME_UP:      "Volume Up",
	KEY_VOLUME_DOWN:    "Volume Down",
	KEY_MUTE:           "Mute",
	KEY_HELP:           "Help",
	KEY_HOME:           "Home",
	KEY_PAGE_UP:        "Page Up",
	KEY_FORWARD_DELETE: "Forward Delete",
	KEY_END:            "End",
	KEY_PAGE_DOWN:      "Page Down",
	KEY_CLEAR:          "Clear",
	KEY_PAD_EQUALS:     "Pad =",
	KEY_PAD_DIVIDE:     "Pad /",
	KEY_PAD_MULTIPLY:   "Pad *",
	KEY_PAD_MINUS:      "Pad -",
	KEY_PAD_PLUS:       "Pad +",
	KEY_PAD_ENTER:      "Pad Enter",
	KEY_PAD_0:          "Pad 0",
	KEY_PAD_1:          "Pad 1",
	KEY_PAD_2:          "Pad 2",
	KEY_PAD_3:          "Pad 3",
	KEY_PAD_4:          "Pad 4",
	KEY_PAD_5:          "Pad 5",
	KEY_PAD_6:          "Pad 6",
	KEY_PAD_7:          "Pad 7",
	KEY_PAD_8:          "Pad 8",
	KEY_PAD_9:          "Pad 9",
	KEY_PAD_DECIMAL:    "Pad .",
}

const (
	MOD_SHIFT uint64 = 0x20000
	MOD_CTRL  uint64 = 0x40000
	MOD_OPT   uint64 = 0x80000
	MOD_CMD   uint64 = 0x100000

	KEY_A            int = 0
	KEY_S            int = 1
	KEY_D            int = 2
	KEY_F            int = 3
	KEY_H            int = 4
	KEY_G            int = 5
	KEY_Z            int = 6
	KEY_X            int = 7
	KEY_C            int = 8
	KEY_V            int = 9
	KEY_B            int = 11
	KEY_Q            int = 12
	KEY_W            int = 13
	KEY_E            int = 14
	KEY_R            int = 15
	KEY_Y            int = 16
	KEY_T            int = 17
	KEY_1            int = 18
	KEY_2            int = 19
	KEY_3            int = 20
	KEY_4            int = 21
	KEY_6            int = 22
	KEY_5            int = 23
	KEY_EQUAL        int = 24
	KEY_9            int = 25
	KEY_7            int = 26
	KEY_MINUS        int = 27
	KEY_8            int = 28
	KEY_0            int = 29
	KEY_O            int = 31
	KEY_U            int = 32
	KEY_I            int = 34
	KEY_P            int = 35
	KEY_L            int = 37
	KEY_J            int = 38
	KEY_K            int = 40
	KEY_N            int = 45
	KEY_M            int = 46
	KEY_COMMA        int = 43
	KEY_PERIOD       int = 47
	KEY_SLASH        int = 44
	KEY_SEMICOLON    int = 41
	KEY_APOSTROPHE   int = 39
	KEY_BACKSLASH    int = 42
	KEY_SECTION_SIGN int = 10

	KEY_F1  int = 122
	KEY_F2  int = 120
	KEY_F3  int = 99
	KEY_F4  int = 118
	KEY_F5  int = 96
	KEY_F6  int = 97
	KEY_F7  int = 98
	KEY_F8  int = 100
	KEY_F9  int = 101
	KEY_F10 int = 109
	KEY_F11 int = 103
	KEY_F12 int = 111

	KEY_ENTER        int = 36
	KEY_TAB          int = 48
	KEY_SPACE        int = 49
	KEY_DELETE       int = 51
	KEY_ESCAPE       int = 53
	KEY_COMMAND      int = 55
	KEY_SHIFT        int = 56
	KEY_CAPSLOCK     int = 57
	KEY_OPTION       int = 58
	KEY_CONTROL      int = 59
	KEY_RIGHTSHIFT   int = 60
	KEY_RIGHTOPTION  int = 61
	KEY_RIGHTCONTROL int = 62
	KEY_FUNCTION     int = 63

	KEY_RIGHTARROW int = 124
	KEY_LEFTARROW  int = 123
	KEY_DOWNARROW  int = 125
	KEY_UPARROW    int = 126

	KEY_VOLUME_UP      int = 72
	KEY_VOLUME_DOWN    int = 73
	KEY_MUTE           int = 74
	KEY_HELP           int = 114
	KEY_HOME           int = 115
	KEY_PAGE_UP        int = 116
	KEY_FORWARD_DELETE int = 117
	KEY_END            int = 119
	KEY_PAGE_DOWN      int = 121
	KEY_CLEAR          int = 71
	KEY_PAD_EQUALS     int = 81
	KEY_PAD_DIVIDE     int = 75
	KEY_PAD_MULTIPLY   int = 67
	KEY_PAD_MINUS      int = 78
	KEY_PAD_PLUS       int = 69
	KEY_PAD_ENTER      int = 76
	KEY_PAD_0          int = 82
	KEY_PAD_1          int = 83
	KEY_PAD_2          int = 84
	KEY_PAD_3          int = 85
	KEY_PAD_4          int = 86
	KEY_PAD_5          int = 87
	KEY_PAD_6          int = 88
	KEY_PAD_7          int = 89
	KEY_PAD_8          int = 91
	KEY_PAD_9          int = 92
	KEY_PAD_DECIMAL    int = 65
)
