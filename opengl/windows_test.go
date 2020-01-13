package opengl_test

import (
	"fmt"
	"testing"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jacekolszak/pixiq/image"
	"github.com/jacekolszak/pixiq/keyboard"
	"github.com/jacekolszak/pixiq/opengl"
)

func TestOpenGL_Open(t *testing.T) {
	t.Run("should constrain width to platform-specific minimum if negative", func(t *testing.T) {
		openGL := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		// when
		win := openGL.Open(-1, 0)
		defer win.Close()
		// then
		require.NotNil(t, win)
		assert.GreaterOrEqual(t, win.Width(), 0)
	})
	t.Run("should constrain height to platform-specific minimum if negative", func(t *testing.T) {
		openGL := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		// when
		win := openGL.Open(0, -1)
		defer win.Close()
		// then
		require.NotNil(t, win)
		assert.GreaterOrEqual(t, win.Height(), 0)
	})
	t.Run("should open Window", func(t *testing.T) {
		openGL := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		// when
		win := openGL.Open(640, 360)
		defer win.Close()
		// then
		require.NotNil(t, win)
		assert.Equal(t, 640, win.Width())
		assert.Equal(t, 360, win.Height())
	})
	t.Run("should open two windows at the same time", func(t *testing.T) {
		openGL := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		// when
		win1 := openGL.Open(640, 360)
		defer win1.Close()
		win2 := openGL.Open(320, 180)
		defer win2.Close()
		// then
		require.NotNil(t, win1)
		assert.Equal(t, 640, win1.Width())
		assert.Equal(t, 360, win1.Height())
		require.NotNil(t, win2)
		assert.Equal(t, 320, win2.Width())
		assert.Equal(t, 180, win2.Height())
	})
	t.Run("should open another Window after first one was closed", func(t *testing.T) {
		openGL := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		win1 := openGL.Open(640, 360)
		win1.Close()
		// when
		win2 := openGL.Open(320, 180)
		defer win2.Close()
		// then
		require.NotNil(t, win2)
		assert.Equal(t, 320, win2.Width())
		assert.Equal(t, 180, win2.Height())
	})
	t.Run("should skip nil option", func(t *testing.T) {
		openGL := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		// when
		win := openGL.Open(0, 0, nil)
		defer win.Close()
	})
	t.Run("zoom <= 1 should not affect the width and height", func(t *testing.T) {
		tests := map[string]struct {
			zoom int
		}{
			"zoom = -1": {
				zoom: -1,
			},
			"zoom = 0": {
				zoom: 0,
			},
			"zoom = 1": {
				zoom: 1,
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				openGL := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				// when
				win := openGL.Open(640, 360, opengl.Zoom(test.zoom))
				defer win.Close()
				// then
				require.NotNil(t, win)
				assert.Equal(t, 640, win.Width())
				assert.Equal(t, 360, win.Height())
			})
		}
	})
	t.Run("zoom should affect the width and height", func(t *testing.T) {
		tests := map[string]struct {
			zoom           int
			expectedWidth  int
			expectedHeight int
		}{
			"zoom = 2": {
				zoom:           2,
				expectedWidth:  1280,
				expectedHeight: 720,
			},
			"zoom = 3": {
				zoom:           3,
				expectedWidth:  1920,
				expectedHeight: 1080,
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				openGL := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				// when
				win := openGL.Open(640, 360, opengl.Zoom(test.zoom))
				defer win.Close()
				// then
				require.NotNil(t, win)
				assert.Equal(t, test.expectedWidth, win.Width())
				assert.Equal(t, test.expectedHeight, win.Height())
			})
		}
	})
}

func TestWindow_Draw(t *testing.T) {
	t.Run("should draw screen image", func(t *testing.T) {
		color1 := image.RGBA(10, 20, 30, 40)
		color2 := image.RGBA(50, 60, 70, 80)
		color3 := image.RGBA(90, 100, 110, 120)
		color4 := image.RGBA(130, 140, 150, 160)

		t.Run("1x1", func(t *testing.T) {
			openGL := opengl.New(mainThreadLoop)
			defer openGL.Destroy()
			window := openGL.Open(1, 1, opengl.NoDecorationHint())
			defer window.Close()
			window.Image().WholeImageSelection().SetColor(0, 0, color1)
			// when
			window.Draw()
			// then
			expected := []image.Color{color1}
			assert.Equal(t, expected, framebufferPixels(0, 0, 1, 1))
		})
		t.Run("1x2", func(t *testing.T) {
			openGL := opengl.New(mainThreadLoop)
			defer openGL.Destroy()
			window := openGL.Open(1, 2, opengl.NoDecorationHint())
			defer window.Close()
			img := window.Image()
			img.WholeImageSelection().SetColor(0, 0, color1)
			img.WholeImageSelection().SetColor(0, 1, color2)
			// when
			window.Draw()
			// then
			expected := []image.Color{color2, color1}
			assert.Equal(t, expected, framebufferPixels(0, 0, 1, 2))
		})
		t.Run("2x1", func(t *testing.T) {
			openGL := opengl.New(mainThreadLoop)
			defer openGL.Destroy()
			window := openGL.Open(2, 1, opengl.NoDecorationHint())
			defer window.Close()
			img := window.Image()
			img.WholeImageSelection().SetColor(0, 0, color1)
			img.WholeImageSelection().SetColor(1, 0, color2)
			// when
			window.Draw()
			// then
			expected := []image.Color{color1, color2}
			assert.Equal(t, expected, framebufferPixels(0, 0, 2, 1))
		})
		t.Run("2x2", func(t *testing.T) {
			openGL := opengl.New(mainThreadLoop)
			defer openGL.Destroy()
			window := openGL.Open(2, 2, opengl.NoDecorationHint())
			defer window.Close()
			img := window.Image()
			selection := img.WholeImageSelection()
			selection.SetColor(0, 0, color1)
			selection.SetColor(1, 0, color2)
			selection.SetColor(0, 1, color3)
			selection.SetColor(1, 1, color4)
			// when
			window.Draw()
			// then
			expected := []image.Color{color3, color4, color1, color2}
			assert.Equal(t, expected, framebufferPixels(0, 0, 2, 2))
		})

		t.Run("zoom < 1 should not change the framebuffer size", func(t *testing.T) {
			for zoom := -1; zoom < 1; zoom++ {
				name := fmt.Sprintf("zoom=%d", zoom)
				t.Run(name, func(t *testing.T) {
					openGL := opengl.New(mainThreadLoop)
					defer openGL.Destroy()
					window := openGL.Open(1, 1, opengl.NoDecorationHint(), opengl.Zoom(zoom))
					defer window.Close()
					img := window.Image()
					img.WholeImageSelection().SetColor(0, 0, color1)
					// when
					window.Draw()
					// then
					expected := []image.Color{color1}
					assert.Equal(t, expected, framebufferPixels(0, 0, 1, 1))
				})
			}
		})

		t.Run("zoom > 1 should make framebuffer zoom times bigger", func(t *testing.T) {
			for zoom := 2; zoom < 4; zoom++ {
				name := fmt.Sprintf("zoom=%d", zoom)
				t.Run(name, func(t *testing.T) {
					openGL := opengl.New(mainThreadLoop)
					defer openGL.Destroy()
					window := openGL.Open(1, 1, opengl.NoDecorationHint(), opengl.Zoom(zoom))
					defer window.Close()
					img := window.Image()
					img.WholeImageSelection().SetColor(0, 0, color1)
					// when
					window.Draw()
					// then
					expected := make([]image.Color, zoom*zoom)
					for i := 0; i < len(expected); i++ {
						expected[i] = color1
					}
					assert.Equal(t, expected, framebufferPixels(0, 0, int32(zoom), int32(zoom)))
				})
			}
		})

		t.Run("two windows", func(t *testing.T) {
			openGL := opengl.New(mainThreadLoop)
			defer openGL.Destroy()
			window1 := windowOfColor(openGL, color1)
			defer window1.Close()
			window2 := windowOfColor(openGL, color2)
			defer window2.Close()
			// when
			window1.Draw()
			// then
			expected := []image.Color{color1}
			assert.Equal(t, expected, framebufferPixels(0, 0, 1, 1))
			// when
			window2.Draw()
			// then
			expected = []image.Color{color2}
			assert.Equal(t, expected, framebufferPixels(0, 0, 1, 1))
		})

		t.Run("two OpenGL instances", func(t *testing.T) {
			openGL1 := opengl.New(mainThreadLoop)
			defer openGL1.Destroy()
			openGL2 := opengl.New(mainThreadLoop)
			defer openGL2.Destroy()
			window1 := windowOfColor(openGL1, color1)
			defer window1.Close()
			window2 := windowOfColor(openGL2, color2)
			defer window2.Close()
			// when
			window1.Draw()
			// then
			expected := []image.Color{color1}
			assert.Equal(t, expected, framebufferPixels(0, 0, 1, 1))
			// when
			window2.Draw()
			// then
			expected = []image.Color{color2}
			assert.Equal(t, expected, framebufferPixels(0, 0, 1, 1))
		})
	})
}

func windowOfColor(openGL *opengl.OpenGL, color image.Color) *opengl.Window {
	var (
		window    = openGL.Open(1, 1, opengl.NoDecorationHint())
		selection = window.Image().WholeImageSelection()
	)
	selection.SetColor(0, 0, color)
	return window
}

func framebufferPixels(x, y, width, height int32) []image.Color {
	size := (height - y) * (width - x)
	frameBuffer := make([]image.Color, size)
	mainThreadLoop.Execute(func() {
		gl.ReadPixels(x, y, width, height, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(frameBuffer))
	})
	return frameBuffer
}

func TestWindow_Poll(t *testing.T) {
	t.Run("should return EmptyEvent and false when there is no keyboard events", func(t *testing.T) {
		openGL := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		win := openGL.Open(1, 1)
		defer win.Close()
		// when
		event, ok := win.Poll()
		// then
		assert.Equal(t, keyboard.EmptyEvent, event)
		assert.False(t, ok)
	})
}

func TestWindow_Zoom(t *testing.T) {
	t.Run("should return specified zoom for window", func(t *testing.T) {
		tests := map[string]struct {
			zoom         int
			expectedZoom int
		}{
			"zoom -1": {
				zoom:         -1,
				expectedZoom: 1,
			},
			"zoom 0": {
				zoom:         0,
				expectedZoom: 1,
			},
			"zoom 1": {
				zoom:         1,
				expectedZoom: 1,
			},
			"zoom 2": {
				zoom:         2,
				expectedZoom: 2,
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				openGL := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				win := openGL.Open(0, 0, opengl.Zoom(test.zoom))
				defer win.Close()
				// when
				zoom := win.Zoom()
				// expect
				assert.Equal(t, test.expectedZoom, zoom)
			})
		}
	})
}

func TestWindow_Image(t *testing.T) {
	t.Run("should provide screen image", func(t *testing.T) {
		tests := map[string]struct {
			width, height int
		}{
			"1x2": {
				width:  1,
				height: 2,
			},
			"3x4": {
				width:  3,
				height: 4,
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				openGL := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				win := openGL.Open(test.width, test.height, opengl.NoDecorationHint())
				defer win.Close()
				// when
				img := win.Image()
				// then
				require.NotNil(t, img)
				assert.Equal(t, test.width, img.Width())
				assert.Equal(t, test.height, img.Height())
			})
		}
	})
	t.Run("zoom should not affect the screen size", func(t *testing.T) {
		tests := map[string]struct {
			zoom int
		}{
			"zoom = -1": {
				zoom: 1,
			},
			"zoom = 0": {
				zoom: 0,
			},
			"zoom = 1": {
				zoom: 1,
			},
			"zoom = 2": {
				zoom: 2,
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				openGL := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				win := openGL.Open(640, 360, opengl.Zoom(test.zoom))
				// when
				screen := win.Image()
				// then
				assert.Equal(t, 640, screen.Width())
				assert.Equal(t, 360, screen.Height())
			})
		}
	})
	t.Run("initial screen is transparent", func(t *testing.T) {
		openGL := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		win := openGL.Open(1, 1, opengl.NoDecorationHint())
		transparent := image.RGBA(0, 0, 0, 0)
		// when
		img := win.Image()
		// then
		selection := img.WholeImageSelection()
		assert.Equal(t, transparent, selection.Color(0, 0))
	})
}