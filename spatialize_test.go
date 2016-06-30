package sound

import (
	"testing"
	"time"
	"os"
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
	s1,s2:=Stereo(NewSound(NewTone(time.Second/440, 1), time.Second*2), vector{2,2})
	s3,s4:=Stereo(NewSound(NewTone(time.Second/1200, 1), time.Second), vector{-2,2})
	Encode(wavFile, 1, 16000, NewCompositor(s1,s3),NewCompositor(s2,s4))


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



