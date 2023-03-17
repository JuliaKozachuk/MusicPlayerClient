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

type Track struct {
	Data     []byte `json:"data"`
	FileName string `json:"fileName`
}

func Start() {
	var result Track
	var req string

	req = ReturnNamemusic()

	resp, err := http.Get("http://localhost:9888/download/" + req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	if err != nil {
		log.Fatal(err)
	}

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
