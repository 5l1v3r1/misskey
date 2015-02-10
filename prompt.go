package main

import (
	"github.com/unixpickle/gogui"
	"strings"
)

type Prompt struct {
	RunesDone int
	Words     []string
	WordsDone int
}

// NewPrompt creates a prompt from a string.
func NewPrompt(text string) *Prompt {
	// TODO: more sophisticated word splitting.
	words := strings.Split(text, " ")
	return &Prompt{Words: words}
}

// Complete returns true if the prompt was fully entered.
func (p *Prompt) Complete() bool {
	word := []rune(p.Words[p.WordsDone])
	return p.WordsDone == len(p.Words)-1 && p.RunesDone == len(word)
}

// Draw draws the prompt in a context.
func (p *Prompt) Draw(c gogui.DrawContext, maxWidth float64) {
	c.SetFont(18, "Helvetica")

	var x, y float64 = 10, 10
	const lineHeight = 25
	const cursorHeight = 20

	for i, word := range p.Words {
		w := WordWidth(c, word)
		if x+w+10 > maxWidth {
			x = 10
			y += lineHeight
		}
		// Draw the word
		for j, ch := range word {
			if i < p.WordsDone || (i == p.WordsDone && j < p.RunesDone) {
				c.SetFill(gogui.Color{0, 0, 0, 1})
			} else {
				c.SetFill(gogui.Color{0.5, 0.5, 0.5, 1})
			}
			c.FillText(string(ch), x, y)
			width, _ := c.TextSize(string(ch))

			if i == p.WordsDone && j == p.RunesDone {
				// Draw a cursor.
				c.SetFill(gogui.Color{0, 0, 0, 1})
				c.FillRect(gogui.Rect{x, y, 2, cursorHeight})
			}

			x += width
		}
		if i == p.WordsDone && p.RunesDone == len([]rune(word)) {
			// Draw a cursor.
			c.SetFill(gogui.Color{0, 0, 0, 1})
			c.FillRect(gogui.Rect{x, y, 2, cursorHeight})
		}
		x += 10
	}
}

// HandleKey should be called to check a key press against the prompt.
func (p *Prompt) HandleKey(c gogui.KeyEvent) {
	currentRunes := []rune(p.Words[p.WordsDone])

	// If they are at the end of the word, they must hit space.
	if len(currentRunes) == p.RunesDone {
		if c.CharCode != 0x20 {
			p.Reset()
		} else {
			p.WordsDone++
			p.RunesDone = 0
		}
		return
	}

	expecting := currentRunes[p.RunesDone]
	if int(expecting) != c.CharCode {
		p.Reset()
		return
	}
	p.RunesDone++
}

// Reset resets the prompt to have nothing completed.
func (p *Prompt) Reset() {
	p.WordsDone = 0
	p.RunesDone = 0
}

// WordWidth computes the width of a given word the way it would be drawn to the
// context.
func WordWidth(c gogui.DrawContext, word string) float64 {
	var res float64
	for _, ch := range word {
		w, _ := c.TextSize(string(ch))
		res += w
	}
	return res
}
