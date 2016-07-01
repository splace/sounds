package sound

import (
	"github.com/splace/signals"
	"math"
	"time"
)

// offset in polar attenuation space from omnidirecional hearing
const humanOmniOffset = 0.5 
const humanOmniAngle = math.Pi * .1
const humanEarOffset = 0.25
const rateOfSound = time.Second/340

type vector struct{
	x,y float64
}

func distance(point vector) float64 {
	return math.Sqrt(point.x*point.x + point.y*point.y)
}

func angle(point vector) float64 {
	return math.Atan2(point.x,point.y)
}


// returns a Sound adjusted for location.
func Spatialized(source Sound, location vector ) Sound {
	return signals.Modulated{Delayed(source,rateOfSound*time.Duration(distance(location))), signals.NewConstant(signals.DB(attenuation(location)))}
}

func attenuation(location vector) float64{
	return angleAttenuation(angle(location))/distance(location)
}

func angleAttenuation(angle float64) float64{
	return math.Sqrt(1+humanOmniOffset*(humanOmniOffset+math.Cos(angle+humanOmniAngle)))
}

func Stereo(source Sound, location vector)(left,right Sound){
	return Spatialized(source,vector{-location.y,location.x-humanEarOffset}),Spatialized(source,vector{location.y,-location.x-humanEarOffset})
}

