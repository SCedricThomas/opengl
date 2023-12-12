# OpenGL meetup

Dependencies required:
```
sudo apt install libgl1-mesa-dev xorg-dev
```

How to learn OpenGL online ?
https://learnopengl.com/


----

OpenGL introduction
-------------------

What is OpenGL ?
- OpenGL is an API specification that provides us with a large set of functions that we can use to manipulate graphics and images.
- This specification is implemented in graphics card driver.
- It's a platform agnostic specification allowing to use the same code on several platform to render the same thing 

How is OpenGL working ?
- It's a huge state machine
- You can create resources in the state machine and bind them to use them
- It only manipulates 3D coords

Core vs Immediate ?
- Immediate => Legacy mode easy to use
- Core => manually and harder 
- Allows to use the rendering pipeline at it's full extend 

How to use OpenGL ?
- First we need to create a window to draw on it 
- Then we can draw on it by using the OpenGL bindings  

=> It comes in 2 parts / libraries we need to initialize:
    - The context management library (window creation and manipulation) 
    - The binding library (searching and loading the specific function in the graphic card driver)


OpenGL resources
----------------

Graphic pipeline ?
The graphic pipeline is a succession of steps allowing to transform 3D coords in 2D coords.
Each steps output is the input of the next one. I only know 2 steps:
 - Vertex: It transforms 3D coordinates into different 3D coordinates. 
 - Fragment: It calculates the final color of a pixel and this is usually the stage where all the advanced OpenGL effects occur
You can configure some steps with small programs => Shaders

What is a Shader ? What is a program ?
- A shader is a piece of code written in GLSL running on the GC allowing to transform the data directly on the GC
- A program is a group of shaders we can load on the CG that will be used when rendering a pixel 

What is a VBO ? What is a VAO ?
- A VBO (Vertex Buffer Object) is an OpenGL buffer storing data that will be sent to the CG to be rendered
- A VAO (Vertex Array Object) is an object allowing us to precise the structure of the data stored in a VBO   
