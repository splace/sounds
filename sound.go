package sound

import (
	"github.com/splace/signals" 
	"io"
	"time"
	//"fmt"
)

const ms = time.Millisecond

// Sounds are Signals that have a duration.
type Sound signals.LimitedSignal

// a Sound made from a Signal limited to a specified end time. 
func NewSound(source signals.Signal, d time.Duration) Sound {
	return signals.Modulated{source, signals.Pulse{signals.X(d.Seconds())}}
}

// Tones are Signals that have a repeat period, but no end time.
type Tone signals.PeriodicSignal

// a Tone made from a Sine wave and a volume.
func NewTone(period time.Duration, volume float64) Tone {
	return signals.NewTone(signals.X(period.Seconds()), signals.DB(volume))
}

// a Tone whose source is a Sound scaled to fit the period, and looped.
func NewSampledTone(period time.Duration, sample Sound, volume float64) Tone {
	return signals.Modulated{signals.Looped{Spedup(sample, float32(sample.MaxX())/float32(period)), signals.X(period.Seconds())}, signals.NewConstant(signals.DB(volume))}
}

// Compositor embeds signals.Compose, which adds together an array of Signals, which can be Sounds.
type Compositor struct {
	signals.Composite
}

// make a Compositor from Signal's, (use directly from a slice of narrower interfaces, it will require a slice promoter.)
func NewCompositor(sources ...signals.Signal) Compositor {
	return Compositor{signals.NewComposite(sources...)}
}

// Silence is a Sound with zero Level.
// can be used to give a duration to a Compositor, that otherwise doesn't contain any Sounds, only neverending Signals.
func Silence(duration time.Duration) (s Sound) {
	return NewSound(signals.NewConstant(0), duration)
}

// encode a Sound as PCM, with a particular sampleRate and sampleBytes precision.
func Encode(too io.Writer, source Sound, sampleRate, sampleBytes int) {
	signals.Encode(too, source, source.MaxX(), uint32(sampleRate), uint8(sampleBytes))
}

