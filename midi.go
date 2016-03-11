package sound

import (
	"math"
	"time"
)

// make a sound from a Midi note number, a time.Duration and a volume percentage.
// duration set to the closest whole number of Periodical.Period()s shorter than the length
func NewMidiNote(n int8, length time.Duration, volume float64) Note {
	return NewNote(NewMidiTone(n, volume), length)
}

// make a continuous Sine wave from a Midi-note number and a volume.
func NewMidiTone(n int8, volume float64) Tone {
	return NewTone(PeriodFromMilliHz(FrequencyMilliHzMidi(n)), volume)
}

// make a continuous wave whose source is a Sound looped and scaled to fit a Midi note period.
// more simply, it is a Tone made from a sound sample.
// samples can come from a file, using PCM decode function, which can be a standard wav encoded instrument sample.
func NewSampledMidiTone(n int8, sample Sound, volume float64) Tone {
	return NewSampledTone(PeriodFromMilliHz(FrequencyMilliHzMidi(n)), sample, volume)
}

const baseNoteNumber = 69
const baseFrequency = 440000 // mHz

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
