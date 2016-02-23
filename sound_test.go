//Formatted:Tue Feb 2 21:16:21 GMT 2016
package sound

import (
	"../signals"
	//"fmt"
	"os"
	"testing"
	"time"
)

func TestSaveTone(t *testing.T) {
	wavFile, err := os.Create("tone.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s1 := signals.NewTone(signals.UnitX/200, 1)
	signals.Encode(wavFile, s1, signals.UnitX, 44100, 2)
}

func TestSaveSound(t *testing.T) {
	wavFile, err := os.Create("noise.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s1 := NewSound(signals.NewNoise(), time.Second/2)
	Encode(wavFile, s1, 44100, 2)
}

func TestSaveFlattenedSound(t *testing.T) {
	wavFile, err := os.Create("fnoise.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s1 := NewSound(signals.NewSegmented(signals.NewNoise(), signals.X(time.Millisecond/2)), time.Second*2)
	Encode(wavFile, s1, 44100, 2)
}

func TestSaveNote(t *testing.T) {
	noteNumber := MidiNoteNumber(OctaveNumber["one-line"], SemitoneNumber["C"])
	wavFile, err := os.Create("middlec.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s1 := NewNoteMidi(noteNumber, 2000*ms, 1)
	Encode(wavFile, s1, 8000, 1)
}
func TestSaveWavSoundAfterSound(t *testing.T) {
	wavFile, err := os.Create("tones.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s1 := NewNote(NewTone(Period(4, SemitoneNumber["D"]), 1), time.Second/3)
	s2 := After(s1, NewNote(NewTone(Period(4, SemitoneNumber["E"]), 1), time.Second/3))
	s3 := After(s2, NewNote(NewTone(Period(4, SemitoneNumber["C"]), 1), time.Second/3))
	s4 := After(s3, NewNote(NewTone(Period(3, SemitoneNumber["C"]), 1), time.Second/3))
	s5 := After(s4, NewNote(NewTone(Period(3, SemitoneNumber["G"]), 1), time.Second*2/3))
	Encode(wavFile, NewComposition(s1, s2, s3, s4, s5), 44100, 1)

}

func TestSaveVibrato(t *testing.T) {
	wavFile, err := os.Create("vibrato.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s := NewNoteMidi(MidiNoteNumber(OctaveNumber["one-line"], SemitoneNumber["C"]), 2000*ms, 1)
	sm := NewNoteMidi(MidiNoteNumber(OctaveNumber["great"], SemitoneNumber["C"]), 2000*ms, 1)
	Encode(wavFile, Modulated(s, sm, 1*ms), 8000, 1)
}
