/*
Package sounds generates and manipulates sounds.

Overview

Fundamental Types

Compositor:- an array of signals.Functions, that is a signals.LimitedFunction.

(the limit for the Compositor, that makes it a signals.LimitedFunction, comes from the contained signals.Functions, so at least one has to be signals.LimitedFunction)

Interfaces

Sounds are signals.LimitedFunction whose MaxX() is set with time.Durations.

generators

NewTone :- make a continuous Sine wave from a period and a volume percentage.

NewNote :- make a sound from a Midi note number, a time.Duration and a volume percentage.


NewToneMidi:- make a continuous signal from a Midi-note number and a volume.

NewNoteMidi :- make a sound from a Midi note number, a time.Duration and a volume percentage.

*/
package sound

