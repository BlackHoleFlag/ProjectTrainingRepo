package main

import (
	"fmt"
	"os"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize sdl: %s\n", err)
		os.Exit(1)
	}

	window, err := sdl.CreateWindow(
		"GO bear go!",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
			os.Exit(2)
		}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	defer renderer.Destroy()
	renderer.Clear()

	var event sdl.Event
	isRunning := true
	
	plr, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("Creating player: ", err)
		return
	}

	for isRunning {
		// handle events, in this case escape key and close window
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				isRunning = false
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					isRunning = false
				}
			}
		}

		renderer.SetDrawColor(255,255,255,255)
		renderer.Clear()

		plr.draw(renderer)
		plr.update()

		renderer.Present()
	}	
}