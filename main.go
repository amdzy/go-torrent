package main

import (
	torrentfile "amdzy/go-torrent/torrentFile"
	"log"
	"os"
	"path"
)

const Port uint16 = 6881

func main() {
	file := os.Args[1]

	tf, err := torrentfile.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	outPath := path.Join(tf.Name)
	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
