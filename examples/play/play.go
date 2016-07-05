// pipe to 'aplay' a scale.
package main

import (
	"io"
	"os/exec"
	"time"
)

import . "github.com/splace/sounds"

func play(s Sound) {
	out, in := io.Pipe()
	go func() {
		Encode(in,2, 44100, s)
		in.Close()
	}()
	cmd := exec.Command("aplay")
	cmd.Stdin=out 
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main(){
	play(NewSound(NewTone(time.Second/440, 1),time.Second/3))
	play(NewSound(NewTone(time.Second/495, 1),time.Second/3))
	play(NewSound(NewTone(time.Second/550, 1),time.Second/3))
	play(NewSound(NewTone(time.Second*3/1760, 1),time.Second/3))
	play(NewSound(NewTone(time.Second/660, 1),time.Second/3))
	play(NewSound(NewTone(time.Second*3/2200, 1),time.Second/3))
	play(NewSound(NewTone(time.Second/825, 1),time.Second/3))
	play(NewSound(NewTone(time.Second/880, 1),time.Second/3))

}


