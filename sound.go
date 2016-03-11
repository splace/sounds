package sound

import (
	"github.com/splace/signals" //"../signals"//
	"io"
	"time"
	//"fmt"
)

const ms = time.Millisecond

// Sounds are Functions that have a duration.
type Sound signals.LimitedFunction

// signals.unitX maps to a time.Second()
// so with signals.x as an int64, gives a range of nanosecond to 290years.(cf time.Duration)
func NewSound(sig signals.Function, d time.Duration) Sound {
	return signals.Modulated{sig, signals.Pulse{signals.X(d.Seconds())}}
}

// Tones are Functions that have a repeat period.
type Tone signals.PeriodicFunction

// make a continuous Sine wave from a period and a volume.
func NewTone(period time.Duration, volume float64) Tone {
	return signals.NewTone(signals.X(period.Seconds()), signals.DB(volume))
}

// make a continuous wave whose source is a Sound scaled to fit the period, and looped.
func NewSampledTone(period time.Duration, sample Sound, volume float64) Tone {
	return signals.Modulated{signals.Looped{Spedup(sample, float32(sample.MaxX())/float32(period)), signals.X(period.Seconds())}, signals.NewConstant(signals.DB(volume))}
}

// Compositor embeds signals.Compose, which adds together an array of Functions, which can be Sounds.
type Compositor struct {
	signals.Compose
}

// make compositor from Function's, (use directly from a slice of narrower interfaces, will require a slice promoter.)
func NewCompositor(c ...signals.Function) Compositor {
	return Compositor{signals.NewCompositor(c...)}
}

// Silence is a Sound with zero Level.
// can be used to give a duration to Compositor, that otherwise dont contain any Sounds, only neverending Functions.
func Silence(d time.Duration) (s Sound) {
	return NewSound(signals.NewConstant(0), d)
}

// encode a Sound as PCM, with a particular sampleRate and sampleBytes precision.
func Encode(w io.Writer, s Sound, sampleRate, sampleBytes uint) {
	signals.Encode(w, s, s.MaxX(), uint32(sampleRate), uint8(sampleBytes))
}


