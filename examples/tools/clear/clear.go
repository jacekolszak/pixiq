package main

import (
	"github.com/jacekolszak/pixiq/colornames"
	"github.com/jacekolszak/pixiq/image"
	"github.com/jacekolszak/pixiq/keyboard"
	"github.com/jacekolszak/pixiq/loop"
	"github.com/jacekolszak/pixiq/opengl"
	"github.com/jacekolszak/pixiq/tools/clear"
	"github.com/jacekolszak/pixiq/tools/glclear"
)

type ClearTool interface {
	SetColor(image.Color)
	Clear(image.Selection)
}

func main() {
	opengl.RunOrDie(func(gl *opengl.OpenGL) {
		window, err := gl.OpenWindow(10, 10, opengl.Zoom(30))
		if err != nil {
			panic(err)
		}
		context := gl.Context()
		tools := []ClearTool{
			glclear.New(context.NewClearCommand()), // GPU one
			clear.New(),                            // CPU one
		}
		tools[0].SetColor(colornames.Cornflowerblue)
		tools[1].SetColor(colornames.Hotpink)
		currentTool := 0
		keys := keyboard.New(window)
		loop.Run(window, func(frame *loop.Frame) {
			screen := frame.Screen()
			var (
				leftEye  = screen.Selection(2, 2).WithSize(2, 2)
				rightEye = screen.Selection(6, 2).WithSize(2, 2)
				nose     = screen.Selection(4, 5).WithSize(2, 2)
				mouth    = screen.Selection(2, 8).WithSize(6, 1)
			)
			tool := tools[currentTool]
			tool.Clear(leftEye)
			tool.Clear(rightEye)
			tool.Clear(nose)
			tool.Clear(mouth)
			keys.Update()
			if keys.JustReleased(keyboard.Space) {
				currentTool += 1
				currentTool = currentTool % len(tools)
			}
		})

	})
}
