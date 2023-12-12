package main

import (
	"runtime"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl" // core OpenGL
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	// Initializing the context manipulation library
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// Creating a OpenGL context
	window, err := glfw.CreateWindow(640, 480, "OPENGL", nil, nil)
	if err != nil {
		panic(err)
	}

	// Using the created context as the current context
	window.MakeContextCurrent()

	// OpenGL binding library initialization
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// Creation of a shader program
	shaderProgram, err := NewShaderProgram(
		"./shaders/vertex.vs",
		"./shaders/fragment.fs",
	)
	if err != nil {
		panic(err)
	}
	defer shaderProgram.Delete()

	vertices := []float32{
		// positions         // colors
		0.5, -0.5, 0.0, 1.0, 0.0, 0.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 1.0, 0.0, // bottom left
		0.0, 0.5, 0.0, 0.0, 0.0, 1.0, // top
	}

	var VBO, VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.GenBuffers(1, &VBO)
	defer gl.DeleteVertexArrays(1, &VAO)
	defer gl.DeleteBuffers(1, &VBO)

	// bind the Vertex Array Object first, then bind and set vertex buffer(s), and then configure vertex attributes(s).
	gl.BindVertexArray(VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(unsafe.Sizeof(float32(0))), gl.Ptr(vertices), gl.STATIC_DRAW)

	// location binding
	// size of the data
	// type of the data
	// normalized boolean
	// stride in bytes
	// offset from the beginning in bytes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*int32(unsafe.Sizeof(float32(0))), nil)
	// Enabling the attribute
	gl.EnableVertexAttribArray(0)

	// color binding
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 6*int32(unsafe.Sizeof(float32(0))), 3*unsafe.Sizeof(float32(0)))
	gl.EnableVertexAttribArray(1)

	// context rendering loop
	for !window.ShouldClose() {

		// render
		gl.ClearColor(0.4, 0.4, 0.4, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// draw our example triangle
		gl.BindVertexArray(VAO)           // bind the VAO to draw from
		shaderProgram.Use()               // bind the program used to render the VAO content
		gl.DrawArrays(gl.TRIANGLES, 0, 3) // draw from the bound VAO

		// Swap rendering buffers to be able to see what we are drawing
		window.SwapBuffers()

		// Poll events occurring on the context window
		glfw.PollEvents()
	}
}
