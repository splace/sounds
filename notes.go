package sound

import (
	"github.com/splace/signals"
	"time"
)

// Notes are Signals that have a repeat period and a Duration.
type Note signals.PeriodicLimitedSignal

// make a Note from a PeriodicSignal, and a time.Duration.
// end moment set to the closest whole number of Periodical.Period()'s shorter than the length
func NewNote(source signals.PeriodicSignal, length time.Duration) Note {
	period := time.Duration(float32(source.Period()) / float32(signals.X(1)) * float32(time.Second))
	length -= (length + period) % period
	return signals.Modulated{source, signals.Pulse{signals.X(length.Seconds())}}
}

var semitoneNumber = map[string]int8{"C": 0, "C#": 1, "Db": 1, "D": 2, "D#": 3, "Eb": 3, "E": 4, "F": 5, "F#": 6, "Gb": 6, "G": 7, "G#": 8, "Ab": 8, "A": 9, "A#": 10, "Bb": 10, "B": 11}
var octaveNumber = map[string]int8{"double-contra": -1, "sub-contra": 0, "contra": 1, "great": 2, "small": 3, "one-line": 4, "once-accented": 4, "two-line": 5, "twice-accented": 5, "three-line": 6, "thrice-accented": 6, "four-line": 7, "four-times-accented": 7, "five-line": 8, "six-line": 8}
var semitonePrefixes = [...]string{"", "", "low", "base", "middle", "treble", "high", ""}
var semitones = [...]string{"C", "C#/Db", "D", "D#/Eb", "E", "F", "F#/Gb", "G", "G#/Ab", "A", "A#/Bb", "B"}

func Period(octave, semiNote interface{}) time.Duration {
	if o,ok:=octave.(int);ok{
		if s,ok:=semiNote.(int);ok{
			return PeriodFromMilliHz(FrequencyMilliHz(int8(o), int8(s)))
		}
		if s,ok:=semiNote.(string);ok{
			return PeriodFromMilliHz(FrequencyMilliHz(int8(o), semitoneNumber[s]))
		}
	}
	if o,ok:=octave.(string);ok{
		if s,ok:=semiNote.(int);ok{
			return PeriodFromMilliHz(FrequencyMilliHz(octaveNumber[o], int8(s)))
		}
		if s,ok:=semiNote.(string);ok{
			return PeriodFromMilliHz(FrequencyMilliHz(octaveNumber[o], semitoneNumber[s]))
		}
	}
	return 0
}

func PeriodFromMilliHz(mhz uint) time.Duration {
	return 1000 * time.Second / time.Duration(mhz)
}

func Frequency(period time.Duration) uint {
	return uint(time.Second / period)
}



