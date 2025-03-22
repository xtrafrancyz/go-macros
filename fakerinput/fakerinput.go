package fakerinput

// #cgo LDFLAGS: -L./ -lFakerInputDll
// #include <stdlib.h>
// #include <stdbool.h>
//
// typedef void* HNDL;
//
// HNDL fakerinput_alloc();
// void fakerinput_free(HNDL vmulti);
// bool fakerinput_connect(HNDL vmulti);
// bool fakerinput_disconnect(HNDL vmulti);
// bool fakerinput_update_keyboard(HNDL vmulti, unsigned char shiftKeyFlags, const unsigned char* keyCodes);
// bool fakerinput_update_keyboard_enhanced(HNDL vmulti, unsigned char mediaKeys, unsigned char enhancedKeys);
// bool fakerinput_update_relative_mouse(HNDL vmulti, unsigned char button, unsigned short x, unsigned short y, unsigned char wheelPosition, unsigned char hWheelPosition);
// bool fakerinput_update_absolute_mouse(HNDL vmulti, unsigned char button, unsigned short x, unsigned short y, unsigned char wheelPosition, unsigned char hWheelPosition);
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

// KeyboardKey represents keyboard key codes
type KeyboardKey uint8

const (
	KeyA                 KeyboardKey = 0x04
	KeyB                 KeyboardKey = 0x05
	KeyC                 KeyboardKey = 0x06
	KeyD                 KeyboardKey = 0x07
	KeyE                 KeyboardKey = 0x08
	KeyF                 KeyboardKey = 0x09
	KeyG                 KeyboardKey = 0x0A
	KeyH                 KeyboardKey = 0x0B
	KeyI                 KeyboardKey = 0x0C
	KeyJ                 KeyboardKey = 0x0D
	KeyK                 KeyboardKey = 0x0E
	KeyL                 KeyboardKey = 0x0F
	KeyM                 KeyboardKey = 0x10
	KeyN                 KeyboardKey = 0x11
	KeyO                 KeyboardKey = 0x12
	KeyP                 KeyboardKey = 0x13
	KeyQ                 KeyboardKey = 0x14
	KeyR                 KeyboardKey = 0x15
	KeyS                 KeyboardKey = 0x16
	KeyT                 KeyboardKey = 0x17
	KeyU                 KeyboardKey = 0x18
	KeyV                 KeyboardKey = 0x19
	KeyW                 KeyboardKey = 0x1A
	KeyX                 KeyboardKey = 0x1B
	KeyY                 KeyboardKey = 0x1C
	KeyZ                 KeyboardKey = 0x1D
	KeyNumber1           KeyboardKey = 0x1E
	KeyNumber2           KeyboardKey = 0x1F
	KeyNumber3           KeyboardKey = 0x20
	KeyNumber4           KeyboardKey = 0x21
	KeyNumber5           KeyboardKey = 0x22
	KeyNumber6           KeyboardKey = 0x23
	KeyNumber7           KeyboardKey = 0x24
	KeyNumber8           KeyboardKey = 0x25
	KeyNumber9           KeyboardKey = 0x26
	KeyNumber0           KeyboardKey = 0x27
	KeyEnter             KeyboardKey = 0x28
	KeyEscape            KeyboardKey = 0x29
	KeyBackspace         KeyboardKey = 0x2A
	KeyTab               KeyboardKey = 0x2B
	KeySpace             KeyboardKey = 0x2C
	KeyMinus             KeyboardKey = 0x2D
	KeyEquals            KeyboardKey = 0x2E
	KeyLeftBracket       KeyboardKey = 0x2F
	KeyRightBracket      KeyboardKey = 0x30
	KeyBackslash         KeyboardKey = 0x31
	KeyNonUSHash         KeyboardKey = 0x32
	KeySemicolon         KeyboardKey = 0x33
	KeyApostrophe        KeyboardKey = 0x34
	KeyGrave             KeyboardKey = 0x35
	KeyComma             KeyboardKey = 0x36
	KeyPeriod            KeyboardKey = 0x37
	KeySlash             KeyboardKey = 0x38
	KeyCapsLock          KeyboardKey = 0x39
	KeyF1                KeyboardKey = 0x3A
	KeyF2                KeyboardKey = 0x3B
	KeyF3                KeyboardKey = 0x3C
	KeyF4                KeyboardKey = 0x3D
	KeyF5                KeyboardKey = 0x3E
	KeyF6                KeyboardKey = 0x3F
	KeyF7                KeyboardKey = 0x40
	KeyF8                KeyboardKey = 0x41
	KeyF9                KeyboardKey = 0x42
	KeyF10               KeyboardKey = 0x43
	KeyF11               KeyboardKey = 0x44
	KeyF12               KeyboardKey = 0x45
	KeyPrintScreen       KeyboardKey = 0x46
	KeyScrollLock        KeyboardKey = 0x47
	KeyPause             KeyboardKey = 0x48
	KeyInsert            KeyboardKey = 0x49
	KeyHome              KeyboardKey = 0x4A
	KeyPageUp            KeyboardKey = 0x4B
	KeyDelete            KeyboardKey = 0x4C
	KeyEnd               KeyboardKey = 0x4D
	KeyPageDown          KeyboardKey = 0x4E
	KeyRightArrow        KeyboardKey = 0x4F
	KeyLeftArrow         KeyboardKey = 0x50
	KeyDownArrow         KeyboardKey = 0x51
	KeyUpArrow           KeyboardKey = 0x52
	KeyNumLock           KeyboardKey = 0x53
	KeyKeypadDivide      KeyboardKey = 0x54
	KeyKeypadMultiply    KeyboardKey = 0x55
	KeyKeypadSubtract    KeyboardKey = 0x56
	KeyKeypadAdd         KeyboardKey = 0x57
	KeyKeypadEnter       KeyboardKey = 0x58
	KeyKeypad1           KeyboardKey = 0x59
	KeyKeypad2           KeyboardKey = 0x5A
	KeyKeypad3           KeyboardKey = 0x5B
	KeyKeypad4           KeyboardKey = 0x5C
	KeyKeypad5           KeyboardKey = 0x5D
	KeyKeypad6           KeyboardKey = 0x5E
	KeyKeypad7           KeyboardKey = 0x5F
	KeyKeypad8           KeyboardKey = 0x60
	KeyKeypad9           KeyboardKey = 0x61
	KeyKeypad0           KeyboardKey = 0x62
	KeyKeypadDecimal     KeyboardKey = 0x63
	KeyKeypadSeparator   KeyboardKey = 0x64
	KeyKeypadApplication KeyboardKey = 0x65
	KeyF13               KeyboardKey = 0x68
	KeyF14               KeyboardKey = 0x69
	KeyF15               KeyboardKey = 0x6A
	KeyF16               KeyboardKey = 0x6B
	KeyF17               KeyboardKey = 0x6C
	KeyF18               KeyboardKey = 0x6D
	KeyF19               KeyboardKey = 0x6E
	KeyF20               KeyboardKey = 0x6F
	KeyF21               KeyboardKey = 0x70
	KeyF22               KeyboardKey = 0x71
	KeyF23               KeyboardKey = 0x72
	KeyF24               KeyboardKey = 0x73
)

// KBDModifier represents keyboard modifier keys
type KBDModifier uint8

const (
	ModLControl KBDModifier = 1
	ModLShift   KBDModifier = 2
	ModLAlt     KBDModifier = 4
	ModLWin     KBDModifier = 8
	ModRControl KBDModifier = 16
	ModRShift   KBDModifier = 32
	ModRAlt     KBDModifier = 64
	ModRWin     KBDModifier = 128
)

// MouseButtons represents mouse button flags
type MouseButton uint8

const (
	MouseLeft     MouseButton = 0x01
	MouseRight    MouseButton = 0x02
	MouseMiddle   MouseButton = 0x04
	MouseXButton1 MouseButton = 0x08
	MouseXButton2 MouseButton = 0x10
)

// KBDReport represents a keyboard report
type KBDReport struct {
	modifiers   map[KBDModifier]struct{}
	pressedKeys map[KeyboardKey]struct{}
}

// NewKBDReport creates a new keyboard report
func NewKBDReport() *KBDReport {
	return &KBDReport{
		modifiers:   make(map[KBDModifier]struct{}),
		pressedKeys: make(map[KeyboardKey]struct{}),
	}
}

// KeyDown adds a key to the pressed keys
func (r *KBDReport) KeyDown(key KeyboardKey) *KBDReport {
	r.pressedKeys[key] = struct{}{}
	return r
}

// KeyUp removes a key from the pressed keys
func (r *KBDReport) KeyUp(key KeyboardKey) *KBDReport {
	delete(r.pressedKeys, key)
	return r
}

// ModifierDown adds a modifier key to the pressed modifiers
func (r *KBDReport) ModifierDown(modifier KBDModifier) *KBDReport {
	r.modifiers[modifier] = struct{}{}
	return r
}

// ModifierUp removes a modifier key from the pressed modifiers
func (r *KBDReport) ModifierUp(modifier KBDModifier) *KBDReport {
	delete(r.modifiers, modifier)
	return r
}

// GetRawShiftKeyFlags returns the raw shift key flags
func (r *KBDReport) GetRawShiftKeyFlags() uint8 {
	var shiftKeys uint8
	for modifier := range r.modifiers {
		shiftKeys |= uint8(modifier)
	}
	return shiftKeys
}

// GetRawKeyCodes returns the raw key codes
func (r *KBDReport) GetRawKeyCodes() [6]uint8 {
	keyCodes := [6]uint8{}
	i := 0
	for key := range r.pressedKeys {
		if i < 6 {
			keyCodes[i] = uint8(key)
			i++
		} else {
			break
		}
	}
	return keyCodes
}

// MouseReport represents a mouse report
type MouseReport struct {
	HeldButtons map[MouseButton]struct{}
	Buttons     uint8
	X           uint16
	Y           uint16
	Wheel       uint8
	HWheel      uint8
}

// NewMouseReport creates a new mouse report
func NewMouseReport() *MouseReport {
	return &MouseReport{
		HeldButtons: make(map[MouseButton]struct{}),
		X:           0,
		Y:           0,
		Wheel:       0,
		HWheel:      0,
		Buttons:     0,
	}
}

// ResetPosition resets the position of the mouse
func (r *MouseReport) ResetPosition() {
	r.X = 0
	r.Y = 0
	r.Wheel = 0
	r.HWheel = 0
}

// ButtonDown presses a mouse button
func (r *MouseReport) ButtonDown(button MouseButton) {
	r.HeldButtons[button] = struct{}{}
	r.Buttons |= uint8(button)
}

// ButtonUp releases a mouse button
func (r *MouseReport) ButtonUp(button MouseButton) {
	delete(r.HeldButtons, button)
	r.Buttons &= ^uint8(button)
}

// Reset resets the mouse report
func (r *MouseReport) Reset() {
	r.HeldButtons = make(map[MouseButton]struct{})
	r.ResetPosition()
	r.Buttons = 0
}

// FakerInput represents a FakerInput device
type FakerInput struct {
	vmulti    C.HNDL
	connected bool
	mu        sync.Mutex
}

// NewFakerInput creates a new FakerInput device
func NewFakerInput() *FakerInput {
	return &FakerInput{
		vmulti:    C.fakerinput_alloc(),
		connected: false,
	}
}

// Connect connects to the FakerInput device
func (f *FakerInput) Connect() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.connected {
		return true
	}

	result := bool(C.fakerinput_connect(f.vmulti))
	f.connected = result
	return result
}

// Disconnect disconnects from the FakerInput device
func (f *FakerInput) Disconnect() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	if !f.connected {
		return true
	}

	result := bool(C.fakerinput_disconnect(f.vmulti))
	f.connected = !result
	return result
}

// UpdateKeyboard updates the keyboard state
func (f *FakerInput) UpdateKeyboard(report *KBDReport) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	if !f.connected {
		return false
	}

	shiftKeyFlags := report.GetRawShiftKeyFlags()
	keyCodes := report.GetRawKeyCodes()

	// Convert Go array to C array
	cKeyCodes := (*C.uchar)(C.malloc(6))
	defer C.free(unsafe.Pointer(cKeyCodes))

	// Copy array data
	for i := 0; i < 6; i++ {
		*(*C.uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(cKeyCodes)) + uintptr(i))) = C.uchar(keyCodes[i])
	}

	return bool(C.fakerinput_update_keyboard(f.vmulti, C.uchar(shiftKeyFlags), cKeyCodes))
}

// UpdateRelativeMouse updates the relative mouse position
func (f *FakerInput) UpdateRelativeMouse(report *MouseReport) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	if !f.connected {
		return false
	}

	return bool(C.fakerinput_update_relative_mouse(
		f.vmulti,
		C.uchar(report.Buttons),
		C.ushort(report.X),
		C.ushort(report.Y),
		C.uchar(report.Wheel),
		C.uchar(report.HWheel),
	))
}

// UpdateAbsoluteMouse updates the absolute mouse position
func (f *FakerInput) UpdateAbsoluteMouse(report *MouseReport) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	if !f.connected {
		return false
	}

	fmt.Printf("%+v\n", report)

	return bool(C.fakerinput_update_absolute_mouse(
		f.vmulti,
		C.uchar(report.Buttons),
		C.ushort(report.X),
		C.ushort(report.Y),
		C.uchar(report.Wheel),
		C.uchar(report.HWheel),
	))
}

// Close frees resources associated with the FakerInput device
func (f *FakerInput) Close() {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.connected {
		C.fakerinput_disconnect(f.vmulti)
		f.connected = false
	}

	C.fakerinput_free(f.vmulti)
	f.vmulti = nil
}
