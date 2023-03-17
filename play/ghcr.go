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

func Hdfhn() {
	playlist := []beep.StreamSeekCloser{}

	for {
		fmt.Println("Введите имя файла (или \"стоп\" для выхода):")
		var filename string
		filename, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		filename = strings.TrimSpace(filename)

		if filename == "стоп" {
			break
		}

		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Ошибка при открытии файла %s: %v\n", filename, err)
			continue
		}

		streamer, format, err := mp3.Decode(f)
		if err != nil {
			fmt.Printf("Ошибка при декодировании файла %s: %v\n", filename, err)
			continue
		}

		playlist = append(playlist, streamer)

		go func(s beep.StreamSeekCloser) {
			done := make(chan bool)
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(beep.Seq(s, beep.Callback(func() {
				done <- true
			})))
			<-done
		}(streamer)

		fmt.Printf("Файл %s добавлен в плейлист\n", filename)
	}

	fmt.Println("Плейлист:")
	for i, streamer := range playlist {
		fmt.Printf("%d: %s (%s)\n", i+1, streamer.Len())
	}

	fmt.Println("Воспроизведение завершено")
}
