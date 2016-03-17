# sounds
Go language generation and manipulation of sounds, using github.com/splace/signals package.
 
Status: (Beta :- stabilising API)

Overview: (see godoc reference below)

Installation:

     go get github.com/splace/sounds   

Example: play a note.

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
			Encode(in, s, 44100,2)
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


Example: tune with electronic organ sampled notes.

[twinkle twinkle little star](https://github.com/splace/sounds/blob/master/test%20output/hNotes.wav)


[![GoDoc](https://godoc.org/github.com/splace/sounds?status.svg)](https://godoc.org/github.com/splace/sounds)

