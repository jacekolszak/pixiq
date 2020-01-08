package internal

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/jacekolszak/pixiq/keyboard"
)

var keymap = map[glfw.Key]keyboard.Key{
	glfw.KeySpace:        keyboard.Space,
	glfw.KeyApostrophe:   keyboard.Apostrophe,
	glfw.KeyComma:        keyboard.Comma,
	glfw.KeyMinus:        keyboard.Minus,
	glfw.KeyPeriod:       keyboard.Period,
	glfw.KeySlash:        keyboard.Slash,
	glfw.Key0:            keyboard.Zero,
	glfw.Key1:            keyboard.One,
	glfw.Key2:            keyboard.Two,
	glfw.Key3:            keyboard.Three,
	glfw.Key4:            keyboard.Four,
	glfw.Key5:            keyboard.Five,
	glfw.Key6:            keyboard.Six,
	glfw.Key7:            keyboard.Seven,
	glfw.Key8:            keyboard.Eight,
	glfw.Key9:            keyboard.Nine,
	glfw.KeySemicolon:    keyboard.Semicolon,
	glfw.KeyEqual:        keyboard.Equal,
	glfw.KeyA:            keyboard.A,
	glfw.KeyB:            keyboard.B,
	glfw.KeyC:            keyboard.C,
	glfw.KeyD:            keyboard.D,
	glfw.KeyE:            keyboard.E,
	glfw.KeyF:            keyboard.F,
	glfw.KeyG:            keyboard.G,
	glfw.KeyH:            keyboard.H,
	glfw.KeyI:            keyboard.I,
	glfw.KeyJ:            keyboard.J,
	glfw.KeyK:            keyboard.K,
	glfw.KeyL:            keyboard.L,
	glfw.KeyM:            keyboard.M,
	glfw.KeyN:            keyboard.N,
	glfw.KeyO:            keyboard.O,
	glfw.KeyP:            keyboard.P,
	glfw.KeyQ:            keyboard.Q,
	glfw.KeyR:            keyboard.R,
	glfw.KeyS:            keyboard.S,
	glfw.KeyT:            keyboard.T,
	glfw.KeyU:            keyboard.U,
	glfw.KeyV:            keyboard.V,
	glfw.KeyW:            keyboard.W,
	glfw.KeyX:            keyboard.X,
	glfw.KeyY:            keyboard.Y,
	glfw.KeyZ:            keyboard.Z,
	glfw.KeyLeftBracket:  keyboard.LeftBracket,
	glfw.KeyBackslash:    keyboard.Backslash,
	glfw.KeyRightBracket: keyboard.RightBracket,
	glfw.KeyGraveAccent:  keyboard.GraveAccent,
	glfw.KeyWorld1:       keyboard.World1,
	glfw.KeyWorld2:       keyboard.World2,
	glfw.KeyEscape:       keyboard.Esc,
	glfw.KeyEnter:        keyboard.Enter,
	glfw.KeyTab:          keyboard.Tab,
	glfw.KeyBackspace:    keyboard.Backspace,
	glfw.KeyInsert:       keyboard.Insert,
	glfw.KeyDelete:       keyboard.Delete,
	glfw.KeyRight:        keyboard.Right,
	glfw.KeyLeft:         keyboard.Left,
	glfw.KeyDown:         keyboard.Down,
	glfw.KeyUp:           keyboard.Up,
	glfw.KeyPageUp:       keyboard.PageUp,
	glfw.KeyPageDown:     keyboard.PageDown,
	glfw.KeyHome:         keyboard.Home,
	glfw.KeyEnd:          keyboard.End,
	glfw.KeyCapsLock:     keyboard.CapsLock,
	glfw.KeyScrollLock:   keyboard.ScrollLock,
	glfw.KeyNumLock:      keyboard.NumLock,
	glfw.KeyPrintScreen:  keyboard.PrintScreen,
	glfw.KeyPause:        keyboard.Pause,
	glfw.KeyF1:           keyboard.F1,
	glfw.KeyF2:           keyboard.F2,
	glfw.KeyF3:           keyboard.F3,
	glfw.KeyF4:           keyboard.F4,
	glfw.KeyF5:           keyboard.F5,
	glfw.KeyF6:           keyboard.F6,
	glfw.KeyF7:           keyboard.F7,
	glfw.KeyF8:           keyboard.F8,
	glfw.KeyF9:           keyboard.F9,
	glfw.KeyF10:          keyboard.F10,
	glfw.KeyF11:          keyboard.F11,
	glfw.KeyF12:          keyboard.F12,
	glfw.KeyF13:          keyboard.F13,
	glfw.KeyF14:          keyboard.F14,
	glfw.KeyF15:          keyboard.F15,
	glfw.KeyF16:          keyboard.F16,
	glfw.KeyF17:          keyboard.F17,
	glfw.KeyF18:          keyboard.F18,
	glfw.KeyF19:          keyboard.F19,
	glfw.KeyF20:          keyboard.F20,
	glfw.KeyF21:          keyboard.F21,
	glfw.KeyF22:          keyboard.F22,
	glfw.KeyF23:          keyboard.F23,
	glfw.KeyF24:          keyboard.F24,
	glfw.KeyF25:          keyboard.F25,
	glfw.KeyKP0:          keyboard.Keypad0,
	glfw.KeyKP1:          keyboard.Keypad1,
	glfw.KeyKP2:          keyboard.Keypad2,
	glfw.KeyKP3:          keyboard.Keypad3,
	glfw.KeyKP4:          keyboard.Keypad4,
	glfw.KeyKP5:          keyboard.Keypad5,
	glfw.KeyKP6:          keyboard.Keypad6,
	glfw.KeyKP7:          keyboard.Keypad7,
	glfw.KeyKP8:          keyboard.Keypad8,
	glfw.KeyKP9:          keyboard.Keypad9,
	glfw.KeyKPDecimal:    keyboard.KeypadDecimal,
	glfw.KeyKPDivide:     keyboard.KeypadDivide,
	glfw.KeyKPMultiply:   keyboard.KeypadMultiply,
	glfw.KeyKPSubtract:   keyboard.KeypadSubtract,
	glfw.KeyKPAdd:        keyboard.KeypadAdd,
	glfw.KeyKPEnter:      keyboard.KeypadEnter,
	glfw.KeyKPEqual:      keyboard.KeypadEqual,
	glfw.KeyLeftShift:    keyboard.LeftShift,
	glfw.KeyLeftControl:  keyboard.LeftControl,
	glfw.KeyLeftAlt:      keyboard.LeftAlt,
	glfw.KeyLeftSuper:    keyboard.LeftSuper,
	glfw.KeyRightShift:   keyboard.RightShift,
	glfw.KeyRightControl: keyboard.RightControl,
	glfw.KeyRightAlt:     keyboard.RightAlt,
	glfw.KeyRightSuper:   keyboard.RightSuper,
	glfw.KeyMenu:         keyboard.Menu,
}
