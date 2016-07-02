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
	l1,r1:=Stereo(NewSound(NewTone(time.Second/800, 1), time.Second*1), vector{4,1})
	l2,r2:=Stereo(Delayed(NewSound(NewTone(time.Second/800, 1), time.Second*1),time.Second*11/10), vector{-4,1})
	Encode(wavFile, 1, 16000, NewCompositor(l1,l2),NewCompositor(r1,r2))


}

func TestSpatializeStereoNarrow(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoNarrow.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	l1,r1:=Stereo(NewSound(NewTone(time.Second/800, 1), time.Second*1), vector{1,4})
	l2,r2:=Stereo(Delayed(NewSound(NewTone(time.Second/800, 1), time.Second*1),time.Second*11/10), vector{-1,4})
	Encode(wavFile, 1, 16000, NewCompositor(l1,l2),NewCompositor(r1,r2))


}

func TestSpatializeStereoVeryNarrow(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoVeryNarrow.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	l1,r1:=Stereo(NewSound(NewTone(time.Second/800, 1), time.Second*1), vector{.25,4})
	l2,r2:=Stereo(Delayed(NewSound(NewTone(time.Second/800, 1), time.Second*1),time.Second*11/10), vector{-.25,4})
	Encode(wavFile, 1, 16000, NewCompositor(l1,l2),NewCompositor(r1,r2))


}


func TestSpatializeStereoNoiseNarrow(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoNoiseNarrow.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	l1,r1:=Stereo(NewSound(signals.NewNoise(), time.Second*1), vector{1,4})
	l2,r2:=Stereo(Delayed(NewSound(signals.NewNoise(), time.Second*1),time.Second*11/10), vector{-1,4})
	Encode(wavFile, 2, 22050, NewCompositor(l1,l2),NewCompositor(r1,r2))
}


func TestSpatializeStereoNoiseVeryNarrow(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoNoiseVeryNarrow.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	l1,r1:=Stereo(NewSound(signals.NewNoise(), time.Second*1), vector{.25,4})
	l2,r2:=Stereo(Delayed(NewSound(signals.NewNoise(), time.Second*1),time.Second*11/10), vector{-.25,4})
	Encode(wavFile, 2, 22050, NewCompositor(l1,l2),NewCompositor(r1,r2))
}


func TestSpatializeStereoFrontBackTone(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoFrontBackTone.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	l1,r1:=Stereo(NewSound(NewTone(time.Second/440, 1), time.Second*1), vector{0,4})
	l2,r2:=Stereo(Delayed(NewSound(NewTone(time.Second/440, 1), time.Second*1),time.Second*11/10), vector{0,-4})
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

/*  Hal3 Sat Jul 2 22:56:09 BST 2016 go version go1.5.1 linux/amd64
=== RUN   TestSpatializeReceeding
{0 0} 0 0
{2 0} 2 1.5707963267948966
{3 0} 3 1.5707963267948966
{4 0} 4 1.5707963267948966
{5 0} 5 1.5707963267948966
{6 0} 6 1.5707963267948966
{7 0} 7 1.5707963267948966
{8 0} 8 1.5707963267948966
{9 0} 9 1.5707963267948966
--- PASS: TestSpatializeReceeding (0.16s)
=== RUN   TestSpatializeStereo
{1 -4.25} 4.366062299143245 2.910501986393896
{1 3.75} 3.881043674065006 0.26060239174734096
{1 3.75} 3.881043674065006 0.26060239174734096
{1 -4.25} 4.366062299143245 2.910501986393896
--- PASS: TestSpatializeStereo (0.21s)
=== RUN   TestSpatializeStereoNarrow
{4 -1.25} 4.190763653560053 1.873681195169868
{4 0.75} 4.0697051490249265 1.3854483767992019
{4 0.75} 4.0697051490249265 1.3854483767992019
{4 -1.25} 4.190763653560053 1.873681195169868
--- PASS: TestSpatializeStereoNarrow (0.23s)
=== RUN   TestSpatializeStereoVeryNarrow
{4 -0.5} 4.031128874149275 1.695151321341658
{4 0} 4 1.5707963267948966
{4 0} 4 1.5707963267948966
{4 -0.5} 4.031128874149275 1.695151321341658
--- PASS: TestSpatializeStereoVeryNarrow (0.17s)
=== RUN   TestSpatializeStereoNoiseNarrow
{4 -1.25} 4.190763653560053 1.873681195169868
{4 0.75} 4.0697051490249265 1.3854483767992019
{4 0.75} 4.0697051490249265 1.3854483767992019
{4 -1.25} 4.190763653560053 1.873681195169868
--- PASS: TestSpatializeStereoNoiseNarrow (0.27s)
=== RUN   TestSpatializeStereoNoiseVeryNarrow
{4 -0.5} 4.031128874149275 1.695151321341658
{4 0} 4 1.5707963267948966
{4 0} 4 1.5707963267948966
{4 -0.5} 4.031128874149275 1.695151321341658
--- PASS: TestSpatializeStereoNoiseVeryNarrow (7.51s)
=== RUN   TestSpatializeStereoFrontBackTone
{4 -0.25} 4.00780488547035 1.6332151367908538
{4 -0.25} 4.00780488547035 1.6332151367908538
{-4 -0.25} 4.00780488547035 -1.6332151367908538
{-4 -0.25} 4.00780488547035 -1.6332151367908538
--- PASS: TestSpatializeStereoFrontBackTone (0.29s)
PASS
ok  	_/home/simon/Dropbox/github/working/sound	8.853s
Sat Jul 2 22:56:20 BST 2016 */

