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
	flag.DurationVar(&gap, "gap", 180*time.Millisecond, "gap between pulses")
	flag.DurationVar(&width, "width", 170*time.Millisecond, "pilse width")
	flag.UintVar(&sampleRate, "rate", 44100, "sample per second")
	flag.UintVar(&sampleBytes,"bytes", 2, "bytes per sample")
	codeString:=flag.String("code", "0123456789*#ABCD", "<digits>*#ABCD for encode")

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

	tones:=NewComposition()
	tones.Sum=append(tones.Sum,NewSound(DTMF.Tones[code[0]],width))
	for _,c:=range(code){
		tones.Sum = append(tones.Sum,AfterPlusOffset(tones.Sum[len(tones.Sum)-1].(Sound), NewSound(DTMF.Tones[c],width),gap))
	}
	Encode(wavFile,tones, sampleRate,sampleBytes )
}



