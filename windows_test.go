package pixiq_test

import (
	"github.com/jacekolszak/pixiq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewWindows(t *testing.T) {
	t.Run("should return Windows object for creating windows", func(t *testing.T) {
		windows := pixiq.NewWindows()
		assert.NotNil(t, windows)
	})
}

func TestWindow_New(t *testing.T) {
	windows := pixiq.NewWindows()
	t.Run("should clamp width to 0 if negative", func(t *testing.T) {
		win := windows.New(-1, 0)
		require.NotNil(t, win)
		assert.Equal(t, 0, win.Width())
	})
	t.Run("should clamp height to 0 if negative", func(t *testing.T) {
		win := windows.New(0, -1)
		require.NotNil(t, win)
		assert.Equal(t, 0, win.Height())
	})
	t.Run("should create window", func(t *testing.T) {
		win := windows.New(1, 2)
		require.NotNil(t, win)
		assert.Equal(t, 1, win.Width())
		assert.Equal(t, 2, win.Height())
	})
}

func TestWindow_Loop(t *testing.T) {
	windows := pixiq.NewWindows()
	window := windows.New(0, 0)

	t.Run("should run callback function until window is closed", func(t *testing.T) {
		executionCount := 0
		// when
		window.Loop(func(frame *pixiq.Frame) {
			executionCount += 1
			if executionCount == 2 {
				frame.CloseWindowEventually()
			}
		})
		// then
		assert.Equal(t, 2, executionCount)
	})

	t.Run("frame should provide Selection for whole window's image", func(t *testing.T) {
		windows := pixiq.NewWindows()
		tests := map[string]struct {
			width, height int
		}{
			"0x0": {
				width:  0,
				height: 0,
			},
			"1x2": {
				width:  1,
				height: 2,
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				window := windows.New(test.width, test.height)
				var selection pixiq.Selection
				// when
				window.Loop(func(frame *pixiq.Frame) {
					selection = frame.ImageSelection()
					frame.CloseWindowEventually()
				})
				// then
				assert.Equal(t, 0, selection.ImageX())
				assert.Equal(t, 0, selection.ImageY())
				assert.Equal(t, test.width, selection.Width())
				assert.Equal(t, test.height, selection.Height())
				assert.NotNil(t, selection.Image())
			})

		}
	})

}
