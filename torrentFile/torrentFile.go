package torrentfile

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
	Name        string `bencode:"name"`
	Length      int    `bencode:"length"`
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
}

type bencodeTorrent struct {
	Announce     string      `bencode:"announce"`
	Comment      string      `bencode:"announce"`
	CreationDate int         `bencode:"creation date"`
	Info         bencodeInfo `bencode:"info"`
}

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func Open(path string) (*TorrentFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bt := bencodeTorrent{}
	err = bencode.Unmarshal(file, &bt)
	if err != nil {
		return nil, err
	}

	torrent, err := bt.toTorrentFile()
	if err != nil {
		return nil, err
	}

	return torrent, nil
}

func (bt *bencodeTorrent) toTorrentFile() (*TorrentFile, error) {
	infoHash, err := bt.Info.hash()
	if err != nil {
		return nil, err
	}

	piecesHashes, err := bt.Info.splitPieceHashes()
	if err != nil {
		return nil, err
	}

	torrentFile := TorrentFile{
		Announce:    bt.Announce,
		Name:        bt.Info.Name,
		Length:      bt.Info.Length,
		InfoHash:    infoHash,
		PieceHashes: piecesHashes,
	}

	return &torrentFile, nil
}

func (bi *bencodeInfo) hash() ([20]byte, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, *bi)
	if err != nil {
		return [20]byte{}, err
	}

	h := sha1.Sum(buf.Bytes())

	return h, nil
}

func (bi *bencodeInfo) splitPieceHashes() ([][20]byte, error) {
	hashLen := 20 // Length of SHA-1 hash
	buf := []byte(bi.Pieces)
	if len(buf)%hashLen != 0 {
		err := fmt.Errorf("received malformed pieces of length %d", len(buf))
		return nil, err
	}

	numHashes := len(buf) / hashLen
	hashes := make([][20]byte, numHashes)

	for i := 0; i < numHashes; i++ {
		copy(hashes[i][:], buf[i*hashLen:(i+1)*hashLen])
	}
	return hashes, nil
}
