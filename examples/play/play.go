// pipe to 'aplay' a scale.
package main

import (
//	"io"
	"os/exec"
	"time"
)

import . "github.com/splace/sounds"

func play(s Sound) {
	cmd := exec.Command("aplay","--rate=44100","--format=S16_LE")
	in,err:=cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	Encode(in, 2, 44100, s)
	in.Close()
	err = cmd.Wait()
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


