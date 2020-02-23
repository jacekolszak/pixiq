package opengl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

type context struct {
	runInOpenGLThread func(func())
}

// GenBuffers generates buffer object names
func (g *context) GenBuffers(n int32, buffers *uint32) {
	g.runInOpenGLThread(func() {
		gl.GenBuffers(n, buffers)
	})
}

// BindBuffer binds a named buffer object
func (g *context) BindBuffer(target uint32, buffer uint32) {
	g.runInOpenGLThread(func() {
		gl.BindBuffer(target, buffer)
	})
}

// BufferData creates and initializes a buffer object's data store
func (g *context) BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	g.runInOpenGLThread(func() {
		gl.BufferData(target, size, data, usage)
	})
}

// BufferSubData updates a subset of a buffer object's data store
func (g *context) BufferSubData(target uint32, offset int, size int, data unsafe.Pointer) {
	g.runInOpenGLThread(func() {
		gl.BufferSubData(target, offset, size, data)
	})
}

// GetBufferSubData returns a subset of a buffer object's data store
func (g *context) GetBufferSubData(target uint32, offset int, size int, data unsafe.Pointer) {
	g.runInOpenGLThread(func() {
		gl.GetBufferSubData(target, offset, size, data)
	})
}

// DeleteBuffers deletes named buffer objects
func (g *context) DeleteBuffers(n int32, buffers *uint32) {
	g.runInOpenGLThread(func() {
		gl.DeleteBuffers(n, buffers)
	})
}

// GenVertexArrays generates vertex array object names
func (g *context) GenVertexArrays(n int32, arrays *uint32) {
	g.runInOpenGLThread(func() {
		gl.GenVertexArrays(n, arrays)
	})
}

// DeleteVertexArrays deletes vertex array objects
func (g *context) DeleteVertexArrays(n int32, arrays *uint32) {
	g.runInOpenGLThread(func() {
		gl.DeleteVertexArrays(n, arrays)
	})
}

// BindVertexArray binds a vertex array object
func (g *context) BindVertexArray(array uint32) {
	g.runInOpenGLThread(func() {
		gl.BindVertexArray(array)
	})
}

// VertexAttribPointer defines an array of generic vertex attribute data
func (g *context) VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
	g.runInOpenGLThread(func() {
		gl.VertexAttribPointer(index, size, xtype, normalized, stride, pointer)
	})
}

// EnableVertexAttribArray enables a generic vertex attribute array
func (g *context) EnableVertexAttribArray(index uint32) {
	g.runInOpenGLThread(func() {
		gl.EnableVertexAttribArray(index)
	})
}

// CreateShader creates a shader object
func (g *context) CreateShader(xtype uint32) uint32 {
	var id uint32
	g.runInOpenGLThread(func() {
		id = gl.CreateShader(xtype)
	})
	return id
}

// ShaderSource replaces the source code in a shader object
func (g *context) ShaderSource(shader uint32, count int32, xstring **uint8, length *int32) {
	g.runInOpenGLThread(func() {
		gl.ShaderSource(shader, count, xstring, length)
	})
}

// CompileShader compiles a shader object
func (g *context) CompileShader(shader uint32) {
	g.runInOpenGLThread(func() {
		gl.CompileShader(shader)
	})
}

// GetShaderiv returns a parameter from a shader object
func (g *context) GetShaderiv(shader uint32, pname uint32, params *int32) {
	g.runInOpenGLThread(func() {
		gl.GetShaderiv(shader, pname, params)
	})
}

// GetShaderInfoLog returns the information log for a shader object
func (g *context) GetShaderInfoLog(shader uint32, bufSize int32, length *int32, infoLog *uint8) {
	g.runInOpenGLThread(func() {
		gl.GetShaderInfoLog(shader, bufSize, length, infoLog)
	})
}

// DeleteShader deletes a shader object
func (g *context) DeleteShader(shader uint32) {
	g.runInOpenGLThread(func() {
		gl.DeleteShader(shader)
	})
}

// AttachShader attaches a shader object to a program object
func (g *context) AttachShader(program uint32, shader uint32) {
	g.runInOpenGLThread(func() {
		gl.AttachShader(program, shader)
	})
}

// LinkProgram links a program object
func (g *context) LinkProgram(program uint32) {
	g.runInOpenGLThread(func() {
		gl.LinkProgram(program)
	})
}

// GetProgramiv returns a parameter from a program object
func (g *context) GetProgramiv(program uint32, pname uint32, params *int32) {
	g.runInOpenGLThread(func() {
		gl.GetProgramiv(program, pname, params)
	})
}

// GetProgramInfoLog returns the information log for a program object
func (g *context) GetProgramInfoLog(program uint32, bufSize int32, length *int32, infoLog *uint8) {
	g.runInOpenGLThread(func() {
		gl.GetProgramInfoLog(program, bufSize, length, infoLog)
	})
}

// UseProgram installs a program object as part of current rendering state
func (g *context) UseProgram(program uint32) {
	g.runInOpenGLThread(func() {
		gl.UseProgram(program)
	})
}

// CreateProgram creates a program object
func (g *context) CreateProgram() uint32 {
	var program uint32
	g.runInOpenGLThread(func() {
		program = gl.CreateProgram()
	})
	return program
}

// GetActiveUniform returns information about an active uniform variable for the specified program object
func (g *context) GetActiveUniform(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	g.runInOpenGLThread(func() {
		gl.GetActiveUniform(program, index, bufSize, length, size, xtype, name)
	})
}

// GetActiveAttrib returns information about an active attribute variable for the specified program object
func (g *context) GetActiveAttrib(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	g.runInOpenGLThread(func() {
		gl.GetActiveAttrib(program, index, bufSize, length, size, xtype, name)
	})
}

// GetAttribLocation returns the location of an attribute variable
func (g *context) GetAttribLocation(program uint32, name *uint8) int32 {
	var loc int32
	g.runInOpenGLThread(func() {
		loc = gl.GetAttribLocation(program, name)
	})
	return loc
}

// Enable enables or disable server-side GL capabilities
func (g *context) Enable(cap uint32) {
	g.runInOpenGLThread(func() {
		gl.Enable(cap)
	})
}

// BindFramebuffer binds a framebuffer to a framebuffer target
func (g *context) BindFramebuffer(target uint32, framebuffer uint32) {
	g.runInOpenGLThread(func() {
		gl.BindFramebuffer(target, framebuffer)
	})
}

// Scissor defines the scissor box
func (g *context) Scissor(x int32, y int32, width int32, height int32) {
	g.runInOpenGLThread(func() {
		gl.Scissor(x, y, width, height)
	})
}

// Viewport sets the viewport
func (g *context) Viewport(x int32, y int32, width int32, height int32) {
	g.runInOpenGLThread(func() {
		gl.Viewport(x, y, width, height)
	})
}

// ClearColor specifies clear values for the color buffers
func (g *context) ClearColor(red float32, green float32, blue float32, alpha float32) {
	g.runInOpenGLThread(func() {
		gl.ClearColor(red, green, blue, alpha)
	})
}

// Clear clears buffers to preset values
func (g *context) Clear(mask uint32) {
	g.runInOpenGLThread(func() {
		gl.Clear(mask)
	})
}

// DrawArrays render primitives from array data
func (g *context) DrawArrays(mode uint32, first int32, count int32) {
	g.runInOpenGLThread(func() {
		gl.DrawArrays(mode, first, count)
	})
}

// Uniform1f specifies the value of a uniform variable for the current program object
func (g *context) Uniform1f(location int32, v0 float32) {
	g.runInOpenGLThread(func() {
		gl.Uniform1f(location, v0)
	})
}

// Uniform2f specifies the value of a uniform variable for the current program object
func (g *context) Uniform2f(location int32, v0 float32, v1 float32) {
	g.runInOpenGLThread(func() {
		gl.Uniform2f(location, v0, v1)
	})
}

// Uniform3f specifies the value of a uniform variable for the current program object
func (g *context) Uniform3f(location int32, v0 float32, v1 float32, v2 float32) {
	g.runInOpenGLThread(func() {
		gl.Uniform3f(location, v0, v1, v2)
	})
}

// Uniform4f specifies the value of a uniform variable for the current program object
func (g *context) Uniform4f(location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	g.runInOpenGLThread(func() {
		gl.Uniform4f(location, v0, v1, v2, v3)
	})
}

// Uniform1i specifies the value of a uniform variable for the current program object
func (g *context) Uniform1i(location int32, v0 int32) {
	g.runInOpenGLThread(func() {
		gl.Uniform1i(location, v0)
	})
}

// Uniform2i specifies the value of a uniform variable for the current program object
func (g *context) Uniform2i(location int32, v0 int32, v1 int32) {
	g.runInOpenGLThread(func() {
		gl.Uniform2i(location, v0, v1)
	})
}

// Uniform3i specifies the value of a uniform variable for the current program object
func (g *context) Uniform3i(location int32, v0 int32, v1 int32, v2 int32) {
	g.runInOpenGLThread(func() {
		gl.Uniform3i(location, v0, v1, v2)
	})
}

// Uniform4i specifies the value of a uniform variable for the current program object
func (g *context) Uniform4i(location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	g.runInOpenGLThread(func() {
		gl.Uniform4i(location, v0, v1, v2, v3)
	})
}

// UniformMatrix3fv specifies the value of a uniform variable for the current program object
func (g *context) UniformMatrix3fv(location int32, count int32, transpose bool, value *float32) {
	g.runInOpenGLThread(func() {
		gl.UniformMatrix3fv(location, count, transpose, value)
	})
}

// UniformMatrix4fv specifies the value of a uniform variable for the current program object
func (g *context) UniformMatrix4fv(location int32, count int32, transpose bool, value *float32) {
	g.runInOpenGLThread(func() {
		gl.UniformMatrix4fv(location, count, transpose, value)
	})
}

// ActiveTexture selects active texture unit
func (g *context) ActiveTexture(texture uint32) {
	g.runInOpenGLThread(func() {
		gl.ActiveTexture(texture)
	})
}

// BindTexture binds a named texture to a texturing target
func (g *context) BindTexture(target uint32, texture uint32) {
	g.runInOpenGLThread(func() {
		gl.BindTexture(target, texture)
	})
}

// GetIntegerv returns the value or values of the specified parameter
func (g *context) GetIntegerv(pname uint32, data *int32) {
	g.runInOpenGLThread(func() {
		gl.GetIntegerv(pname, data)
	})
}

// GenTextures generates texture names
func (g *context) GenTextures(n int32, textures *uint32) {
	g.runInOpenGLThread(func() {
		gl.GenTextures(n, textures)
	})
}

// TexImage2D specifies a two-dimensional texture image
func (g *context) TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	g.runInOpenGLThread(func() {
		gl.TexImage2D(target, level, internalformat, width, height, border, format, xtype, pixels)
	})
}

// TexParameteri sets texture parameter
func (g *context) TexParameteri(target uint32, pname uint32, param int32) {
	g.runInOpenGLThread(func() {
		gl.TexParameteri(target, pname, param)
	})
}

// GenFramebuffers generates framebuffer object names
func (g *context) GenFramebuffers(n int32, framebuffers *uint32) {
	g.runInOpenGLThread(func() {
		gl.GenFramebuffers(n, framebuffers)
	})
}

// FramebufferTexture2D attaches a level of a texture object as a logical buffer to the currently bound framebuffer object
func (g *context) FramebufferTexture2D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	g.runInOpenGLThread(func() {
		gl.FramebufferTexture2D(target, attachment, textarget, texture, level)
	})
}

// TexSubImage2D specifies a two-dimensional texture subimage
func (g *context) TexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	g.runInOpenGLThread(func() {
		gl.TexSubImage2D(target, level, xoffset, yoffset, width, height, format, xtype, pixels)
	})
}

// GetTexImage returns a texture image
func (g *context) GetTexImage(target uint32, level int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	g.runInOpenGLThread(func() {
		gl.GetTexImage(target, level, format, xtype, pixels)
	})
}

// GetError returns error information
func (g *context) GetError() uint32 {
	var code uint32
	g.runInOpenGLThread(func() {
		code = gl.GetError()
	})
	return code
}

// GoStr takes a null-terminated string returned by OpenGL and constructs a
// corresponding Go string.
func (g *context) GoStr(cstr *uint8) string {
	return gl.GoStr(cstr)
}

// Strs takes a list of Go strings (with or without null-termination) and
// returns their C counterpart.
//
// The returned free function must be called once you are done using the strings
// in order to free the memory.
//
// If no strings are provided as a parameter this function will panic.
func (g *context) Strs(strs ...string) (cstrs **uint8, free func()) {
	return gl.Strs(strs...)
}
