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
	return NewSound(signals.Shifted(signals.Reversed{source}, source.MaxX()), time.Duration(source.MaxX()))
}

// returns a sound that is rate adjusted depending on the value of a signal, potentially another sound. 
func Modulated(source Sound, modulation signals.Signal, factor time.Duration) Sound {
	return NewSound(signals.RateModulated{source, modulation, signals.X(factor.Seconds())}, time.Duration(source.MaxX()))
}

