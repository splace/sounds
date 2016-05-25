/*
Package sounds generates and manipulates sounds.

Overview

a 'higher level' package than github.com/splace/signals, tailored for sounds.
x-axis a time.Duration and y-axis ranging from +1 to -1.
composition to allow overlapping and some methods to help with music.

Interfaces

Sound:- time limited Signal (a signals.LimitedSignal.)

Tone:- neverending repeating wave, (a signals.PeriodicSignal.)

Note:- a Sound with a duration set to a whole number of Period's, (a signals.PeriodicLimitedSignal.)

(along with basic access with whats in this package, these interfaces can be used by func's from the signals package)

Types

Compositor:- an array of signals.Signal, potentially Sound's, that is itself a Sound, and is generated by adding together its contents.

(the end time of a Compositor (that makes it a signals.LimitedSignal), comes from the longest contained signals.Signal's, so at least one should be a Sound)

*/
package sound
