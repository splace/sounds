package sound

import (
	"github.com/splace/signals"	//"../signals"
	"math"
	"time"
)

// make a sound from a Midi note number, a time.Duration and a volume percentage.
// duration set to the closest whole number of Periodical.Period()s shorter than the length
func NewNoteMidi(n int8, length time.Duration, volume float64) Sound {
	return NewNote(NewToneMidi(n, volume), length)
}

// make a continuous signal from a Midi-note number and a volume.
func NewToneMidi(n int8, volume float64) signals.Multiplex {
	return NewTone(PeriodFromCentiHz(FrequencyCentiHzMidi(n)), volume)
}

const baseNoteNumber = 69
const baseFrequency = 44000

func FrequencyCentiHzMidi(noteNumber int8) int {
	return int(baseFrequency * math.Pow(2, float64(noteNumber-baseNoteNumber)/12))
}

func FrequencyCentiHz(octave, semiNote int8) int {
	return FrequencyCentiHzMidi(MidiNoteNumber(octave, semiNote))
}

func MidiNoteNumber(octave, semiNote int8) int8 {
	return (octave+1)*12 + semiNote
}
func NameFromMidiNoteNumber(noteNumber int8) string {
	return SemitonePrefixes[noteNumber/12] + Semitones[noteNumber%12]
}
