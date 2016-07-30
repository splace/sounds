package sound

import (
	"github.com/splace/signals"
	"io"
	"time"
)

const ms = time.Millisecond

// Sounds are Signals that have a duration.
type Sound signals.LimitedSignal

// make a Sound from a Signal (potentially unlimited), limited to a duration.
func NewSound(source signals.Signal, d time.Duration) Sound {
	return signals.Modulated{source, signals.Pulse{signals.X(d.Seconds())}}
}

// Tones are Signals that have a repeat period, but no end.
type Tone signals.PeriodicSignal

// a Tone made from a Sine wave and a volume.
func NewTone(period time.Duration, volume float64) Tone {
	return signals.Modulated{signals.Sine{signals.X(period.Seconds())}, signals.NewConstant(signals.DB(volume))}
}

// Tone made from a Sound, spedup or slowed down, to fit the required period, and looped.
func NewSampledTone(period time.Duration, sample Sound, volume float64) Tone {
	return signals.Modulated{signals.Looped{Spedup(sample, float32(sample.MaxX())/float32(period)), signals.X(period.Seconds())}, signals.NewConstant(signals.DB(volume))}
}

// Sequence embeds signals.Sequenced, which appends together an array of Signals.
type Sequence struct {
	signals.Sequenced
}

// make a Sequence of Sound's
func NewSequencer(sources ...Sound) Sequence {
	return Sequence{signals.NewSequence(SoundsToLimitedSignals(sources)...)}
}

// converts to []LimitedSignal
func SoundsToLimitedSignals(s []Sound) []signals.LimitedSignal {
	out := make([]signals.LimitedSignal, len(s))
	for i := range out {
		out[i] = s[i].(signals.LimitedSignal)
	}
	return out
}

// Compositor embeds signals.Compose, which adds together an array of Signals, which can be Sounds.
type Compositor struct {
	signals.Composite
}

// make a Compositor from Signal's, (to be used directly from a slice of narrower interfaces, it will require one of the slice promoter functions.)
func NewCompositor(sources ...signals.Signal) Compositor {
	return Compositor{signals.NewComposite(sources...)}
}

// Silence is a Sound with zero Level.
// can be used to give a duration to a Compositor, that might otherwise not contain any Sounds, only neverending Signals.
func Silence(duration time.Duration) (s Sound) {
	return NewSound(signals.NewConstant(0), duration)
}

// encode Sounds as multichannel PCM, with a particular sampleRate and sampleBytes precision.
func Encode(w io.Writer, sampleBytes, sampleRate int, sources ...signals.LimitedSignal) {
	max:= signals.X(0)
	for _, s := range sources {
		if sls, ok := s.(signals.LimitedSignal); ok {
			if newmax := sls.MaxX(); newmax > max {
				max = newmax
			}
		}
	}
	signals.Encode(w, uint8(sampleBytes), uint32(sampleRate), max, signals.PromoteToSignals(sources)...)
}

// turn a slice of Sounds into a slice of those sounds each with a suffix sound appended.
func SoundsSuffixed(suffix Sound, sounds ...Sound) []Sound {
	out := make([]Sound, len(sounds))
	for i := range out {
		out[i] = NewSequencer(sounds[i],suffix)
	}
	return out
}

// turn a slice of Sounds into a slice of those sounds each with a prefix sound appended.
func SoundsPrefixed(prefix Sound, sounds ...Sound) []Sound {
	out := make([]Sound, len(sounds))
	for i := range out {
		out[i] = NewSequencer(prefix, sounds[i])
	}
	return out
}

// turn a slice of Sounds into a slice of those sounds with the same separator sound between each.
func SoundsSeparated(sep Sound, sounds ...Sound) []Sound {
	out := make([]Sound, len(sounds)*2-1)
	for i := range out {
		if i/2*2==i {
			out[i] = sounds[i/2]
			}else{
			out[i]=sep
			}
	}
	return out
}



