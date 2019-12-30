package pixiq

// NewWindows returns a factory of Window objects.
func NewWindows(images *Images, windows SystemWindows) *Windows {
	return &Windows{images: images, systemWindows: windows}
}

// SystemWindows is an API for creating system windows
type SystemWindows interface {
	// Open creates and show system window
	Open(width, height int) SystemWindow
}

// SystemWindow is an operating system window
type SystemWindow interface {
	Draw(image *Image)
}

// Windows is a factory of Window objects
type Windows struct {
	images        *Images
	systemWindows SystemWindows
}

// New creates a new window with width and height given in pixels.
func (w *Windows) New(width, height int) *Window {
	if width < 0 {
		width = 0
	}
	if height < 0 {
		height = 0
	}
	return &Window{
		width:  width,
		height: height,
		image:  w.images.New(width, height),
		window: w.systemWindows.Open(width, height),
	}
}

// Window is area where image will be drawn
type Window struct {
	width  int
	height int
	image  *Image
	window SystemWindow
}

// Width returns width of the window in pixels
func (w *Window) Width() int {
	return w.width
}

// Height returns height of the window in pixels
func (w *Window) Height() int {
	return w.height
}

// Loop starts the main loop. It will execute onEachFrame function for each frame, as soon as window is closed. This
// function blocks the current goroutine.
func (w *Window) Loop(onEachFrame func(frame *Frame)) {
	frame := &Frame{}
	frame.image = w.image
	for !frame.closeWindow {
		onEachFrame(frame)
		w.window.Draw(w.image)
	}
}

// Frame provides the whole window image which will be drawn on a screen after making modifications
type Frame struct {
	closeWindow bool
	image       *Image
}

// CloseWindowEventually closes the window as soon as onEachFrame function is finished
func (w *Frame) CloseWindowEventually() {
	w.closeWindow = true
}

// Image returns the whole window Image, which can be modified
func (w *Frame) Image() *Image {
	return w.image
}