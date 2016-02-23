//Formatted:Tue Feb 2 21:16:02 GMT 2016
package sound

import (
	"../signals"
	"time"
)

// Notes are Tones with their duration set to the closest whole number of Periodical.Period()s shorter than the length
func NewNote(sig signals.PeriodicFunction, length time.Duration) Sound {
	period := time.Duration(float32(sig.Period()) / float32(signals.UnitX) * float32(time.Second))
	length -= (length + period) % period
	// length-=(length+period/2)%period  // half cycle matching
	return Sound{signals.Multiplex{sig, signals.Pulse{signals.X(length)}}}
}

var SemitoneNumber = map[string]int8{"C": 0, "C#": 1, "Db": 1, "D": 2, "D#": 3, "Eb": 3, "E": 4, "F": 5, "F#": 6, "Gb": 6, "G": 7, "G#": 8, "Ab": 8, "A": 9, "A#": 10, "Bb": 10, "B": 11}
var OctaveNumber = map[string]int8{"double-contra": -1, "sub-contra": 0, "contra": 1, "great": 2, "small": 3, "one-line": 4, "once-accented": 4, "two-line": 5, "twice-accented": 5, "three-line": 6, "thrice-accented": 6, "four-line": 7, "four-times-accented": 7, "five-line": 8, "six-line": 8}
var SemitonePrefixes = [...]string{"", "", "low", "base", "middle", "treble", "high", ""}
var Semitones = [...]string{"C", "C#/Db", "D", "D#/Eb", "E", "F", "F#/Gb", "G", "G#/Ab", "A", "A#/Bb", "B"}

func Period(octave, semiNote int8) time.Duration {
	return PeriodFromCentiHz(FrequencyCentiHz(octave, semiNote))
}

func PeriodFromCentiHz(chz int) time.Duration {
	return 100 * time.Second / time.Duration(chz)
}
