# sounds
Go language generation and manipulation of sounds, built on github.com/splace/signals package.
 
Status: (Beta :- stabilising API)

Overview/docs: [![GoDoc](https://godoc.org/github.com/splace/sounds?status.svg)](https://godoc.org/github.com/splace/sounds)
uses Signals: [![GoDoc](https://godoc.org/github.com/splace/signals?status.svg)](https://godoc.org/github.com/splace/signals) 

Installation:

     go get github.com/splace/sounds   

Example: play a note. ("aplay" command, or something like it, doesn't seem to exist on windows.)

	package main

	import (
		"io"
		"os/exec"
		"time"
	)

	import . "github.com/splace/sounds"

	func play(s Sound) {
		cmd := exec.Command("aplay")
		out, in := io.Pipe()
		go func() {
			Encode(in, 2, 44100, s)
			in.Close()
		}()
		cmd.Stdin=out 
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}

	func main(){
		play(NewSound(NewTone(time.Second/440, 1),time.Second))
	}


Example: A tune with sampled notes.

[twinkle twinkle little star](https://github.com/splace/sounds/blob/master/test%20output/hNotes.wav)


