package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/jacekolszak/pixiq/gl"
	"github.com/jacekolszak/pixiq/glfw/internal"
	"github.com/jacekolszak/pixiq/image"
	"github.com/jacekolszak/pixiq/keyboard"
)

// Window is an implementation of loop.Screen and keyboard.EventSource
type Window struct {
	glfwWindow             *glfw.Window
	mainThreadLoop         *MainThreadLoop
	screenPolygon          *screenPolygon
	keyboardEvents         *internal.KeyboardEvents
	requestedWidth         int
	requestedHeight        int
	zoom                   int
	screenImage            *image.Image
	screenAcceleratedImage *gl.AcceleratedImage
	api                    gl.API
	context                *gl.Context
	program                *gl.Program
}

// Draw draws a screen image to the invisible buffer. It will be shown in window
// after SwapImages is called.
func (w *Window) Draw() {
	w.screenImage.Upload()
	var width, height int
	w.mainThreadLoop.Execute(func() {
		width, height = w.glfwWindow.GetFramebufferSize()
	})
	w.api.Viewport(0, 0, int32(width), int32(height))
	w.screenPolygon.draw()
}

// SwapImages makes last drawn image visible in window.
func (w *Window) SwapImages() {
	w.mainThreadLoop.Execute(w.glfwWindow.SwapBuffers)
}

// Close closes the window and cleans resources.
func (w *Window) Close() {
	w.mainThreadLoop.Execute(w.glfwWindow.Destroy)
}

// ShouldClose reports the value of the close flag of the window.
// The flag is set to true when user clicks Close button or hits ALT+F4/CMD+Q.
func (w *Window) ShouldClose() bool {
	var shouldClose bool
	w.mainThreadLoop.Execute(func() {
		shouldClose = w.glfwWindow.ShouldClose()
	})
	return shouldClose
}

// Width returns the actual width of the window in pixels. It may be different
// than requested width used when window was open due to platform limitation.
// If zooming is used the width is multiplied by zoom.
func (w *Window) Width() int {
	var width int
	w.mainThreadLoop.Execute(func() {
		width, _ = w.glfwWindow.GetSize()
	})
	return width
}

// Height returns the actual height of the window in pixels. It may be different
// than requested height used when window was open due to platform limitation.
// If zooming is used the height is multiplied by zoom.
func (w *Window) Height() int {
	var height int
	w.mainThreadLoop.Execute(func() {
		_, height = w.glfwWindow.GetSize()
	})
	return height
}

// Zoom returns the actual zoom. It is the zoom given during opening the window,
// unless zoom < 1 was given - then the actual zoom is 1.
func (w *Window) Zoom() int {
	return w.zoom
}

// Poll retrieves and removes next keyboard Event. If there are no more
// events false is returned. It implements keyboard.EventSource method.
func (w *Window) Poll() (keyboard.Event, bool) {
	var (
		event keyboard.Event
		ok    bool
	)
	w.mainThreadLoop.Execute(func() {
		event, ok = w.keyboardEvents.Poll()
	})
	return event, ok
}

// Image returns the screen's image
func (w *Window) Image() *image.Image {
	return w.screenImage
}

// ContextAPI returns window-specific OpenGL's context. Useful for accessing
// window's framebuffer.
func (w *Window) ContextAPI() gl.API {
	return w.api
}