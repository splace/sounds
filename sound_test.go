package sound

import (
	"fmt"
	"github.com/splace/signals" //"../signals"//
	"os"
	"testing"
	"time"
)

func TestSaveTone(t *testing.T) {
	wavFile, err := os.Create("./test output/tone.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s1 := signals.Sine{signals.X(.005)}
	signals.Encode(wavFile, 2, 44100, signals.X(1), s1)
}

func TestSaveSound(t *testing.T) {
	wavFile, err := os.Create("./test output/noise.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s1 := NewSound(signals.NewNoise(), time.Second/2)
	Encode(wavFile, s1, 44100, 2)
}

func TestSaveFlattenedSound(t *testing.T) {
	wavFile, err := os.Create("./test output/fnoise.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s1 := NewSound(signals.NewSegmented(signals.NewNoise(), signals.X(.0005)), time.Second/2)
	Encode(wavFile, s1, 44100, 2)
}

func TestSaveNote(t *testing.T) {
	noteNumber := MidiNote("one-line","C")
	wavFile, err := os.Create("./test output/middlec.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s1 := NewMidiNote(noteNumber, 2000*ms, 1)
	Encode(wavFile, s1, 8000, 1)
}
func TestLoad(t *testing.T) {
	stream, err := os.Open("middlec.wav")
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	noises, err := signals.Decode(stream)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(noises))
}

func TestLoadChannels(t *testing.T) {
	stream, err := os.Open("pcm0808s.wav")
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	noises, err := signals.Decode(stream)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(noises))
}

func TestSaveSignal(t *testing.T) {
	wavFile, err := os.Create("./test output/delayedlimitedsquaresignal.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s1 := Delayed(NewSound(signals.Square{signals.X(.1)}, 1000*ms), 2000*ms)
	Encode(wavFile, s1, 8000, 1)
}

func TestSaveModifiedNote(t *testing.T) {
	wavFile2, err := os.Create("./test output/NoteModded.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile2.Close()
	s2 := Spedup(NewSound(NewTone(time.Millisecond, 1), time.Second), .264) // makes a middle c
	Encode(wavFile2, s2, 8000, 1)
}

func TestSaveModifiedWav(t *testing.T) {
	stream, err := os.Open("8k8bitpcm.wav")
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	wavFile, err := os.Create("./test output/8k8bitpcmSpedup.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	noises, err := signals.Decode(stream)
	//noises[0].Interpolate = true // interpolation because the save frequency, 44.1k, is going to be much more than stored, 8k.
	Encode(wavFile, Spedup(noises[0].(Sound), 1.2), 44100, 1)
	//wav.Encode(noises[0], wavFile, 44100,1)
}
func TestSaveWavSoundAfterSound(t *testing.T) {
	wavFile, err := os.Create("./test output/tones.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()

	s1 := NewNote(NewTone(Period(4, "D"), 1), time.Second/3)
	s2 := After(s1, NewNote(NewTone(Period(4, "E"), 1), time.Second/3))
	s3 := After(s2, NewNote(NewTone(Period(4, 0), 1), time.Second/3))
	s4 := After(s3, NewNote(NewTone(Period("small", 0), 1), time.Second/3))
	s5 := After(s4, NewNote(NewTone(Period("small", "G"), 1), time.Second*2/3))
	Encode(wavFile, NewCompositor(s1, s2, s3, s4, s5), 44100, 1)

}

func TestSaveVibrato(t *testing.T) {
	wavFile, err := os.Create("./test output/vibrato.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	s := NewMidiNote(MidiNote("one-line", "C"), 2000*ms, 1)
	sm := NewMidiNote(MidiNote("great","C"), 2000*ms, 1)
	Encode(wavFile, Modulated(s, sm, 1*ms), 8000, 1)
}

func TestSaveADSRModulate(t *testing.T) {
	wavFile, err := os.Create("./test output/ADSRModulate.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	sm := signals.Looped{signals.NewADSREnvelope(signals.X(.1), signals.X(.1), signals.X(.1), signals.Y(.7), signals.X(.1)), signals.X(.4)}

	s := NewMidiNote(MidiNote("two-line","C"), 3500*ms, 100)
	Encode(wavFile, Modulated(s, sm, 20*ms), 8000, 1)
}

func TestSaveHarmonicNotes(t *testing.T) {
	stream, err := os.Open("sample.wav")
	if err != nil {
		panic(err)
	}
	sample, err := signals.Decode(stream)
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	wavFile, err := os.Create("./test output/hNotes.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	sustainedEnv := func(length time.Duration) signals.LimitedSignal {
		return signals.NewADSREnvelope(signals.X(.025), signals.X(.1), signals.X(length.Seconds()), signals.Y(.5), signals.X(.5))
	}

	noteDuration := 180 * ms
	sustainDuration := 50 * ms
	notes := Compositor{}
	addMidiNote := func(t Compositor, noteNum int8, length, gap uint8) Compositor {
		//noteAndGap := Sound{signals.Product{NewToneMidi(noteNum, 80), sustainedEnv(signals.MultiplyInterval(length, sustainDuration))}, 140*ms+signals.MultiplyInterval(length+gap, noteDuration)}
		envNoteAndGap := NewSound(signals.Modulated{NewSampledMidiTone(noteNum, sample[0], .7), sustainedEnv(time.Duration(length) * sustainDuration)}, 140*ms+time.Duration(length+gap)*noteDuration)
		if len(t.Composite) == 0 {
			return NewCompositor(append(t.Composite, envNoteAndGap))
		}
		return NewCompositor(append(t.Composite, AfterPlusOffset(t.Composite[len(t.Composite)-1].(Sound), envNoteAndGap, -140*ms)))
	}

	TwinkleTwinkleLittleStar := []int8{60, 60, 67, 67, 69, 69, -67, 65, 65, 64, 64, 62, 62, -60}
	var noteLength uint8 = 1
	for _, note := range TwinkleTwinkleLittleStar {
		if note < 0 {
			notes = addMidiNote(notes, -note, noteLength*2+2, 0)
		} else {
			notes = addMidiNote(notes, note, noteLength*2, 0)
		}
	}
	// measured frequencies: 261 392  440 392 349 330 293.5 261 checked
	Encode(wavFile, notes, 44100, 1)
}

/*
func BenchmarkOne(b *testing.B) {
	b.StopTimer()
}

*/

