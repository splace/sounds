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
	l1,r1:=Stereo(NewSound(NewTone(time.Second/300, 1), time.Second*1), vector{4,1})
	l2,r2:=Stereo(Delayed(NewSound(NewTone(time.Second/300, 1), time.Second*1),time.Second*11/10), vector{-4,1})
	Encode(wavFile, 1, 16000, NewCompositor(l1,l2),NewCompositor(r1,r2))


}

func TestSpatializeStereoNarrow(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoNarrow.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	l1,r1:=Stereo(NewSound(NewTone(time.Second/300, 1), time.Second*1), vector{1,4})
	l2,r2:=Stereo(Delayed(NewSound(NewTone(time.Second/300, 1), time.Second*1),time.Second*11/10), vector{-1,4})
	Encode(wavFile, 1, 16000, NewCompositor(l1,l2),NewCompositor(r1,r2))


}

func TestSpatializeStereoVeryNarrow(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoVeryNarrow.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	l1,r1:=Stereo(NewSound(NewTone(time.Second/300, 1), time.Second*1), vector{.25,4})
	l2,r2:=Stereo(Delayed(NewSound(NewTone(time.Second/300, 1), time.Second*1),time.Second*11/10), vector{-.25,4})
	Encode(wavFile, 1, 16000, NewCompositor(l1,l2),NewCompositor(r1,r2))


}


func TestSpatializeStereoNoise(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoNoise.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	// noise is segmented at 2k, to remove high frequencies, otherwise since noise is white and procedural, it will have infinity frequency components, distroying any phasing.
	l1,r1:=Stereo(NewSound(&signals.Segmented{Signal:signals.NewNoise(),Width:signals.X(.0005)}, time.Second/10), vector{4,1})
	l2,r2:=Stereo(Delayed(NewSound(&signals.Segmented{Signal:signals.NewNoise(),Width:signals.X(.0005)}, time.Second/10),time.Second*11/100), vector{-4,1})
	Encode(wavFile, 4, 44100, NewCompositor(l1,l2),NewCompositor(r1,r2))
}


func TestSpatializeStereoNoiseNarrow(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoNoiseNarrow.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	// noise is segmented at 2k, to remove high frequencies, otherwise since noise is white and procedural, it will have infinity frequency components, distroying any phasing.
	l1,r1:=Stereo(NewSound(&signals.Segmented{Signal:signals.NewNoise(),Width:signals.X(.0005)}, time.Second/10), vector{1,4})
	l2,r2:=Stereo(Delayed(NewSound(&signals.Segmented{Signal:signals.NewNoise(),Width:signals.X(.0005)}, time.Second/10),time.Second*11/100), vector{-1,4})
	Encode(wavFile, 4, 44100, NewCompositor(l1,l2),NewCompositor(r1,r2))
}


func TestSpatializeStereoNoiseVeryNarrow(t *testing.T) {
	wavFile, err := os.Create("./test output/stereoNoiseVeryNarrow.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	// noise is sampled, in a PCM, at the same rate as the eventual encoding, to stop encode sampling from infinite frequencies, and distroying phasing.
	l1,r1:=Stereo(signals.NewPCMSignal(signals.NewNoise(), signals.X(.1), 22050, 2) , vector{.25,4})
	l2,r2:=Stereo(Delayed(signals.NewPCMSignal(signals.NewNoise(), signals.X(.1), 22050, 2),time.Second*11/100), vector{-.25,4})
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



