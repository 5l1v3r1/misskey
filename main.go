package main

import (
	"fmt"
	"github.com/unixpickle/gogui"
	"os"
	"strconv"
)

const WindowSize = 400

var prompts = []*Prompt{
	NewPrompt("The quick brown fox jumps over the lazy dog."),
	NewPrompt("Pack my box with five dozen liquor jugs."),
}
var currentPrompt = 0
var mistakeCount = 0

func drawCanvas(c gogui.DrawContext) {
	// Draw a white backdrop
	c.SetFill(gogui.Color{1, 1, 1, 1})
	c.FillRect(gogui.Rect{0, 0, WindowSize, WindowSize})
	
	// Draw the prompt
	prompts[currentPrompt].Draw(c, WindowSize)
	
	// Draw the number of mistakes
	c.SetFont(18, "Helvetica")
	c.SetFill(gogui.Color{0, 0, 0, 1})
	c.FillText("Mistakes: "+strconv.Itoa(mistakeCount), 10, WindowSize-30)
}

func main() {
	gogui.RunOnMain(setup)
	gogui.Main(&gogui.AppInfo{Name: "MissKey"})
}

func setup() {
	bounds := gogui.Rect{0, 0, WindowSize, WindowSize}
	
	// Create the window.
	window, err := gogui.NewWindow(bounds)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	window.SetTitle("MissKey")
	window.SetCloseHandler(func() {
		os.Exit(0)
	})
	
	// Create the canvas.
	canvas, err := gogui.NewCanvas(bounds)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	canvas.SetDrawHandler(drawCanvas)
	window.Add(canvas)
	
	// Present the window.
	window.Center()
	window.Show()
	
	canvas.NeedsUpdate()
	
	window.SetKeyPressHandler(func(e gogui.KeyEvent) {
		p := prompts[currentPrompt]
		if !p.HandleKey(e) {
			mistakeCount++
		}
		if p.Complete() {
			p.Reset()
			currentPrompt = (currentPrompt+1) % len(prompts)
		}
		canvas.NeedsUpdate()
	})
}
