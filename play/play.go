package play

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func Play() {
	m := ReturnNamemusic()

	f, err := os.Open(m)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// streamer— это то, что мы можем использовать, чтобы сыграть песню
	// format- сведения о песне, самое главное, о ее частоте дискретизации
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(1, streamer), Paused: false}
	speaker.Play(ctrl)
	for {
		fmt.Print("Press [ENTER] to pause/resume. ")
		fmt.Scanln()

		speaker.Lock()
		ctrl.Paused = !ctrl.Paused
		speaker.Unlock()
	}

	//speaker.Play(ctrl)

}

// ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
// for {
// 	fmt.Print("Press [ENTER] to pause/resume. ")
// 	fmt.Scanln()

// 	speaker.Lock()
// 	ctrl.Paused = !ctrl.Paused
// 	speaker.Unlock()
// }
