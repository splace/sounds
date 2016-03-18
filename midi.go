package sound

import (
	"math"
	"time"
)

// make a Note from a Midi note number, a time.Duration and a volume percentage.
// duration set to the closest whole number of Periodical.Period()s shorter than the length
func NewMidiNote(noteNumber int8, length time.Duration, volume float64) Note {
	return NewNote(NewMidiTone(noteNumber, volume), length)
}

// make a Tone, a continuous Sine wave, from a Midi-note number and a volume.
func NewMidiTone(noteNumber int8, volume float64) Tone {
	return NewTone(PeriodFromMilliHz(FrequencyMilliHzMidi(noteNumber)), volume)
}

// make a Tone, whose source is a Sound looped and scaled to fit a Midi note period.
// more simply, it is a Tone made from a sound sample.
// samples can come from a file, using PCM decode function, which can be a standard wav encoded instrument sample.
func NewSampledMidiTone(noteNumber int8, sample Sound, volume float64) Tone {
	return NewSampledTone(PeriodFromMilliHz(FrequencyMilliHzMidi(noteNumber)), sample, volume)
}

const baseNoteNumber = 69
const baseFrequency = 440000 // mHz

// frequency as an int, in 1/1000 of a Hz units
func FrequencyMilliHzMidi(noteNumber int8) uint {
	return uint(baseFrequency * math.Pow(2, float64(noteNumber-baseNoteNumber)/12))
}

// frequency as in int, in 1/1000 of a Hz units
func FrequencyMilliHz(octave, semiNote int8) uint {
	return FrequencyMilliHzMidi(MidiNoteNumber(octave, semiNote))
}

func MidiNoteNumber(octave, semiNote int8) int8 {
	return (octave+1)*12 + semiNote
}

func MidiNote(octave, semiNote string) int8 {
	return (octaveNumber[octave]+1)*12 + semitoneNumber[semiNote]
}

func NameFromMidiNoteNumber(noteNumber int8) string {
	return semitonePrefixes[noteNumber/12] + semitones[noteNumber%12]
}

