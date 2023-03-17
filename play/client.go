package play

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//	type List_Track struct {
//		Tracks []Track
//	}
type Track struct {
	Data     []byte `json:"data"`
	FileName string `json:"fileName`
}

// myPlaylist := Playlist{Name: "My Awesome Playlist", Songs: []string{"Song 1", "Song 2"}}
// AddSongToPlaylist(&myPlaylist, "Song 3")
// fmt.Println(myPlaylist.Songs)

func Start() {
	var result Track
	var req string
	//req, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	//req = strings.TrimSpace(req)
	req = ReturnNamemusic()

	resp, err := http.Get("http://localhost:9888/download/" + req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// err = ioutil.WriteFile("output.txt", file, 0644)
	// if err != nil {
	//     panic(err)
	// }

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(body))
	//m := (string(body))
	//return []byte(m)

	fmt.Println(result.FileName)

	err = ioutil.WriteFile(result.FileName, result.Data, 0644)
	if err != nil {
		panic(err)

	}

}
func ReturnNamemusic() string {
	var name string
	fmt.Print("Type an MP3 file name: ")
	name, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	name = strings.TrimSpace(name)

	return name
}
