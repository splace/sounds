/*
Package sounds generates and manipulates sounds.

Overview

Interfaces

Sound:- has constructors for various Tones/Notes, and modifiers for basic alterations, using time.Duration rather than the abstract units of the signals package. (is a signals.LimitedFunction)

Fundamental Types

Compositor:- an array of signals.Functions, that is a signals.LimitedFunction, and so is a Sound.

(the limit for the Compositor, that makes it a signals.LimitedFunction, comes from the contained signals.Functions, so at least one has to be a signals.LimitedFunction)

*/
package sound

