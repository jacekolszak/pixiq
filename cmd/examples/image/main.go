package main

import (
	"fmt"

	"github.com/jacekolszak/pixiq"
	"github.com/jacekolszak/pixiq/opengl"
)

var white = pixiq.RGBA(255, 255, 255, 255)

// This program shows how to create an Image and manipulate its pixels.
// It does not open any windows.
func main() {
	// OpenGL object is needed for implementation of AcceleratedImages.
	// Unfortunately some functions used by opengl package must be executed
	// from the main thread. This program is executed in the main thread,
	// so we can pass it to the opengl by executing StartMainThreadLoop.
	// The function will block and will execute our callback in a different
	// thread.
	opengl.StartMainThreadLoop(func(loop *opengl.MainThreadLoop) {
		// Create OpenGL object.
		gl := opengl.New(loop)

		// Create pixiq.Images which is basically a factory of images.
		// Inject OpenGL's implementation of pixiq.AcceleratedImages.
		images := pixiq.NewImages(gl.AcceleratedImages())

		// Create 2x2 image. Width parameter (as always) is before Height.
		// All given in pixels.
		image := images.New(2, 2)

		// Image can be modified via Selection. Here we create a selection
		// spanning the whole Image. The Selection starts at (0,0) and have
		// a size of the Image - width=2 and height=2
		wholeSelection := image.WholeImageSelection()

		// Set the pixel color to white at x=0, y=0. X is always before Y.
		wholeSelection.SetColor(0, 1, white)
		// Get the pixel color.
		color := wholeSelection.Color(0, 1)
		// The pixel color at (0,1) is {255 255 255 255}.
		fmt.Printf("The pixel color at (0,1) is %v.\n", color)

		// Create Selection starting at x=1, y=1
		selection := image.Selection(1, 1).WithSize(1, 1)
		// Use Selection local coordinates (0,0) to access Image at (1,1).
		selection.SetColor(0, 0, white)
		color = wholeSelection.Color(1, 1)
		// The pixel color at (1,1) is {255 255 255 255}.
		fmt.Printf("The pixel color at (1,1) is %v.\n", color)

		// Access Image outside the Selection.
		color = selection.Color(-1, 0)
		// The pixel color at (0,1) is {255 255 255 255}.
		fmt.Printf("The pixel color at (0,1) is %v.\n", color)

		// Access pixels outside the Image always return transparent color.
		color = wholeSelection.Color(-1, 0)
		// The pixel color at (-1,0) is {0 0 0 0}.
		fmt.Printf("The pixel color at (-1,0) is %v.\n", color)
	})
}