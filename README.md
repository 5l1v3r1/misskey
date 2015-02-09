# misskey

This is a simple Desktop typing test which starts over every time the typist makes a mistake. I was motivated to make this after seeing my friend try to type in Dvorak&mdash;he makes frequent mistakes instead of actually *learning* where the keys are.

# Implementation

The current implementation of *misskey* is written in Go with [gogui](https://github.com/unixpickle/gogui) for its interface. The test sentences and keyboard layout are hard-coded and generic.

# Usage

Everything in this section requires that you have the Go programming language installed and configured with a `$GOPATH`.

To use misskey, you must install gogui first:

    go get github.com/unixpickle/gogui
    go install github.com/unixpickle/gogui

To run misskey itself, you should change to the misskey directory and run it like so:

    go run *.go

# TODO

 * Add a cursor underneath the current letter.
 * Add some on-screen instructions
 * Display the number of correct and incorrect letters they have typed.
 * Possibly display their WPM
