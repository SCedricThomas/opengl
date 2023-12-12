package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	shaderType uint32
	handle     uint32
}

func NewShader(filename string, shaderType uint32) (Shader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Shader{}, fmt.Errorf("failed to open file: %v", err)
	}
	source, err := io.ReadAll(file)
	if err != nil {
		return Shader{}, fmt.Errorf("failed to read: %v", err)
	}
	handle := gl.CreateShader(shaderType)
	csources, free := gl.Strs(string(source) + "\x00")
	gl.ShaderSource(handle, 1, csources, nil)
	free()
	gl.CompileShader(handle)
	var status int32
	gl.GetShaderiv(handle, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(handle, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(handle, logLength, nil, gl.Str(log))
		return Shader{}, fmt.Errorf("failed to compile %v: %v", filename, log)
	}
	return Shader{
		shaderType: shaderType,
		handle:     handle,
	}, nil
}

func (s Shader) DeleteHandle() {
	gl.DeleteShader(s.handle)
}

func (s Shader) Handle() uint32 {
	return s.handle
}
