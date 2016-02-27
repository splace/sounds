// lookup DTMF tones from key 
package DTMF

import . "github.com/splace/signals"
import "time"

// the minimum inter-digit interval shall be  45msec, the minimum pulse duration shall be 50msec, and the minimum duty cycle for ANSI-compliance shall be 100msec

// make a Function from two Sine waves scaled and added together.
func DuelTone(p1,p2 time.Duration) Stack {
	return Stack{Sine{X(p1)},Sine{X(p2)}}
}

var Tones = map[rune]Function{
	'0': DuelTone(time.Second/941, time.Second/1336),
	'1': DuelTone(time.Second/697, time.Second/1209),
	'2': DuelTone(time.Second/697, time.Second/1336),
	'3': DuelTone(time.Second/697, time.Second/1477),
	'4': DuelTone(time.Second/770, time.Second/1209),
	'5': DuelTone(time.Second/770, time.Second/1336),
	'6': DuelTone(time.Second/770, time.Second/1477),
	'7': DuelTone(time.Second/852, time.Second/1209),
	'8': DuelTone(time.Second/852, time.Second/1336),
	'9': DuelTone(time.Second/852, time.Second/1477),
	'A': DuelTone(time.Second/697, time.Second/1633),
	'B': DuelTone(time.Second/770, time.Second/1633),
	'C': DuelTone(time.Second/852, time.Second/1633),
	'D': DuelTone(time.Second/941, time.Second/1633),
	'*': DuelTone(time.Second/941, time.Second/1209),
	'#': DuelTone(time.Second/941, time.Second/1477),
}

