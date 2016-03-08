package sound

import (
	"github.com/splace/signals" //"../signals"//
	"time"
)

func Delayed(s Sound, offset time.Duration) Sound {
	return NewSound(signals.Shifted{signals.Modulated{signals.Heavyside{}, s}, signals.X(offset.Seconds())}, offset+time.Duration(s.MaxX()))
}

func Spedup(s Sound, f float32) Sound {
	return NewSound(signals.Spedup{s, f}, time.Duration(float32(s.MaxX())/f))
}

func After(p, s Sound) Sound {
	return Delayed(s, time.Duration(p.MaxX()))
}

func AfterPlusOffset(p, s Sound, offset time.Duration) Sound {
	return Delayed(s, time.Duration(p.MaxX())+offset)
}

func Reversed(s Sound) Sound {
	return NewSound(signals.Shifted{signals.Reversed{s}, s.MaxX()}, time.Duration(s.MaxX()))
}

func Modulated(s Sound, ms signals.Function, factor time.Duration) Sound {
	return NewSound(signals.RateModulated{s, ms, signals.X(factor.Seconds())}, time.Duration(s.MaxX()))
}
