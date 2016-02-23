/*
Package sounds generates and manipulates sounds.

Overview

Fundamental Types

Compositor:- an array of signals.Functions, that is a signals.LimitedFunction.

(the limit for the Compositor, that makes it a signals.LimitedFunction, comes from the contained signals.Functions, so at least one has to be signals.LimitedFunction)

Interfaces

Sounds are signals.LimitedFunction whose MaxX() is set with a time.Duration.

(see signals package for advanced manipulation of signals.LimitedFunction.)

*/
package sound

