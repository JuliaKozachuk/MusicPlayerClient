package play

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Queue struct {
	Name      string
	streamers []beep.StreamSeekCloser
}

func Start_queue() {
	songNameCh := make(chan string)
	exitCh := make(chan bool)
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))

	myPlaylist := Queue{Name: "My Awesome Playlist", streamers: []beep.StreamSeekCloser{}}

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			select {
			case <-exitCh:
				return
			default:
				songName, _ := reader.ReadString('\n')
				songName = strings.TrimSpace(songName)
				songNameCh <- songName
			}
		}
	}()

	for {
		select {
		case songName := <-songNameCh:
			f, err := os.Open(songName)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// Decode it.
			streamer, format, err := mp3.Decode(f)
			if err != nil {
				fmt.Println(err)
				continue
			}

			myPlaylist.streamers = append(myPlaylist.streamers, streamer)

			length := format.SampleRate.D(streamer.Len())
			fmt.Println(length)
			fmt.Println(myPlaylist.streamers)
			fmt.Printf("Файл %s добавлен в плейлист\n", songName)

			if len(myPlaylist.streamers) == 1 {
				go func() {
					speaker.Play(myPlaylist.streamers[0])
				}()
			}

		case <-time.After(time.Second):
			if len(myPlaylist.streamers) > 1 {
				go func() {
					speaker.Lock()
					speaker.Clear()

					streamer := myPlaylist.streamers[0]
					myPlaylist.streamers = myPlaylist.streamers[1:]
					speaker.Play(streamer)
					speaker.Unlock()

				}()
			}
		}
	}
}
