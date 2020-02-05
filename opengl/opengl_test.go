package opengl_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jacekolszak/pixiq/image"
	"github.com/jacekolszak/pixiq/opengl"
)

var mainThreadLoop *opengl.MainThreadLoop

func TestMain(m *testing.M) {
	var exit int
	opengl.StartMainThreadLoop(func(main *opengl.MainThreadLoop) {
		mainThreadLoop = main
		exit = m.Run()
	})
	os.Exit(exit)
}

func TestNew(t *testing.T) {
	t.Run("should panic when MainThreadLoop is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			_, _ = opengl.New(nil)
		})
	})
	t.Run("should create OpenGL using supplied MainThreadLoop", func(t *testing.T) {
		// when
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// then
		assert.NotNil(t, openGL)
	})
	t.Run("should create 2 objects working at the same time", func(t *testing.T) {
		for i := 0; i < 2; i++ {
			openGL, err := opengl.New(mainThreadLoop)
			require.NoError(t, err)
			defer openGL.Destroy()
		}
	})
}

func TestOpenGL_NewImage(t *testing.T) {
	t.Run("should return error for negative width", func(t *testing.T) {
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// when
		img, err := openGL.NewImage(-1, 0)
		// then
		require.Error(t, err)
		assert.Nil(t, img)
	})
	t.Run("should return error for negative height", func(t *testing.T) {
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// when
		img, err := openGL.NewImage(0, -1)
		// then
		require.Error(t, err)
		assert.Nil(t, img)
	})
	t.Run("should create Image", func(t *testing.T) {
		tests := map[string]struct {
			width  int
			height int
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
				openGL, err := opengl.New(mainThreadLoop)
				require.NoError(t, err)
				defer openGL.Destroy()
				// when
				img, err := openGL.NewImage(test.width, test.height)
				// then
				require.NoError(t, err)
				assert.NotNil(t, img)
				assert.Equal(t, test.width, img.Width())
				assert.Equal(t, test.height, img.Height())
			})
		}
	})
}

func TestOpenGL_NewTexture(t *testing.T) {
	t.Run("should return error for negative width", func(t *testing.T) {
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// when
		img, err := openGL.NewTexture(-1, 0)
		// then
		require.Error(t, err)
		assert.Nil(t, img)
	})
	t.Run("should return error for negative height", func(t *testing.T) {
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// when
		img, err := openGL.NewTexture(0, -1)
		// then
		require.Error(t, err)
		assert.Nil(t, img)
	})
	t.Run("should create AcceleratedImage", func(t *testing.T) {
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// when
		img, err := openGL.NewTexture(0, 0)
		// then
		require.NoError(t, err)
		assert.NotNil(t, img)
	})
}

func TestTexture_Upload(t *testing.T) {
	color1 := image.RGBA(10, 20, 30, 40)
	color2 := image.RGBA(50, 60, 70, 80)
	color3 := image.RGBA(90, 100, 110, 120)
	color4 := image.RGBA(130, 140, 150, 160)

	t.Run("should upload pixels", func(t *testing.T) {
		tests := map[string]struct {
			width, height int
			inputColors   []image.Color
		}{
			"0x0": {
				width:       0,
				height:      0,
				inputColors: []image.Color{},
			},
			"1x1": {
				width:       1,
				height:      1,
				inputColors: []image.Color{color1},
			},
			"2x1": {
				width:       2,
				height:      1,
				inputColors: []image.Color{color1, color2},
			},
			"1x2": {
				width:       1,
				height:      2,
				inputColors: []image.Color{color1, color2},
			},
			"2x2": {
				width:       2,
				height:      2,
				inputColors: []image.Color{color1, color2, color3, color4},
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				openGL, err := opengl.New(mainThreadLoop)
				require.NoError(t, err)
				defer openGL.Destroy()
				img, err := openGL.NewTexture(test.width, test.height)
				require.NoError(t, err)
				// when
				img.Upload(test.inputColors)
				// then
				output := make([]image.Color, len(test.inputColors))
				img.Download(output)
				assert.Equal(t, test.inputColors, output)
			})
		}
	})
	t.Run("2 OpenGL contexts", func(t *testing.T) {
		gl1, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer gl1.Destroy()
		gl2, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer gl2.Destroy()
		img1, err := gl1.NewTexture(1, 1)
		require.NoError(t, err)
		img2, err := gl2.NewTexture(1, 1)
		require.NoError(t, err)
		// when
		img1.Upload([]image.Color{color1})
		img2.Upload([]image.Color{color2})
		// then
		output := make([]image.Color, 1)
		img1.Download(output)
		assert.Equal(t, []image.Color{color1}, output)
		// and
		img2.Download(output)
		assert.Equal(t, []image.Color{color2}, output)
	})
}

func TestRunOrDie(t *testing.T) {
	t.Run("should run provided callback", func(t *testing.T) {
		var callbackExecuted bool
		mainThreadLoop.Execute(func() {
			opengl.RunOrDie(func(gl *opengl.OpenGL) {
				callbackExecuted = true
			})
		})
		assert.True(t, callbackExecuted)
	})
	t.Run("should start a MainThreadLoop and create OpenGL object", func(t *testing.T) {
		var (
			actualGL *opengl.OpenGL
		)
		mainThreadLoop.Execute(func() {
			opengl.RunOrDie(func(gl *opengl.OpenGL) {
				actualGL = gl
			})
		})
		assert.NotNil(t, actualGL)
	})
}

func TestOpenGL_OpenWindow(t *testing.T) {
	t.Run("should constrain width to platform-specific minimum if negative", func(t *testing.T) {
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// when
		win, err := openGL.OpenWindow(-1, 0)
		require.NoError(t, err)
		defer win.Close()
		// then
		require.NotNil(t, win)
		assert.GreaterOrEqual(t, win.Width(), 0)
	})
	t.Run("should constrain height to platform-specific minimum if negative", func(t *testing.T) {
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// when
		win, err := openGL.OpenWindow(0, -1)
		require.NoError(t, err)
		defer win.Close()
		// then
		require.NotNil(t, win)
		assert.GreaterOrEqual(t, win.Height(), 0)
	})
	t.Run("should open Window", func(t *testing.T) {
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// when
		win, err := openGL.OpenWindow(640, 360)
		require.NoError(t, err)
		defer win.Close()
		// then
		require.NotNil(t, win)
		assert.Equal(t, 640, win.Width())
		assert.Equal(t, 360, win.Height())
	})
	t.Run("should open two windows at the same time", func(t *testing.T) {
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// when
		win1, err := openGL.OpenWindow(640, 360)
		require.NoError(t, err)
		defer win1.Close()
		win2, err := openGL.OpenWindow(320, 180)
		require.NoError(t, err)
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
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		win1, err := openGL.OpenWindow(640, 360)
		require.NoError(t, err)
		win1.Close()
		// when
		win2, err := openGL.OpenWindow(320, 180)
		require.NoError(t, err)
		defer win2.Close()
		// then
		require.NotNil(t, win2)
		assert.Equal(t, 320, win2.Width())
		assert.Equal(t, 180, win2.Height())
	})
	t.Run("should skip nil option", func(t *testing.T) {
		openGL, err := opengl.New(mainThreadLoop)
		require.NoError(t, err)
		defer openGL.Destroy()
		// when
		win, err := openGL.OpenWindow(0, 0, nil)
		require.NoError(t, err)
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
				openGL, err := opengl.New(mainThreadLoop)
				require.NoError(t, err)
				defer openGL.Destroy()
				// when
				win, err := openGL.OpenWindow(640, 360, opengl.Zoom(test.zoom))
				require.NoError(t, err)
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
				openGL, err := opengl.New(mainThreadLoop)
				require.NoError(t, err)
				defer openGL.Destroy()
				// when
				win, err := openGL.OpenWindow(640, 360, opengl.Zoom(test.zoom))
				require.NoError(t, err)
				defer win.Close()
				// then
				require.NotNil(t, win)
				assert.Equal(t, test.expectedWidth, win.Width())
				assert.Equal(t, test.expectedHeight, win.Height())
			})
		}
	})
}

func TestOpenGL_CompileVertexShader(t *testing.T) {
	t.Run("should return error for incorrect shader", func(t *testing.T) {
		tests := map[string]string{
			"golang code": "package main\nfunc main() {}",
		}
		for name, source := range tests {
			t.Run(name, func(t *testing.T) {
				openGL, _ := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				// when
				shader, err := openGL.CompileVertexShader(source)
				// then
				assert.Error(t, err)
				assert.Nil(t, shader)
			})
		}
	})
	t.Run("should compile shader", func(t *testing.T) {
		tests := map[string]string{
			"GLSL 1.10": "void main() {}",
			"minimal": `
				#version 330 core
				void main() {
					gl_Position = vec4(0., 0., 0., 0.);
				}
				`,
		}
		for name, source := range tests {
			t.Run(name, func(t *testing.T) {
				openGL, _ := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				// when
				shader, err := openGL.CompileVertexShader(source)
				// then
				require.NoError(t, err)
				assert.NotNil(t, shader)
			})
		}
	})
	t.Run("should not panic for empty shader", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		// when
		_, _ = openGL.CompileVertexShader("")
	})
}

func TestOpenGL_CompileFragmentShader(t *testing.T) {
	t.Run("should return error for incorrect shader", func(t *testing.T) {
		tests := map[string]string{
			"golang code": "package main\nfunc main() {}",
		}
		for name, source := range tests {
			t.Run(name, func(t *testing.T) {
				openGL, _ := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				// when
				shader, err := openGL.CompileFragmentShader(source)
				assert.Error(t, err)
				assert.Nil(t, shader)
			})
		}
	})
	t.Run("should compile shader", func(t *testing.T) {
		tests := map[string]string{
			"GLSL 1.10": "void main() {}",
			"minimal": `
				#version 330 core
				void main() {}
				`,
		}
		for name, source := range tests {
			t.Run(name, func(t *testing.T) {
				openGL, _ := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				// when
				shader, err := openGL.CompileFragmentShader(source)
				require.NoError(t, err)
				assert.NotNil(t, shader)
			})
		}
	})
	t.Run("should not panic for empty shader", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		// when
		_, _ = openGL.CompileFragmentShader("")
	})
}

func TestOpenGL_LinkProgram(t *testing.T) {
	t.Run("should return error when vertex shader is nil", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		fragmentShader, _ := openGL.CompileFragmentShader("")
		// when
		program, err := openGL.LinkProgram(nil, fragmentShader)
		// then
		assert.Error(t, err)
		assert.Nil(t, program)
	})
	t.Run("should return error when fragment shader is nil", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		vertexShader, _ := openGL.CompileVertexShader("")
		// when
		program, err := openGL.LinkProgram(vertexShader, nil)
		// then
		assert.Error(t, err)
		assert.Nil(t, program)
	})
	t.Run("should return error", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		vertexShader, err := openGL.CompileVertexShader(`
								#version 330 core
								void noMain() {}
								`)
		require.NoError(t, err)
		fragmentShader, err := openGL.CompileFragmentShader(`
								#version 330 core
								void noMainEither() {}
								`)
		require.NoError(t, err)
		// when
		program, err := openGL.LinkProgram(vertexShader, fragmentShader)
		// then
		assert.Error(t, err)
		assert.Nil(t, program)
	})
	t.Run("should return program", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		vertexShader, _ := openGL.CompileVertexShader(`
								#version 330 core
								void main() {
									gl_Position = vec4(0., 0., 0., 0.);
								}
								`)
		fragmentShader, _ := openGL.CompileFragmentShader(`
								#version 330 core
								void main() {}
								`)
		// when
		program, err := openGL.LinkProgram(vertexShader, fragmentShader)
		// then
		require.NoError(t, err)
		assert.NotNil(t, program)
	})
}

func TestProgram_AcceleratedCommand(t *testing.T) {
	t.Run("should return error when passed command is nil", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		program := workingProgram(openGL)
		// when
		cmd, err := program.AcceleratedCommand(nil)
		assert.Error(t, err)
		assert.Nil(t, cmd)
	})
	t.Run("should return command", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		program := workingProgram(openGL)
		// when
		cmd, err := program.AcceleratedCommand(&commandMock{})
		require.NoError(t, err)
		assert.NotNil(t, cmd)
	})
}

func TestAcceleratedCommand_Run(t *testing.T) {
	t.Run("should execute command", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		program := workingProgram(openGL)
		texture, _ := openGL.NewTexture(1, 1)
		output := image.AcceleratedImageSelection{
			Image: texture,
		}
		tests := map[string]struct {
			selections []image.AcceleratedImageSelection
		}{
			"empty selections": {
				selections: []image.AcceleratedImageSelection{},
			},
			"one selection": {
				selections: []image.AcceleratedImageSelection{{}},
			},
			"two selections": {
				selections: []image.AcceleratedImageSelection{{}, {}},
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				command := &commandMock{}
				acceleratedCommand, _ := program.AcceleratedCommand(command)
				// when
				err := acceleratedCommand.Run(output, test.selections)
				// then
				require.NoError(t, err)
				assert.Equal(t, 1, command.executionCount)
				assert.Equal(t, test.selections, command.selections)
				assert.NotNil(t, command.renderer)
			})
		}
	})
	t.Run("should return error when command returned error", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		program := workingProgram(openGL)
		command, _ := program.AcceleratedCommand(&failingCommand{})
		// when
		err := command.Run(image.AcceleratedImageSelection{}, []image.AcceleratedImageSelection{})
		assert.Error(t, err)
	})
}

func TestOpenGL_NewFloatVertexBuffer(t *testing.T) {
	t.Run("should return error when size is negative", func(t *testing.T) {
		tests := map[string]int{
			"size -1": -1,
			"size -2": -2,
		}
		for name, size := range tests {
			t.Run(name, func(t *testing.T) {
				openGL, _ := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				// when
				buffer, err := openGL.NewFloatVertexBuffer(size)
				// then
				assert.Error(t, err)
				assert.Nil(t, buffer)
			})
		}
	})
	t.Run("should create FloatVertexBuffer", func(t *testing.T) {
		tests := map[string]int{
			"size 0": 0,
			"size 1": 1,
		}
		for name, size := range tests {
			t.Run(name, func(t *testing.T) {
				openGL, _ := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				// when
				buffer, err := openGL.NewFloatVertexBuffer(size)
				// then
				require.NoError(t, err)
				assert.NotNil(t, buffer)
				// and
				assert.Equal(t, size, buffer.Size())
			})
		}
	})
}

func TestFloatVertexBuffer_Upload(t *testing.T) {
	// TODO Finish this
	t.Run("should return error when trying to upload slice bigger than size", func(t *testing.T) {
		tests := map[string]struct {
			size int
			data []float32
		}{
			"size 0, data len 1": {
				data: []float32{1},
			},
			"size 1, data len 2": {
				data: []float32{1, 2},
			},
		}
		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				openGL, _ := opengl.New(mainThreadLoop)
				defer openGL.Destroy()
				buffer, _ := openGL.NewFloatVertexBuffer(test.size)
				// when
				err := buffer.Upload(0, test.data)
				assert.Error(t, err)
			})
		}
	})
	// TODO Finish this
	t.Run("should upload data", func(t *testing.T) {
		openGL, _ := opengl.New(mainThreadLoop)
		defer openGL.Destroy()
		buffer, _ := openGL.NewFloatVertexBuffer(1)
		input := []float32{1}
		// when
		err := buffer.Upload(0, input)
		// then
		require.NoError(t, err)
		// and
		output := make([]float32, 1)
		err = buffer.Download(0, output)
		require.NoError(t, err)
		assert.Equal(t, input, output)
	})
}

//
//func TestFloatVertexBuffer_Download(t *testing.T) {
//	t.Run("should download data", func(t *testing.T) {
//		tests := map[string][]float32{
//			"1 element":  {1.0},
//			"2 elements": {1.0, 2.0},
//		}
//		for name, data := range tests {
//			t.Run(name, func(t *testing.T) {
//				openGL, _ := opengl.New(mainThreadLoop)
//				defer openGL.Destroy()
//				buffer, _ := openGL.NewFloatVertexBuffer(data)
//				// when
//				output := make([]float32, len(data))
//				err := buffer.Download(output)
//				// then
//				require.NoError(t, err)
//				assert.Equal(t, data, output)
//			})
//		}
//	})
//	t.Run("should return error for deleted buffer", func(t *testing.T) {
//		openGL, _ := opengl.New(mainThreadLoop)
//		defer openGL.Destroy()
//		buffer, _ := openGL.NewFloatVertexBuffer([]float32{1.0})
//		buffer.Delete()
//		// when
//		err := buffer.Download(make([]float32, 1))
//		assert.Error(t, err)
//	})
//}

func workingProgram(openGL *opengl.OpenGL) *opengl.Program {
	vertexShader, _ := openGL.CompileVertexShader(`
								#version 330 core
								void main() {
									gl_Position = vec4(0., 0., 0., 0.);
								}
								`)
	fragmentShader, _ := openGL.CompileFragmentShader(`
								#version 330 core
								void main() {}
								`)
	program, _ := openGL.LinkProgram(vertexShader, fragmentShader)
	return program
}

type commandMock struct {
	executionCount int
	selections     []image.AcceleratedImageSelection
	renderer       *opengl.Renderer
}

func (f *commandMock) RunGL(renderer *opengl.Renderer, selections []image.AcceleratedImageSelection) error {
	f.executionCount++
	f.selections = selections
	f.renderer = renderer
	return nil
}

type failingCommand struct{}

func (f *failingCommand) RunGL(renderer *opengl.Renderer, selections []image.AcceleratedImageSelection) error {
	return errors.New("command failed")
}
