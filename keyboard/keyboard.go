package keyboard

// EventsSource is a source of keyboard Events.
type EventsSource interface {
	// Fills output slice with events and returns number of events.
	Poll(output []Event) int
}

// UnknownKey returns instance of unknown Key.
// ScanCode is platform-specific code.
func UnknownKey(scanCode int) Key {
	return Key{}
}

// NewKey returns new instance of immutable Key.
func NewKey(token Token, scanCode int) Key {
	return Key{}
}

// NewReleasedEvent returns new instance of Event when key was released.
func NewReleasedEvent(key Key) Event {
	return Event{}
}

// NewPressedEvent returns new instance of Event when key was pressed.
func NewPressedEvent(key Key) Event {
	return Event{}
}

// Event describes what happened with the key. Whether it was pressed or released.
type Event struct {
	typ eventType
	key Key
}

// eventType is used because using polymorphism means heap allocation and we don't
// want to generate garbage (really? StarCraft e-sport players can perform up to 300 APM,
// which means 600 event objects per minute - maybe it is not that much).
type eventType byte

const (
	pressed  eventType = 1
	released eventType = 2
)

// Key contains numbers identifying the key.
type Key struct {
	token    Token
	scanCode int
}

// Token is platform-independent number identifying the key. It may be
// Unknown, then ScanCode should be used instead.
type Token int

const (
	unknown Token = 0
	// A is 65
	A Token = 65
)

// New creates Keyboard instance.
func New(source EventsSource) *Keyboard {
	return nil
}

// Keyboard provides a read-only information about the current state of the
// keyboard, such as what keys are currently pressed.
type Keyboard struct {
}

// Update updates the state of the keyboard by polling events queued since last time
// the function was executed.
func (k *Keyboard) Update() {

}

// Pressed returns true if given key is currently pressed.
func (k Keyboard) Pressed(keyToken Token) bool {
	return false
}
