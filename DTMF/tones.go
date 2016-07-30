// lookup DTMF tones from key, this is a plugin for; github.com/splace/signals, for the generation and manipulation of Telephone tones.
package DTMF

import . "github.com/splace/signals"

// the minimum inter-digit interval shall be  45msec, the minimum pulse duration shall be 50msec, and the minimum duty cycle for ANSI-compliance shall be 100msec

var Tones = map[rune]Stacked{
	'0': Stacked{Sine{X(1.0/941)},Sine{X(1.0/1336)}},
	'1': Stacked{Sine{X(1.0/697)},Sine{X(1.0/1209)}},
	'2': Stacked{Sine{X(1.0/697)},Sine{X(1.0/1336)}},
	'3': Stacked{Sine{X(1.0/697)},Sine{X(1.0/1477)}},
	'4': Stacked{Sine{X(1.0/770)},Sine{X(1.0/1209)}},
	'5': Stacked{Sine{X(1.0/770)},Sine{X(1.0/1336)}},
	'6': Stacked{Sine{X(1.0/770)},Sine{X(1.0/1477)}},
	'7': Stacked{Sine{X(1.0/852)},Sine{X(1.0/1209)}},
	'8': Stacked{Sine{X(1.0/852)},Sine{X(1.0/1336)}},
	'9': Stacked{Sine{X(1.0/852)},Sine{X(1.0/1477)}},
	'A': Stacked{Sine{X(1.0/697)},Sine{X(1.0/1633)}},
	'B': Stacked{Sine{X(1.0/770)},Sine{X(1.0/1633)}},
	'C': Stacked{Sine{X(1.0/852)},Sine{X(1.0/1633)}},
	'D': Stacked{Sine{X(1.0/941)},Sine{X(1.0/1633)}},
	'*': Stacked{Sine{X(1.0/941)},Sine{X(1.0/1209)}},
	'#': Stacked{Sine{X(1.0/941)},Sine{X(1.0/1477)}},
}

