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


