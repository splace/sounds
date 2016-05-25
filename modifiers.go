package sound

import (
	"github.com/splace/signals" //"../signals"//
	"time"
)

func Delayed(source Sound, offset time.Duration) Sound {
	return NewSound(signals.Shifted{signals.Modulated{signals.Heavyside{}, source}, signals.X(offset.Seconds())}, offset+time.Duration(source.MaxX()))
}

func Spedup(source Sound, factor float32) Sound {
	return NewSound(signals.Spedup{source, factor}, time.Duration(float32(source.MaxX())/factor))
}

func After(waitedFor, source Sound) Sound {
	return Delayed(source, time.Duration(waitedFor.MaxX()))
}

func AfterPlusOffset(waitedFor, source Sound, offset time.Duration) Sound {
	return Delayed(source, time.Duration(waitedFor.MaxX())+offset)
}

func Reversed(source Sound) Sound {
	return NewSound(signals.Shifted{signals.Reversed{source}, source.MaxX()}, time.Duration(source.MaxX()))
}

func Modulated(source Sound, modulation signals.Signal, factor time.Duration) Sound {
	return NewSound(signals.RateModulated{source, modulation, signals.X(factor.Seconds())}, time.Duration(source.MaxX()))
}
