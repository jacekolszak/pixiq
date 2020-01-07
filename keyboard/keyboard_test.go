package keyboard_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jacekolszak/pixiq/keyboard"
)

func TestNew(t *testing.T) {
	t.Run("should panic when source is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			keyboard.New(nil)
		})
	})
	t.Run("should create a keyboard instance", func(t *testing.T) {
		source := &fakeEventSource{}
		// when
		keys := keyboard.New(source)
		// then
		assert.NotNil(t, keys)
	})
}

func TestNewKey(t *testing.T) {
	t.Run("should panic for invalid tokens", func(t *testing.T) {
		tests := []keyboard.Token{0, 64, 67}
		for _, token := range tests {
			testName := fmt.Sprintf("for %v", token)
			t.Run(testName, func(t *testing.T) {
				assert.Panics(t, func() {
					keyboard.NewKey(token)
				})
			})
		}
	})
	t.Run("should create key using package variable", func(t *testing.T) {
		key := keyboard.A
		assert.False(t, key.IsUnknown())
		assert.Equal(t, 'A', key.Token().Rune())
	})
	t.Run("should create new key using token", func(t *testing.T) {
		tests := map[string]struct {
			token keyboard.Token
		}{
			"A": {
				token: 65,
			},
			"B": {
				token: 66,
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				key := keyboard.NewKey(test.token)
				assert.Equal(t, test.token, key.Token())
				assert.False(t, key.IsUnknown())
			})
		}
	})
	t.Run("two keys with token should be equal", func(t *testing.T) {
		key1 := keyboard.NewKey(65)
		key2 := keyboard.NewKey(65)
		assert.Equal(t, key1, key2)
	})
	t.Run("two keys with scanCode should be equal", func(t *testing.T) {
		key1 := keyboard.NewUnknownKey(0)
		key2 := keyboard.NewUnknownKey(0)
		assert.Equal(t, key1, key2)
	})
}

func TestNewUnknownKey(t *testing.T) {
	t.Run("should create unknown key using scan code", func(t *testing.T) {
		// when
		key := keyboard.NewUnknownKey(0)
		// then
		assert.True(t, key.IsUnknown())
		assert.Equal(t, 0, key.ScanCode())
	})
}

func TestNewPressedEvent(t *testing.T) {
	t.Run("should create PressedEvent", func(t *testing.T) {
		// when
		keyboard.NewPressedEvent(keyboard.A)
	})
}

func TestKeyboard_Pressed(t *testing.T) {
	t.Run("before Update was called, Pressed returns false for all keys", func(t *testing.T) {
		tests := []keyboard.Key{keyboard.A, keyboard.B}
		for _, key := range tests {
			testName := fmt.Sprintf("for key: %v", key)
			t.Run(testName, func(t *testing.T) {
				source := newFakeEventSourceWithPressedEvents(keyboard.A)
				keys := keyboard.New(source)
				// when
				pressed := keys.Pressed(key)
				// then
				assert.False(t, pressed)
			})
		}
	})
	t.Run("after Update was called", func(t *testing.T) {
		tests := map[string]struct {
			source             keyboard.EventSource
			expectedPressed    []keyboard.Key
			expectedNotPressed []keyboard.Key
		}{
			"one event": {
				source:             newFakeEventSourceWithPressedEvents(keyboard.A),
				expectedPressed:    []keyboard.Key{keyboard.A},
				expectedNotPressed: []keyboard.Key{keyboard.B},
			},
			"two events for keys B and A": {
				source:          newFakeEventSourceWithPressedEvents(keyboard.B, keyboard.A),
				expectedPressed: []keyboard.Key{keyboard.A, keyboard.B},
			},
			"two events for keys A and B": {
				source:          newFakeEventSourceWithPressedEvents(keyboard.A, keyboard.B),
				expectedPressed: []keyboard.Key{keyboard.A, keyboard.B},
			},
			"one event for unknown key": {
				source:             newFakeEventSourceWithPressedEvents(keyboard.NewUnknownKey(0)),
				expectedPressed:    []keyboard.Key{keyboard.NewUnknownKey(0)},
				expectedNotPressed: []keyboard.Key{keyboard.NewUnknownKey(1)},
			},
			"two events: unknown key and A": {
				source:             newFakeEventSourceWithPressedEvents(keyboard.NewUnknownKey(0), keyboard.A),
				expectedPressed:    []keyboard.Key{keyboard.NewUnknownKey(0), keyboard.A},
				expectedNotPressed: []keyboard.Key{keyboard.NewUnknownKey(1)},
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				keys := keyboard.New(test.source)
				// when
				keys.Update()
				// then
				for _, expectedPressedKey := range test.expectedPressed {
					assert.True(t, keys.Pressed(expectedPressedKey))
				}
				for _, expectedNotPressedKey := range test.expectedNotPressed {
					assert.False(t, keys.Pressed(expectedNotPressedKey))
				}
			})
		}
	})
}

func newFakeEventSourceWithPressedEvents(keys ...keyboard.Key) *fakeEventSource {
	source := &fakeEventSource{}
	source.events = []keyboard.Event{}
	for _, key := range keys {
		event := keyboard.NewPressedEvent(key)
		source.events = append(source.events, event)
	}
	return source
}

type fakeEventSource struct {
	events []keyboard.Event
}

func (f *fakeEventSource) Poll() (keyboard.Event, bool) {
	if len(f.events) > 0 {
		event := f.events[0]
		f.events = f.events[1:]
		return event, true
	}
	return keyboard.Event{}, false // TODO Maybe we need an EmptyEvent?
}
