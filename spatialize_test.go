package sound

import (
	"testing"
	"time"
	"os"
	"github.com/splace/signals"
)


func TestSpatializeReceeding(t *testing.T) {
	s:=NewCompositor()
	s.Composite=append(s.Composite,Spatialized(NewSound(NewTone(time.Second/440, 1), time.Second/2), vector{0, 0}))
	for location:=(vector{2,0}); location.x<10;location.x=location.x+1{
		fs := Spatialized(NewSound(NewTone(time.Second/440, 1), time.Second/2), location)
		s.Composite=append(s.Composite,After(s.Composite[len(s.Composite)-1].(Sound),fs))
	}
	wavFile, err := os.Create("./test output/receeding.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile, 1, 8000, s)

}



func TestSpatializeStereo(t *testing.T) {
	wavFile, err := os.Create("./test output/stereo.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	l1,r1:=Stereo(NewSound(NewTone(time.Second/800, 1), time.Second*1), vector{1,4})
	l2,r2:=Stereo(Delayed(NewSound(NewTone(time.Second/800, 1), time.Second*1),time.Second*11/10), vector{-1,4})
	Encode(wavFile, 1, 16000, NewCompositor(l1,l2),NewCompositor(r1,r2))


}


func TestSpatializeStereoNoise(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoNoise.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	l1,r1:=Stereo(NewSound(signals.NewNoise(), time.Second*1), vector{4,2})
	l2,r2:=Stereo(Delayed(NewSound(signals.NewNoise(), time.Second*1),time.Second*11/10), vector{-4,2})
	Encode(wavFile, 2, 22050, NewCompositor(l1,l2),NewCompositor(r1,r2))


}


func TestSpatializeStereoTone(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoTone.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	l1,r1:=Stereo(NewSound(NewTone(time.Second/440, 1), time.Second*1), vector{4,2})
	l2,r2:=Stereo(Delayed(NewSound(NewTone(time.Second/440, 1), time.Second*1),time.Second*11/10), vector{-4,2})
	Encode(wavFile, 2, 22050, NewCompositor(l1,l2),NewCompositor(r1,r2))


}

//
//func TestSpatializePanning(t *testing.T) {
//	s1:=NewCompositor()
//	s1.Composite=append(s1.Composite,Spatialized(NewSound(NewTone(time.Second/440, 1), time.Second/2), vector{1,0}))
//	for location:=(vector{1,1}); location.y<=5;location.y=location.y+1{
//		fs := Spatialized(NewSound(NewTone(time.Second/440, 1), time.Second/2), location)
//		s1.Composite=append(s1.Composite,After(s1.Composite[len(s1.Composite)-1].(Sound),fs))
//	}
//	s2:=NewCompositor()
//	for si :=range(s1.Composite){
//		s2.Composite=append(s2.Composite,Spatialized([]signals.Signal(s1.Composite[si].(signals.Modulated)).[0], location))
//	}
//	//s2:=NewCompositor(signals.PromoteToSignals(s1)...)
//	wavFile, err := os.Create("./test output/panning.wav")
//	if err != nil {
//		panic(err)
//	}
//	defer wavFile.Close()
//	Encode(wavFile, 1, 8000, s1,s2)
//
//}


/*  Hal3 Fri Jul 1 00:28:27 BST 2016 go version go1.5.1 linux/amd64
=== RUN   TestSpatializeReceeding
--- PASS: TestSpatializeReceeding (0.17s)
=== RUN   TestSpatializeStereo
--- PASS: TestSpatializeStereo (7.10s)
PASS
ok  	_/home/simon/Dropbox/github/working/sound	7.305s
Fri Jul 1 00:28:36 BST 2016 */
/*  Hal3 Fri Jul 1 00:38:20 BST 2016 go version go1.5.1 linux/amd64
=== RUN   TestSpatializeReceeding
--- PASS: TestSpatializeReceeding (0.19s)
=== RUN   TestSpatializeStereo
--- PASS: TestSpatializeStereo (0.21s)
=== RUN   TestSpatializeStereoNoise
--- PASS: TestSpatializeStereoNoise (7.17s)
=== RUN   TestSpatializeStereoTone
--- PASS: TestSpatializeStereoTone (0.29s)
PASS
ok  	_/home/simon/Dropbox/github/working/sound	7.884s
Fri Jul 1 00:38:30 BST 2016 */

