// command line tool for generating telephone tone dialing codes (DTMF) 
package main

import . "github.com/splace/sounds"  // "../../sound" //
import  "github.com/splace/sounds/examples/DTMF" // "./DTMF" // 
import "os"
import "flag"
import "time"
import	"log"

func main() {
	var gap,width time.Duration
	var sampleRate,sampleBytes uint
	flag.DurationVar(&gap, "gap", 80*time.Millisecond, "gap between pulses")
	flag.DurationVar(&width, "width", 70*time.Millisecond, "pulse width")
	flag.UintVar(&sampleRate, "rate", 44100, "sample per second")
	flag.UintVar(&sampleBytes,"bytes", 2, "bytes per sample")
	codeString:=flag.String("code", "0123456789", "<digits>*#ABCD to be encoded.")

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
	if len(code)==0{log.Fatal("need a least 1 key.")}

	tones:=NewCompositor()
	tones.Composite=append(tones.Composite,NewSound(DTMF.Tones[code[0]],width))
	for _,c:=range(code[1:]){
		// add new Sound, a DTMF.Tone with width, that starts when the previous entry in the slice ends, plus a gap.
		// Compose is a slice of Function, an interface that doesn't have to contain Sounds, so needs a type assertion to Sound, a type that has an end.  
		tones.Composite = append(tones.Composite,AfterPlusOffset(tones.Composite[len(tones.Composite)-1].(Sound), NewSound(DTMF.Tones[c],width),gap))
	}
	Encode(wavFile,tones, int(sampleRate),int(sampleBytes) )
}


