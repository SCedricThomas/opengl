# OpenGL meetup

Dependencies required:
```
sudo apt install libgl1-mesa-dev xorg-dev
```

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

Core vs Immediate ?
- Immediate => Legacy mode easy to use
- Core => manually and harder 
- Allows to use the rendering pipeline at it's full extend 

How to use OpenGL ?
- It comes in 2 parts / libraries:
    - The context management library => Window creation and manipulation 
    - The binding library searching and loading the specific function in the graphic card driver
- Creation of the context 
- Using the OpenGL bindings to draw on it  

OpenGL resources
----------------

VAO / VBO ?

What is a Shader ? What is a program ?

Graphic pipeline ?
 - Vertex:
 - Fragment:

