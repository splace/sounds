// command line tool for generating telephone tone dialing codes (DTMF)
package main

import . "github.com/splace/sounds"
import "github.com/splace/sounds/DTMF"
import "os"
import "flag"
import "time"
import "log"



func main() {
	var gap, width time.Duration
	flag.DurationVar(&gap, "gap", 50*time.Millisecond, "gap between pulses")
	flag.DurationVar(&width, "width", 80*time.Millisecond, "pulse width")
	var sampleRate, sampleBytes uint
	flag.UintVar(&sampleRate, "rate", 22050, "sample per second")
	flag.UintVar(&sampleBytes, "bytes", 2, "bytes per sample")
	codeString := flag.String("code", "0123456789", "<digits>*#ABCD to be encoded.")

	flag.Parse()
	file := flag.Args()
	code := []rune(*codeString)
	var wavFile *os.File
	var err error
	if len(file) > 0 {
		wavFile, err = os.Create(file[0])
		if err != nil {
			log.Fatal(file[0], "Creation failure", err.Error())
		}
		defer wavFile.Close()
	} else {
		log.Fatal("need an output file name argument.")
	}
	if len(code) == 0 {
		log.Fatal("need a least 1 key.")
	}
	// the minimum inter-digit interval shall be  45msec, the minimum pulse duration shall be 50msec, and the minimum duty cycle for ANSI-compliance shall be 100msec
	if gap<45*time.Millisecond || width<50*time.Millisecond || gap+width<100*time.Millisecond {
		log.Println("Warning: Non-ANSI-complient timings.")
	}
	
	tones := make([]Sound, len(code))
	for i, c := range code {
		tones[i] = NewNote(DTMF.Tones[c], width)
	}
	Encode(wavFile, int(sampleBytes), int(sampleRate), NewSequencer(SoundsSeparated(Silence(gap),tones...)...))
}


