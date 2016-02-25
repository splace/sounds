package sound

import (
	"github.com/splace/signals"	//"../signals"//
	"io"
	"time"
	//"fmt"
)

const ms = time.Millisecond

// Sounds are Functions that have a duration.
type Sound signals.LimitedFunction

func NewSound(sig signals.Function, d time.Duration) Sound {
	return signals.Multiplex{sig, signals.Pulse{signals.X(d)}}
}

// encode a Sound as PCM, with a particular sampleRate and sampleBytes precision.
func Encode(w io.Writer, s Sound, sampleRate, sampleBytes uint) {
	signals.Encode(w, s, s.MaxX(), uint32(sampleRate), uint8(sampleBytes))
}

// make a continuous Sine wave from a period and a volume.
func NewTone(period time.Duration, volume float64) signals.PeriodicFunction {
	return signals.NewTone(signals.X(period), signals.DB(volume))
}
// make a continuous wave whose source is a Sound scaled to fit the period, and looped.
func NewSampledTone(period time.Duration, sample Sound, volume float64) signals.PeriodicFunction {
	return signals.Multiplex{signals.Looped{Spedup(sample, float32(sample.MaxX())/float32(period)), signals.X(period)}, signals.NewConstant(signals.DB(volume))}
}

// Compositor contains signals.Compose, an array of Functions, which can be Sounds.
type Compositor struct{
	signals.Compose
}

func NewCompositor(c ...signals.Function) Compositor {
	return Compositor{signals.NewSum(c...)}
}

// Silence is a Sound with zero Level.
// can be used to give a duration to Compositor, that otherwise dont contain any Sounds, only neverending Functions.
func Silence(d time.Duration) (s Sound) {
	return NewSound(signals.NewConstant(0), d)
}



