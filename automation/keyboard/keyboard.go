package keyboard

import (
	"github.com/go-vgo/robotgo"
	"github.com/haroflow/go-macros/automation"
	"github.com/haroflow/go-macros/fakerinput"
)

func Commands() []automation.Command {
	moduleName := "keyboard"
	return []automation.Command{
		{
			ModuleName:  moduleName,
			MethodName:  "type",
			Parameters:  "message: string",
			Description: "Types a message on the keyboard.",
			Action:      Type,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "press",
			Parameters:  "key: string, ...",
			Description: "Presses one or more keys simultaneously. For a complete list of available keys, see: https://github.com/go-vgo/robotgo/blob/master/docs/keys.md",
			Action:      Press,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "down",
			Parameters:  "key: string",
			Description: "Presses a key down.",
			Action:      Down,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "up",
			Parameters:  "key: string",
			Description: "Releases a key.",
			Action:      Up,
		},
	}
}

func Type(msg string) {
	robotgo.TypeStr(msg)
}

func Down(key string) {
	if fakerinput.Enabled {
		kbdReport(func(r *fakerinput.KBDReport) {
			if k, ok := fakerKeys[key]; ok {
				r.KeyDown(k)
			} else if c, ok := fakerControls[key]; ok {
				r.ModifierDown(c)
			}
		})
		return
	}

	_ = robotgo.KeyToggle(key, "down")
}

func Up(key string) {
	if fakerinput.Enabled {
		kbdReport(func(r *fakerinput.KBDReport) {
			if k, ok := fakerKeys[key]; ok {
				r.KeyUp(k)
			} else if c, ok := fakerControls[key]; ok {
				r.ModifierUp(c)
			}
		})
		return
	}

	_ = robotgo.KeyToggle(key, "up")
}

func Press(key string, other ...string) {
	if fakerinput.Enabled {
		kbdReport(func(r *fakerinput.KBDReport) {
			if k, ok := fakerKeys[key]; ok {
				r.KeyDown(k)
			} else if c, ok := fakerControls[key]; ok {
				r.ModifierDown(c)
			}
		})
		robotgo.MilliSleep(3)
		kbdReport(func(r *fakerinput.KBDReport) {
			if k, ok := fakerKeys[key]; ok {
				r.KeyUp(k)
			} else if c, ok := fakerControls[key]; ok {
				r.ModifierUp(c)
			}
		})
		return
	}

	//https://github.com/go-vgo/robotgo/blob/master/docs/keys.md
	if len(other) > 0 {
		_ = robotgo.KeyTap(key, other)
	} else {
		_ = robotgo.KeyTap(key)
	}
}

func kbdReport(cb func(*fakerinput.KBDReport)) {
	r := fakerinput.NewKBDReport()
	cb(r)
	fakerinput.FI.UpdateKeyboard(r)
}

var fakerKeys = map[string]fakerinput.KeyboardKey{
	"0":           fakerinput.KeyNumber0,
	"1":           fakerinput.KeyNumber1,
	"2":           fakerinput.KeyNumber2,
	"3":           fakerinput.KeyNumber3,
	"4":           fakerinput.KeyNumber4,
	"5":           fakerinput.KeyNumber5,
	"6":           fakerinput.KeyNumber6,
	"7":           fakerinput.KeyNumber7,
	"8":           fakerinput.KeyNumber8,
	"9":           fakerinput.KeyNumber9,
	"a":           fakerinput.KeyA,
	"b":           fakerinput.KeyB,
	"c":           fakerinput.KeyC,
	"d":           fakerinput.KeyD,
	"e":           fakerinput.KeyE,
	"f":           fakerinput.KeyF,
	"g":           fakerinput.KeyG,
	"h":           fakerinput.KeyH,
	"i":           fakerinput.KeyI,
	"j":           fakerinput.KeyJ,
	"k":           fakerinput.KeyK,
	"l":           fakerinput.KeyL,
	"m":           fakerinput.KeyM,
	"n":           fakerinput.KeyN,
	"o":           fakerinput.KeyO,
	"p":           fakerinput.KeyP,
	"q":           fakerinput.KeyQ,
	"r":           fakerinput.KeyR,
	"s":           fakerinput.KeyS,
	"t":           fakerinput.KeyT,
	"u":           fakerinput.KeyU,
	"v":           fakerinput.KeyV,
	"w":           fakerinput.KeyW,
	"x":           fakerinput.KeyX,
	"y":           fakerinput.KeyY,
	"z":           fakerinput.KeyZ,
	"enter":       fakerinput.KeyEnter,
	"esc":         fakerinput.KeyEscape,
	"escape":      fakerinput.KeyEscape,
	"backspace":   fakerinput.KeyBackspace,
	"space":       fakerinput.KeySpace,
	"tab":         fakerinput.KeyTab,
	"capslock":    fakerinput.KeyCapsLock,
	"-":           fakerinput.KeyMinus,
	"=":           fakerinput.KeyEquals,
	"[":           fakerinput.KeyLeftBracket,
	"]":           fakerinput.KeyRightBracket,
	"\\":          fakerinput.KeyBackslash,
	";":           fakerinput.KeySemicolon,
	"'":           fakerinput.KeyApostrophe,
	"`":           fakerinput.KeyGrave,
	",":           fakerinput.KeyComma,
	".":           fakerinput.KeyPeriod,
	"/":           fakerinput.KeySlash,
	"printscreen": fakerinput.KeyPrintScreen,
	"scrolllock":  fakerinput.KeyScrollLock,
	"pause":       fakerinput.KeyPause,
	"insert":      fakerinput.KeyInsert,
	"home":        fakerinput.KeyHome,
	"pageup":      fakerinput.KeyPageUp,
	"delete":      fakerinput.KeyDelete,
	"end":         fakerinput.KeyEnd,
	"pagedown":    fakerinput.KeyPageDown,
	"right":       fakerinput.KeyRightArrow,
	"left":        fakerinput.KeyLeftArrow,
	"down":        fakerinput.KeyDownArrow,
	"up":          fakerinput.KeyUpArrow,
	"num_lock":    fakerinput.KeyNumLock,
	"num0":        fakerinput.KeyKeypad0,
	"num1":        fakerinput.KeyKeypad1,
	"num2":        fakerinput.KeyKeypad2,
	"num3":        fakerinput.KeyKeypad3,
	"num4":        fakerinput.KeyKeypad4,
	"num5":        fakerinput.KeyKeypad5,
	"num6":        fakerinput.KeyKeypad6,
	"num7":        fakerinput.KeyKeypad7,
	"num8":        fakerinput.KeyKeypad8,
	"num9":        fakerinput.KeyKeypad9,
	"num/":        fakerinput.KeyKeypadDivide,
	"num*":        fakerinput.KeyKeypadMultiply,
	"num-":        fakerinput.KeyKeypadSubtract,
	"num+":        fakerinput.KeyKeypadAdd,
	"num_enter":   fakerinput.KeyKeypadEnter,
	"num.":        fakerinput.KeyKeypadDecimal,
	"f1":          fakerinput.KeyF1,
	"f2":          fakerinput.KeyF2,
	"f3":          fakerinput.KeyF3,
	"f4":          fakerinput.KeyF4,
	"f5":          fakerinput.KeyF5,
	"f6":          fakerinput.KeyF6,
	"f7":          fakerinput.KeyF7,
	"f8":          fakerinput.KeyF8,
	"f9":          fakerinput.KeyF9,
	"f10":         fakerinput.KeyF10,
	"f11":         fakerinput.KeyF11,
	"f12":         fakerinput.KeyF12,
	"f13":         fakerinput.KeyF13,
	"f14":         fakerinput.KeyF14,
	"f15":         fakerinput.KeyF15,
	"f16":         fakerinput.KeyF16,
	"f17":         fakerinput.KeyF17,
	"f18":         fakerinput.KeyF18,
	"f19":         fakerinput.KeyF19,
	"f20":         fakerinput.KeyF20,
	"f21":         fakerinput.KeyF21,
	"f22":         fakerinput.KeyF22,
	"f23":         fakerinput.KeyF23,
	"f24":         fakerinput.KeyF24,
}

var fakerControls = map[string]fakerinput.KBDModifier{
	"ctrl":   fakerinput.ModLControl,
	"lctrl":  fakerinput.ModLControl,
	"rctrl":  fakerinput.ModRControl,
	"alt":    fakerinput.ModLAlt,
	"lalt":   fakerinput.ModLAlt,
	"ralt":   fakerinput.ModRAlt,
	"shift":  fakerinput.ModLShift,
	"lshift": fakerinput.ModLShift,
	"rshift": fakerinput.ModRShift,
}
