// lookup DTMF tones from key 
package DTMF

import . "github.com/splace/signals"

// the minimum inter-digit interval shall be  45msec, the minimum pulse duration shall be 50msec, and the minimum duty cycle for ANSI-compliance shall be 100msec

var Tones = map[rune]Stack{
	'0': Stack{Sine{X(1.0/941)},Sine{X(1.0/1336)}},
	'1': Stack{Sine{X(1.0/697)},Sine{X(1.0/1209)}},
	'2': Stack{Sine{X(1.0/697)},Sine{X(1.0/1336)}},
	'3': Stack{Sine{X(1.0/697)},Sine{X(1.0/1477)}},
	'4': Stack{Sine{X(1.0/770)},Sine{X(1.0/1209)}},
	'5': Stack{Sine{X(1.0/770)},Sine{X(1.0/1336)}},
	'6': Stack{Sine{X(1.0/770)},Sine{X(1.0/1477)}},
	'7': Stack{Sine{X(1.0/852)},Sine{X(1.0/1209)}},
	'8': Stack{Sine{X(1.0/852)},Sine{X(1.0/1336)}},
	'9': Stack{Sine{X(1.0/852)},Sine{X(1.0/1477)}},
	'A': Stack{Sine{X(1.0/697)},Sine{X(1.0/1633)}},
	'B': Stack{Sine{X(1.0/770)},Sine{X(1.0/1633)}},
	'C': Stack{Sine{X(1.0/852)},Sine{X(1.0/1633)}},
	'D': Stack{Sine{X(1.0/941)},Sine{X(1.0/1633)}},
	'*': Stack{Sine{X(1.0/941)},Sine{X(1.0/1209)}},
	'#': Stack{Sine{X(1.0/941)},Sine{X(1.0/1477)}},
}

