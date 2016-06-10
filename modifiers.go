package sound

import (
	"github.com/splace/signals" 
	"time"
)

//  returns a Sound that is another Sound with silence prepended.   
func Delayed(source Sound, offset time.Duration) Sound {
	return Sound(signals.ShiftedLimitedSignal{source, signals.X(offset.Seconds())})
}

// returns a Sound that is another Sound rate adjusted.
func Spedup(source Sound, factor float32) Sound {
	return Sound(signals.CompressedLimitedSignal{source, factor})
}

// returns a Delayed Sound, so it starts just as another sounds ends.
func After(waitedFor, source Sound) Sound {
	return Delayed(source, time.Duration(waitedFor.MaxX()))
}

// returns a Delayed Sound, so it starts from an offset after another sounds ends.
func AfterPlusOffset(waitedFor, source Sound, offset time.Duration) Sound {
	return Delayed(source, time.Duration(waitedFor.MaxX())+offset)
}

// returns a sound in reversed.
func Reversed(source Sound) Sound {
	return NewSound(signals.Shift(signals.Reversed{source}, source.MaxX()), time.Duration(source.MaxX()))
}

// returns a sound that is rate adjusted depending on the value of a signal, potentially another sound. 
func Modulated(source Sound, modulation signals.Signal, factor time.Duration) Sound {
	return NewSound(signals.RateModulated{source, modulation, signals.X(factor.Seconds())}, time.Duration(source.MaxX()))
}
/*  Hal3 Fri Jun 10 16:28:22 BST 2016 go version go1.5.1 linux/amd64
=== RUN   TestSaveTone
--- PASS: TestSaveTone (0.03s)
=== RUN   TestSaveSound
--- PASS: TestSaveSound (0.74s)
=== RUN   TestSaveFlattenedSound
--- PASS: TestSaveFlattenedSound (0.08s)
=== RUN   TestSaveNote
--- PASS: TestSaveNote (0.02s)
=== RUN   TestLoad
1
--- PASS: TestLoad (0.06s)
=== RUN   TestLoadChannels
2
--- PASS: TestLoadChannels (0.26s)
=== RUN   TestSaveSignal
--- PASS: TestSaveSignal (0.02s)
=== RUN   TestSaveModifiedNote
--- PASS: TestSaveModifiedNote (0.08s)
=== RUN   TestSaveModifiedWav
--- PASS: TestSaveModifiedWav (0.24s)
=== RUN   TestSaveWavSoundAfterSound
--- PASS: TestSaveWavSoundAfterSound (0.09s)
=== RUN   TestSaveVibrato
--- PASS: TestSaveVibrato (0.02s)
=== RUN   TestSaveADSRModulate
--- PASS: TestSaveADSRModulate (0.03s)
=== RUN   TestSaveHarmonicNotes
--- PASS: TestSaveHarmonicNotes (0.81s)
PASS
ok  	_/home/simon/Dropbox/github/working/sound	2.480s
Fri Jun 10 16:28:26 BST 2016 */

