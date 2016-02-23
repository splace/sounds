package sound

import (
	"../signals"
	"io"
	"time"
	//"fmt"
)

const ms = time.Millisecond

// Samples are Functions that have a duration.
type Sound struct {
	signals.LimitedFunction
}

func NewSound(sig signals.Function, d time.Duration) Sound {
	return Sound{signals.Multiplex{sig, signals.Pulse{signals.X(d)}}}
}

// Silence is a Sound with zero Level.
// can be used to give a duration to Stacks, that otherwise dont contain any Sounds, only neverending Functions.
func Silence(d time.Duration) (s Sound) {
	return NewSound(signals.NewConstant(0), d)
}

func Encode(w io.Writer, s Sound, sampleRate, sampleBytes uint) {
	signals.Encode(w, s, s.MaxX(), uint32(sampleRate), uint8(sampleBytes))
}

// Tracks is a Sound, made from adding together Functions.
func NewTracks(c ...signals.Function) Sound {
	return Sound{signals.NewSum(c...)}
}
