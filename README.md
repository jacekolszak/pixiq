# Pixiq

[![CircleCI](https://circleci.com/gh/jacekolszak/pixiq.svg?style=svg)](https://circleci.com/gh/jacekolszak/pixiq)
[![GoDoc](https://godoc.org/github.com/jacekolszak/pixiq?status.svg)](http://godoc.org/github.com/jacekolszak/pixiq)
[![Go Report Card](https://goreportcard.com/badge/github.com/jacekolszak/pixiq)](https://goreportcard.com/report/github.com/jacekolszak/pixiq)

Create Pixel Art games in Golang with fun and ease.

## What you can do with Pixiq? (or you will be able to do as soon as it is released ;)

+ draw images on a screen in real time using your favourite [Go programming language](https://golang.org/)
+ manipulate every single pixel directly or with the use of tools
+ handle user input

## What is Pixel Art?

+ is the art of making digital images where the creator place every single pixel of the image
+ it is all about limits: low resolution, limited palette, no rotation besides square angles and only integer scaling
+ no automatic anti-aliasing, no filters, blur and other fancy tools
+ more information can be found in [The Pixel Art Tutorial](http://pixeljoint.com/forum/forum_posts.asp?TID=11299)

## Installation

+ [Go 1.13+](https://golang.org/dl/)
+ Install dependencies:
    + Ubuntu/Debian: `sudo apt-get install libgl1-mesa-dev xorg-dev`
    + CentOS/Fedora: `sudo yum install libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel`
    + MacOS: `xcode-select --install`
+ `go get -u github.com/jacekolszak/pixiq`

## Hello world!

```go
package main

import (
	"github.com/jacekolszak/pixiq/colornames"
	"github.com/jacekolszak/pixiq/loop"
	"github.com/jacekolszak/pixiq/opengl"
)

func main() {
	opengl.RunOrDie(func(gl *opengl.OpenGL) {
		window, err := gl.OpenWindow(80, 16, opengl.Zoom(5))
		if err != nil {
			panic(err)
		}
		loop.Run(window, func(frame *loop.Frame) {
			screen := frame.Screen()
			screen.SetColor(40, 8, colornames.White)
		})
	})
}
```

More examples you can find in [examples](examples) directory.

## Project status

At the moment I have a vision how the API should look like and how it should be programmed.
I have made a quick and dirty implementation as a proof of concept. 
Now it is time to start from scratch in order to deliver the highest possible quality.

The project is using [semantic versioning](https://semver.org/). Current version 
is `0.X.Y` which basically means that future versions may introduce incompatible 
API changes. More about architecture can be found in [architecture document](docs/architecture.md).

## Project goals

+ Create Go API which is just fun to use. It will provide tools known from Pixel Art software.
+ Create Development Tools similar to Chrome Developer Tools to support the development process
+ Make it fast - image manipulation requires a lot of computation, therefore Pixiq should be well optimized
