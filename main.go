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
	NewPrompt("We promptly judged antique ivory buckles for the next prize."),
	NewPrompt("Sixty zippers were quickly picked from the woven jute bag."),
	NewPrompt("Crazy Fredrick bought many very exquisite opal jewels."),
	NewPrompt("Jump by vow of quick, lazy strength in Oxford."),
	NewPrompt("The five boxing wizards jump quickly."),
}
var currentPrompt = 0
var mistakeCount = 0
var correctCount = 0
var instructions = "Type & don't make mistakes"

func drawCanvas(c gogui.DrawContext) {
	// Draw a white backdrop
	c.SetFill(gogui.Color{1, 1, 1, 1})
	c.FillRect(gogui.Rect{0, 0, WindowSize, WindowSize})
	
	// Draw the prompt
	prompts[currentPrompt].Draw(c, WindowSize)
	
	// Draw stats and instructions
	c.SetFill(gogui.Color{0, 0, 0, 1})
	c.SetFont(18, "Helvetica")
	w, _ := c.TextSize(instructions)
	c.FillText(instructions, (WindowSize-w)/2, WindowSize-120)
	c.SetFont(16, "Helvetica")
	c.FillText("Mistakes: "+strconv.Itoa(mistakeCount), 10, WindowSize-30)
	c.FillText("Correct: "+strconv.Itoa(correctCount), 10, WindowSize-60)
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
		} else {
			correctCount++
		}
		if p.Complete() {
			p.Reset()
			currentPrompt = (currentPrompt+1) % len(prompts)
		}
		canvas.NeedsUpdate()
	})
}
