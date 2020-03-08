package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

type SongData struct {
	Length       string `json:"lengthSecound"`
	Codec        string `json:"codec"`
	Bitrate      string `json:"bitrate"`
	Artist       string `json:"artist"`
	Album        string `json:"album"`
	Date         string `json:"date"`
	Genre        string `json:"genre"`
	TrackNummber string `json:"trackNummber"`
	Title        string `json:"title"`
}

func main() {
	fmt.Println("Starting...")
	conn, err := net.Dial("tcp", ":3333")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	tmp := make([]byte, 1024)
	for {
		n, err := conn.Read(tmp)

		if err != nil {
			break
		}

		data := tmp[:n]
		dataByteArray := bytes.Split(data, []byte("|"))
		dataArray := strings.Split(string(data), "|")

		if dataArray[0] == "111" {
			song, err := parseJson(dataByteArray[4])
			if err != nil {
				fmt.Println("pasedJson failed")
			}
			fmt.Println(song)
			writeNowPlaying(song)
		}

	}
}

func parseJson(jsonBlob []byte) (SongData, error) {
	var song SongData
	err := json.Unmarshal(jsonBlob, &song)
	if err != nil {
		return song, err
	}

	return song, err
}

func writeNowPlaying(song SongData) {
	var toWrite = []byte(song.Artist + " - " + song.Title)
	ioutil.WriteFile("./np.txt", toWrite, 0777)
}
