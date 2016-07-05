// pipe to 'aplay' one second of middle c. 
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

/*  Hal3 Tue Jul 5 20:32:53 BST 2016 go version go1.5.1 linux/amd64
Tue Jul 5 20:32:53 BST 2016 */
/*  Hal3 Tue Jul 5 20:33:25 BST 2016 go version go1.5.1 linux/amd64
Tue Jul 5 20:33:25 BST 2016 */
/*  Hal3 Tue Jul 5 20:33:39 BST 2016 go version go1.5.1 linux/amd64
Tue Jul 5 20:33:39 BST 2016 */
/*  Hal3 Tue Jul 5 20:33:47 BST 2016 go version go1.5.1 linux/amd64
Tue Jul 5 20:33:57 BST 2016 */
/*  Hal3 Tue Jul 5 20:34:53 BST 2016 go version go1.5.1 linux/amd64
Tue Jul 5 20:34:58 BST 2016 */

