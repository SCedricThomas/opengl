package main

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type ShaderProgram struct {
	handle uint32
}

func NewShaderProgram(vertexShaderPath, fragmentShaderPath string) (ShaderProgram, error) {
	// Creating the shader program
	shaderProgram := gl.CreateProgram()

	// Loading the vertex shader
	vertexShader, err := NewShader(vertexShaderPath, gl.VERTEX_SHADER)
	if err != nil {
		return ShaderProgram{}, fmt.Errorf("failed to compile vertex shader: %v", err)
	}
	gl.AttachShader(shaderProgram, vertexShader.Handle())
	defer vertexShader.DeleteHandle()

	// Loading the fragment shader
	fragmentShader, err := NewShader(fragmentShaderPath, gl.FRAGMENT_SHADER)
	if err != nil {
		return ShaderProgram{}, fmt.Errorf("failed to compile fragment shader: %v", err)
	}
	gl.AttachShader(shaderProgram, fragmentShader.Handle())
	defer fragmentShader.DeleteHandle()

	// Linking the program
	gl.LinkProgram(shaderProgram)

	// Checking compilation result
	var status int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)
		fmt.Printf("%v\n", logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))
		return ShaderProgram{}, fmt.Errorf("failed to link program: %v", log)
	}

	return ShaderProgram{
		handle: shaderProgram,
	}, nil
}

func (sp ShaderProgram) Handle() uint32 {
	return sp.handle
}

func (sp ShaderProgram) Use() {
	gl.UseProgram(sp.handle)
}

func (sp ShaderProgram) Delete() {
	gl.DeleteProgram(sp.handle)
}
