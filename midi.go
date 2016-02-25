package sound

import (
	"github.com/splace/signals"	//"../signals"//
	"math"
	"time"
)

// make a sound from a Midi note number, a time.Duration and a volume percentage.
// duration set to the closest whole number of Periodical.Period()s shorter than the length
func NewNoteMidi(n int8, length time.Duration, volume float64) Sound {
	return NewNote(NewMidiTone(n, volume), length)
}

// make a continuous Sine wave from a Midi-note number and a volume.
func NewMidiTone(n int8, volume float64) signals.PeriodicFunction {
	return NewTone(PeriodFromMilliHz(FrequencyMilliHzMidi(n)), volume)
}

// make a continuous wave whose source is a Sound scaled to fit the period, and looped. more simply it is a Tone made from a sound sample and, using the PCM decode function, can be used with standard wav encoded instrument samples.
func NewSampledMidiTone(n int8, sample Sound, volume float64) signals.PeriodicFunction {
	return NewSampledTone(PeriodFromMilliHz(FrequencyMilliHzMidi(n)), sample, volume) 
}

const baseNoteNumber = 69
const baseFrequency = 440000   // mHz

// frequency as an int, in 1/1000 of a Hz units 
func FrequencyMilliHzMidi(noteNumber int8) int {
	return int(baseFrequency * math.Pow(2, float64(noteNumber-baseNoteNumber)/12))
}

// frequency as in int, in 1/1000 of a Hz units
func FrequencyMilliHz(octave, semiNote int8) int {
	return FrequencyMilliHzMidi(MidiNoteNumber(octave, semiNote))
}

func MidiNoteNumber(octave, semiNote int8) int8 {
	return (octave+1)*12 + semiNote
}

func NameFromMidiNoteNumber(noteNumber int8) string {
	return SemitonePrefixes[noteNumber/12] + Semitones[noteNumber%12]
}


