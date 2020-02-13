package opengl

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func compileProgram(vertexShaderSrc, fragmentShaderSrc string) (*program, error) {
	vertexShader, err := compileVertexShader(vertexShaderSrc)
	if err != nil {
		return nil, err
	}
	defer vertexShader.delete()

	fragmentShader, err := compileFragmentShader(fragmentShaderSrc)
	if err != nil {
		return nil, err
	}
	defer fragmentShader.delete()

	program, err := linkProgram(vertexShader, fragmentShader)
	if err != nil {
		return nil, err
	}

	return program, nil
}

type program struct {
	id uint32
}

func (p *program) use() {
	gl.UseProgram(p.id)
}

func (p *program) activeUniformLocations() map[string]int32 {
	locationsByName := map[string]int32{}
	var count, bufSize, length, nameMaxLength int32
	var xtype uint32
	gl.GetProgramiv(p.id, gl.ACTIVE_UNIFORM_MAX_LENGTH, &nameMaxLength)
	name := make([]byte, nameMaxLength)
	gl.GetProgramiv(p.id, gl.ACTIVE_UNIFORMS, &count)
	for location := int32(0); location < count; location++ {
		gl.GetActiveUniform(p.id, uint32(location), nameMaxLength, &bufSize, &length, &xtype, &name[0])
		goName := gl.GoStr(&name[0])
		locationsByName[goName] = location
	}
	return locationsByName
}

type attribute struct {
	typ  Type
	name string
}

func (p *program) attributes() []attribute {
	var count, bufSize, length, nameMaxLength int32
	var xtype uint32
	gl.GetProgramiv(p.id, gl.ACTIVE_ATTRIBUTE_MAX_LENGTH, &nameMaxLength)
	name := make([]byte, nameMaxLength)
	gl.GetProgramiv(p.id, gl.ACTIVE_ATTRIBUTES, &count)
	attributes := make([]attribute, count)
	for location := int32(0); location < count; location++ {
		gl.GetActiveAttrib(p.id, uint32(location), nameMaxLength, &bufSize, &length, &xtype, &name[0])
		attributes[location] = attribute{typ: valueOf(xtype), name: gl.GoStr(&name[0])}
	}
	return attributes
}

func linkProgram(shaders ...*shader) (*program, error) {
	programID := gl.CreateProgram()
	for _, shader := range shaders {
		gl.AttachShader(programID, shader.id)
	}
	gl.LinkProgram(programID)
	var success int32
	gl.GetProgramiv(programID, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var infoLogLen int32
		gl.GetProgramiv(programID, gl.INFO_LOG_LENGTH, &infoLogLen)
		infoLog := make([]byte, infoLogLen)
		if infoLogLen > 0 {
			gl.GetProgramInfoLog(programID, infoLogLen, nil, &infoLog[0])
		}
		return nil, fmt.Errorf("error linking program: %s", string(infoLog))
	}
	return &program{
		id: programID,
	}, nil
}

type shader struct {
	id uint32
}

func compileVertexShader(src string) (*shader, error) {
	return compileShader(gl.VERTEX_SHADER, src)
}

func compileFragmentShader(src string) (*shader, error) {
	return compileShader(gl.FRAGMENT_SHADER, src)
}

func compileShader(xtype uint32, src string) (*shader, error) {
	if src == "" {
		src = " "
	}
	shaderID := gl.CreateShader(xtype)
	srcXString, free := gl.Strs(src)
	defer free()
	length := int32(len(src))
	gl.ShaderSource(shaderID, 1, srcXString, &length)
	gl.CompileShader(shaderID)
	var success int32
	gl.GetShaderiv(shaderID, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLen int32
		gl.GetShaderiv(shaderID, gl.INFO_LOG_LENGTH, &logLen)
		infoLog := make([]byte, logLen)
		if logLen > 0 {
			gl.GetShaderInfoLog(shaderID, logLen, nil, &infoLog[0])
		}
		return nil, fmt.Errorf("error compiling shader: %s", string(infoLog))
	}
	return &shader{id: shaderID}, nil
}

func (s *shader) delete() {
	gl.DeleteShader(s.id)
}

const vertexShaderSrc = `
#version 330 core

layout(location = 0) in vec2 vertexPosition;
layout(location = 1) in vec2 texturePosition;

out vec2 position;

void main() {
	gl_Position = vec4(vertexPosition, 0.0, 1.0);
	position = texturePosition;
}
`

const fragmentShaderSrc = `
#version 330 core

in vec2 position;

out vec4 color;

uniform sampler2D tex;

void main() {
	color = texture(tex, position);
}
`
