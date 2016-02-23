package sound

import (
	"github.com/splace/signals"
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

func Encode(w io.Writer, s Sound, sampleRate, sampleBytes uint) {
	signals.Encode(w, s, s.MaxX(), uint32(sampleRate), uint8(sampleBytes))
}


func NewTone(period time.Duration, volume float64) signals.Multiplex {
	return signals.NewTone(signals.X(period), signals.DB(volume))
}

// Compositor contains signals.Sum, an array of Functions, which can be Sounds.
type Compositor struct{
	signals.Compose
}

func NewCompositor(c ...signals.Function) Compositor {
	return Compositor{signals.NewSum(c...)}
}

// Silence is a Sound with zero Level.
// can be used to give a duration to Stacks, that otherwise dont contain any Sounds, only neverending Functions.
func Silence(d time.Duration) (s Sound) {
	return NewSound(signals.NewConstant(0), d)
}



