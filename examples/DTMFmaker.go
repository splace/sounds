package main

import "./DTMF"
import . "../../sound"
import "os"
import "flag"
import "time"
import	"log"

func main() {
	var gap,width time.Duration
	var sampleRate,sampleBytes uint
	flag.DurationVar(&gap, "gap", 80*time.Millisecond, "gap between pulses")
	flag.DurationVar(&width, "width", 70*time.Millisecond, "pilse width")
	flag.UintVar(&sampleRate, "rate", 44100, "sample per second")
	flag.UintVar(&sampleBytes,"bytes", 2, "bytes per sample")
	codeString:=flag.String("code", "0123456789", "number or *#ABCD to encode")

	flag.Parse()
	file := flag.Args()
	code := []rune(*codeString)
	var wavFile *os.File
	var err error
	if len(file)>0 {
		wavFile, err = os.Create(file[0])
		if err != nil {
			log.Fatal(file[0],"Creation failure",err.Error())
		}
		defer wavFile.Close()
	}else{
		log.Fatal("need output file name argument.")
	}

	tone0 := NewSample(DTMF.Tones[code[0]],width)
	tone1 := AfterPlusOffset(tone0, NewSample(DTMF.Tones[code[1]],width),gap)
	tone2 := AfterPlusOffset(tone1, NewSample(DTMF.Tones[code[2]],width),gap)
	tone3 := AfterPlusOffset(tone2, NewSample(DTMF.Tones[code[3]],width),gap)
	tone4 := AfterPlusOffset(tone3, NewSample(DTMF.Tones[code[4]],width),gap)
	tone5 := AfterPlusOffset(tone4, NewSample(DTMF.Tones[code[5]],width),gap)
	tone6 := AfterPlusOffset(tone5, NewSample(DTMF.Tones[code[6]],width),gap)
	tone7 := AfterPlusOffset(tone6, NewSample(DTMF.Tones[code[7]],width),gap)
	tone8 := AfterPlusOffset(tone7, NewSample(DTMF.Tones[code[8]],width),gap)
	tone9 := AfterPlusOffset(tone8, NewSample(DTMF.Tones[code[9]],width),gap)
	Encode(wavFile,NewTracks(tone0, tone1, tone2, tone3, tone4, tone5, tone6, tone7, tone8, tone9), sampleRate,sampleBytes )
}

/*  Hal3 Tue Feb 23 01:37:31 GMT 2016 go version go1.5.1 linux/amd64
Tue Feb 23 01:37:32 GMT 2016 */

