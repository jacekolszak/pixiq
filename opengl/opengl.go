// Package opengl makes it possible to use Pixiq on PCs with Linux, Windows or MacOS.
// It provides a method for creating OpenGL-accelerated image.Image and Window which
// is an implementation of loop.Screen and keyboard.EventSource.
// Under the hood it is using OpenGL API and GLFW for manipulating windows
// and handling user input.
package opengl

import (
	"log"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/jacekolszak/pixiq/image"
	"github.com/jacekolszak/pixiq/keyboard"
	"github.com/jacekolszak/pixiq/opengl/internal"
)

// New creates OpenGL instance.
// MainThreadLoop is needed because some GLFW functions has to be called
// from the main thread.
// There is a possibility to create multiple OpenGL objects. Please note though
// that some platforms may limit this number. In integration tests you should
// always remember to destroy the object after test by executing Destroy method,
// because eventually the number of objects may reach the mentioned limit.
func New(mainThreadLoop *MainThreadLoop) *OpenGL {
	if mainThreadLoop == nil {
		panic("nil MainThreadLoop")
	}
	var (
		mainWindow *glfw.Window
		err        error
	)
	mainThreadLoop.Execute(func() {
		err = glfw.Init()
		if err != nil {
			return
		}
		mainWindow, err = createWindow(mainThreadLoop, nil)
		if err != nil {
			return
		}
	})
	if err != nil {
		panic(err)
	}
	openGL := &OpenGL{
		mainThreadLoop:     mainThreadLoop,
		bindWindowToThread: mainThreadLoop.bind(mainWindow),
		stopPollingEvents:  make(chan struct{}),
		windows: &Windows{
			mainWindow:     mainWindow,
			mainThreadLoop: mainThreadLoop,
		},
	}
	openGL.windows.newTexture = openGL.newTexture
	go openGL.startPollingEvents(openGL.stopPollingEvents)
	return openGL
}

// Run is a shorthand method for creating Pixiq objects with OpenGL acceleration
// and Windows. It runs the given callback function and blocks. It was created
// mainly for educational purposes to save a few keystrokes.
func Run(main func(gl *OpenGL)) {
	StartMainThreadLoop(func(mainThreadLoop *MainThreadLoop) {
		openGL := New(mainThreadLoop)
		defer openGL.Destroy()
		main(openGL)
	})
}

func createWindow(mainThreadLoop *MainThreadLoop, share *glfw.Window) (*glfw.Window, error) {
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Visible, glfw.False)
	glfw.WindowHint(glfw.CocoaRetinaFramebuffer, glfw.False)
	// FIXME: For some reason XVFB does not change the frame buffer size after
	// resizing the window to higher values. That's why the window created
	// here has size equal to the biggest window used in integration tests
	// See: TestWindow_Draw() in opengl_test.go
	win, err := glfw.CreateWindow(3, 3, "OpenGL Pixiq Window", nil, share)
	if err != nil {
		return nil, err
	}
	mainThreadLoop.bind(win)()
	if err := gl.Init(); err != nil {
		return nil, err
	}
	return win, nil
}

// OpenGL provides method for creating OpenGL-accelerated image.Image and provides
// Windows object for opening windows.
type OpenGL struct {
	mainThreadLoop     *MainThreadLoop
	bindWindowToThread func()
	windows            *Windows
	stopPollingEvents  chan struct{}
}

// Windows returns object for opening system windows. Each open Window
// is a pixiq.Screen implementation.
func (g *OpenGL) Windows() *Windows {
	return g.windows
}

// Destroy cleans all the OpenGL resources associated with this instance.
// This method has to be called in integration tests to clean resources after
// each test. Otherwise on some platforms you may reach the limit of active
// OpenGL contexts.
func (g *OpenGL) Destroy() {
	g.stopPollingEvents <- struct{}{}
	mainThreadLoop := g.windows.mainThreadLoop
	mainThreadLoop.Execute(func() {
		mainThreadLoop.bind(g.windows.mainWindow)()
		g.windows.mainWindow.Destroy()
	})
}

func (g *OpenGL) startPollingEvents(stop <-chan struct{}) {
	// fixme: make it configurable
	ticker := time.NewTicker(4 * time.Millisecond) // 250Hz
	for {
		<-ticker.C
		select {
		case <-stop:
			return
		default:
			g.windows.mainThreadLoop.Execute(glfw.PollEvents)
		}
	}
}

// NewImage creates an *image.Image which is using OpenGL acceleration
// under-the-hood.
//
// Example:
//
//	   gl := opengl.New(loop)
//	   defer gl.Destroy()
//	   img := gl.NewImage(2, 2)
//
// To avoid coupling with opengl you should define your own factory function
// for creating images and use it instead of directly accessing opengl.OpenGL:
//
//	   type NewImage func(width, height) *image.Image
//
func (g *OpenGL) NewImage(width, height int) *image.Image {
	return image.New(width, height, g.NewAcceleratedImage(width, height))
}

// Windows is used for opening system windows.
type Windows struct {
	mainThreadLoop *MainThreadLoop
	// mainWindow contains textures and user programs
	mainWindow *glfw.Window
	newTexture func(width int, height int) *texture
}

// Open creates and shows Window.
func (w *Windows) Open(width, height int, options ...WindowOption) *Window {
	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}
	// FIXME: EventBuffer size should be configurable
	keyboardEvents := internal.NewKeyboardEvents(keyboard.NewEventBuffer(32))
	screenTexture := w.newTexture(width, height)
	screenImage := image.New(width, height, screenTexture)
	win := &Window{
		mainThreadLoop:  w.mainThreadLoop,
		keyboardEvents:  keyboardEvents,
		requestedWidth:  width,
		requestedHeight: height,
		screenTexture:   screenTexture,
		screenImage:     screenImage,
		zoom:            1,
	}
	var err error
	w.mainThreadLoop.Execute(func() {
		win.glfwWindow, err = createWindow(w.mainThreadLoop, w.mainWindow)
		if err != nil {
			return
		}
		win.glfwWindow.SetKeyCallback(win.keyboardEvents.OnKeyCallback)
		win.program, err = compileProgram()
		if err != nil {
			return
		}
		win.screenPolygon = newScreenPolygon(
			win.program.vertexPositionLocation,
			win.program.texturePositionLocation)
		for _, option := range options {
			if option == nil {
				log.Println("nil option given when opening the window")
				continue
			}
			option(win)
		}
		win.glfwWindow.SetSize(win.requestedWidth*win.zoom, win.requestedHeight*win.zoom)
		win.glfwWindow.Show()
	})
	if err != nil {
		panic(err)
	}
	return win
}

// WindowOption is an option used when opening the window.
type WindowOption func(window *Window)

// NoDecorationHint is Window hint hiding the border, close widget, etc.
// Exact behaviour depends on the platform.
func NoDecorationHint() WindowOption {
	return func(win *Window) {
		win.glfwWindow.SetAttrib(glfw.Decorated, glfw.False)
	}
}

// Title sets the window title.
func Title(title string) WindowOption {
	return func(window *Window) {
		window.glfwWindow.SetTitle(title)
	}
}

// Zoom makes window/pixels bigger zoom times.
func Zoom(zoom int) WindowOption {
	return func(window *Window) {
		if zoom > 0 {
			window.zoom = zoom
		}
	}
}

// Window is an implementation of loop.Screen and keyboard.EventSource
type Window struct {
	glfwWindow      *glfw.Window
	program         *program
	mainThreadLoop  *MainThreadLoop
	screenPolygon   *screenPolygon
	keyboardEvents  *internal.KeyboardEvents
	requestedWidth  int
	requestedHeight int
	zoom            int
	screenImage     *image.Image
	screenTexture   *texture
}

// Draw draws a screen image to the invisible buffer. It will be shown in window
// after SwapImages is called.
func (w *Window) Draw() {
	w.screenImage.Upload()
	w.mainThreadLoop.Execute(func() {
		w.mainThreadLoop.bind(w.glfwWindow)()
		w.program.use()
		width, height := w.glfwWindow.GetFramebufferSize()
		gl.Viewport(0, 0, int32(width), int32(height))
		gl.BindTexture(gl.TEXTURE_2D, w.screenTexture.TextureID())
		w.screenPolygon.draw()
	})
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

// NewAcceleratedImage returns an OpenGL-accelerated implementation of image.AcceleratedImage
func (g *OpenGL) NewAcceleratedImage(width, height int) image.AcceleratedImage {
	return g.newTexture(width, height)
}

func (g *OpenGL) newTexture(width, height int) *texture {
	var id uint32
	g.mainThreadLoop.Execute(func() {
		g.bindWindowToThread()
		gl.GenTextures(1, &id)
		gl.BindTexture(gl.TEXTURE_2D, id)
		gl.TexImage2D(
			gl.TEXTURE_2D,
			0,
			gl.RGBA,
			int32(width),
			int32(height),
			0,
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			gl.Ptr(nil),
		)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	})
	return &texture{
		id:                 id,
		width:              width,
		height:             height,
		mainThreadLoop:     g.mainThreadLoop,
		bindWindowToThread: g.bindWindowToThread,
	}
}

type texture struct {
	id                 uint32
	width, height      int
	mainThreadLoop     *MainThreadLoop
	bindWindowToThread func()
}

func (t *texture) TextureID() uint32 {
	return t.id
}

func (t *texture) Upload(pixels []image.Color) {
	t.mainThreadLoop.Execute(func() {
		t.bindWindowToThread()
		gl.BindTexture(gl.TEXTURE_2D, t.id)
		gl.TexSubImage2D(
			gl.TEXTURE_2D,
			0,
			int32(0),
			int32(0),
			int32(t.width),
			int32(t.height),
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			gl.Ptr(pixels),
		)
	})
}
func (t *texture) Download(output []image.Color) {
	t.mainThreadLoop.Execute(func() {
		t.bindWindowToThread()
		gl.BindTexture(gl.TEXTURE_2D, t.id)
		gl.GetTexImage(
			gl.TEXTURE_2D,
			0,
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			gl.Ptr(output),
		)
	})
}
