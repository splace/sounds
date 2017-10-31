/*
Package sounds generates and manipulates sounds. A 'high level' wrapper for github.com/splace/signals, tailored for sound.
Generates tones, notes, sampled notes.
Reads/saves PCM files.

Overview

x-axis a time.Duration and y-axis ranging from +1 to -1.
composition to allow overlapping and some methods to help with music.

Basic Types

Sound:- time limited Signal (implements signals.LimitedSignal.)

Tone:- never-ending repeating wave, (implements signals.PeriodicSignal.)

Note:- a Sound with a duration set to a whole number of Period's, (implements signals.PeriodicLimitedSignal.)

Compositor:- an array of signals.Signal, potentially Sound's, that is itself a Sound, and is generated by adding together its contents.

(the end time of a Compositor (that makes it a signals.LimitedSignal), comes from the longest contained signals.Signal's, so at least one should be a Sound, not, for example, all Tones. Timed Silence could be used for this.)

*/
package sound
